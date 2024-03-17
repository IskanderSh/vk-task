package services

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/IskanderSh/vk-task/internal/entities"
)

type ActorService struct {
	log     *slog.Logger
	storage ActorStorage
}

type ActorStorage interface {
	CreateActor(actor *entities.CreateActor) error
}

type FilmService struct {
	log     *slog.Logger
	storage FilmStorage
}

type FilmStorage interface {
}

type UserService struct {
	log     *slog.Logger
	storage UserStorage
}

type UserStorage interface {
	CreateUser(user *entities.CreateUser) error
}

var (
	PasswordMinChars = 4
	PasswordMaxChars = 40

	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidEmail       = errors.New("incorrect email")
	ErrInvalidPassword    = errors.New(fmt.Sprintf("password should be more than %d symbols and less than %d",
		PasswordMinChars, PasswordMaxChars))
)

func NewActorService(log *slog.Logger, storage ActorStorage) *ActorService {
	return &ActorService{
		log:     log,
		storage: storage,
	}
}

func NewFilmService(log *slog.Logger, storage FilmStorage) *FilmService {
	return &FilmService{
		log:     log,
		storage: storage,
	}
}

func NewUserService(log *slog.Logger, storage UserStorage) *UserService {
	return &UserService{
		log:     log,
		storage: storage,
	}
}
