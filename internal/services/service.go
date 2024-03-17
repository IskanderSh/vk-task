package services

import (
	"errors"
	"log/slog"
	"time"
)

type ActorService struct {
	log     *slog.Logger
	storage ActorStorage
}

type ActorStorage interface {
	CreateActor(name, sex string, birthday time.Time) error
}

type FilmService struct {
	log     *slog.Logger
	storage FilmStorage
}

type FilmStorage interface {
}

var (
	validSex = []string{"male", "female"}

	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectTime      = errors.New("invalid time, it should be less then now")
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
