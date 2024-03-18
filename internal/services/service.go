package services

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/IskanderSh/vk-task/internal/entities"
)

type ActorService struct {
	log     *slog.Logger
	storage ActorStorage
}

type ActorStorage interface {
	CreateActor(actor *entities.Actor) error
	GetActor(name string) (*entities.Actor, error)
	UpdateActor(actor *entities.UpdateActor) error
	DeleteActor(name string) error
}

type FilmService struct {
	log     *slog.Logger
	storage FilmStorage
}

type FilmStorage interface {
	CreateFilm(film *entities.Film) error
	GetFilm(name string) (*entities.Film, error)
	UpdateFilm(film *entities.UpdateFilm) error
	DeleteFilm(name string) error
}

type UserService struct {
	log     *slog.Logger
	storage UserStorage
}

type UserStorage interface {
	CreateUser(user *entities.User) error
	GetUser(email string) (*entities.User, error)
}

var (
	PasswordMinChars    = 4
	PasswordMaxChars    = 40
	TokenExpirationTime = time.Hour * 720

	ErrDuplicateName      = errors.New("actor with this name already exists")
	ErrDuplicateEmail     = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidEmail       = errors.New("incorrect email")
	ErrInvalidPassword    = errors.New(fmt.Sprintf("password should be more than %d symbols and less than %d",
		PasswordMinChars, PasswordMaxChars))
	ErrUserNotFound  = errors.New("no user with this email")
	ErrActorNotFound = errors.New("no actor with this name")
	ErrFilmNotFound  = errors.New("no film with this name")
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
