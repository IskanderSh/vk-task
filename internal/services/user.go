package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	validator "github.com/IskanderSh/vk-task/internal/lib/validation"
)

var (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func (s *UserService) AddUser(ctx context.Context, input *models.UserSignUp) error {
	const op = "service.AddUser"

	if !validator.Matches(*input.Email, validator.ValidEmail) {
		return wrapper.Wrap(op, ErrInvalidEmail)
	}

	if !validator.StringValueBetween(*input.Password, PasswordMinChars, PasswordMaxChars) {
		return wrapper.Wrap(op, ErrInvalidPassword)
	}

	if !validator.PermittedValue(*input.Role, validator.ValidRole...) {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), 12)
	if err != nil {
		return wrapper.Wrap(op, err)
	}

	user := entities.User{
		Email:    *input.Email,
		Password: string(hashPassword),
		Role:     *input.Role,
	}

	if err := s.storage.CreateUser(&user); err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, input *models.UserSignIn) (string, error) {
	const op = "service.Authenticate"

	if !validator.Matches(*input.Email, validator.ValidEmail) {
		return "", wrapper.Wrap(op, ErrInvalidEmail)
	}

	if !validator.StringValueBetween(*input.Password, PasswordMinChars, PasswordMaxChars) {
		return "", wrapper.Wrap(op, ErrInvalidPassword)
	}

	actor, err := s.storage.GetUser(*input.Email)
	if err != nil {
		return "", wrapper.Wrap(op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(actor.Password), []byte(*input.Password)); err != nil {
		return "", wrapper.Wrap(op, ErrInvalidPassword)
	}

	payload := jwt.MapClaims{
		"sub":  actor.Email,
		"role": actor.Role,
		"exp":  time.Now().Add(TokenExpirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(signingKey)
	if err != nil {
		return "", wrapper.Wrap(op, err)
	}

	return t, nil
}
