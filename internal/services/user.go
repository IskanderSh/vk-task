package services

import (
	"context"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	validator "github.com/IskanderSh/vk-task/internal/lib/validation"
)

func (s *UserService) AddUser(ctx context.Context, input *models.UserSignUp) error {
	const op = "service.AddUser"

	if !validator.Matches(input.Email, validator.ValidEmail) {
		return wrapper.Wrap(op, ErrInvalidEmail)
	}

	if !validator.StringValueBetween(input.Password, PasswordMinChars, PasswordMaxChars) {
		return wrapper.Wrap(op, ErrInvalidPassword)
	}

	if !validator.PermittedValue(input.Role, validator.ValidRole...) {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	user := entities.CreateUser{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	if err := s.storage.CreateUser(&user); err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}
