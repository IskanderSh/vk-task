package services

import "log/slog"

type ActorService struct {
	log     *slog.Logger
	storage ActorStorage
}

type ActorStorage interface {
}

type FilmService struct {
	log     *slog.Logger
	storage FilmStorage
}

type FilmStorage interface {
}

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
