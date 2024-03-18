package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	validator "github.com/IskanderSh/vk-task/internal/lib/validation"
	storage "github.com/IskanderSh/vk-task/internal/storage/film"
)

func (s *FilmService) AddFilm(ctx context.Context, input *models.Film) error {
	const op = "service.AddFilm"

	log := s.log.With(
		slog.String("op", op))

	if ok := validator.StringValueBetween(*input.Name, 4, 30); !ok {
		log.Error("incorrect name param")
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	if ok := validator.StringValueBetween(input.Description, 0, 1000); !ok {
		log.Error("incorrect description param")
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	tm, err := validator.ParseTime(input.Date)
	if err != nil {
		log.Error("invalid date param")
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	film := entities.Film{
		Name:        *input.Name,
		Description: input.Description,
		Date:        *tm,
		Rating:      int(*input.Rating),
		Actors:      input.Actors,
	}

	if err := s.storage.CreateFilm(&film); err != nil {
		if errors.Is(err, storage.ErrDuplicateName) {
			return ErrDuplicateName
		}
		return wrapper.Wrap(op, err)
	}
	return nil
}

func (s *FilmService) UpdateFilm(ctx context.Context, input *models.UpdateFilm) error {
	const op = "service.film.UpdateFilm"

	log := s.log.With(slog.String("op", op))

	_, err := s.storage.GetFilm(input.Name)
	if err != nil {
		return ErrFilmNotFound
	}

	film := entities.UpdateFilm{
		Name: input.Name,
	}

	if input.Description != "" {
		film.Description = &input.Description
	}

	if input.Date.String() != "" {
		tm, err := validator.ParseTime(&input.Date)
		if err != nil {
			log.Error("invalid time param")
			return wrapper.Wrap(op, ErrInvalidCredentials)
		}
		film.Date = tm
	}

	if input.Rating != nil {
		rt := int(*input.Rating)
		film.Rating = &rt
	}

	if input.Actors != nil {
		film.Actors = &input.Actors
	}

	if err := s.storage.UpdateFilm(&film); err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}

func (s *FilmService) DeleteFilm(ctx context.Context, name string) error {
	const op = "service.film.UpdateFilm"

	log := s.log.With(slog.String("op", op))

	_, err := s.storage.GetFilm(name)
	if err != nil {
		log.Warn("film with that name don't exists")
		return ErrFilmNotFound
	}

	if err := s.storage.DeleteFilm(name); err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}
