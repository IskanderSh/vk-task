package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	"github.com/IskanderSh/vk-task/internal/lib/validation"
	storage "github.com/IskanderSh/vk-task/internal/storage/actor"
)

func (s *ActorService) AddActor(ctx context.Context, input *models.Actor) error {
	const op = "service.AddActor"

	log := s.log.With(
		slog.String("op", op))

	if ok := validator.PermittedValue(*input.Sex, validator.ValidSex...); !ok {
		log.Error("incorrect sex param")
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	tm, err := validator.ParseTime(input.Birthday)
	if err != nil {
		log.Error("invalid time param")
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	actor := entities.Actor{
		Name:     *input.Name,
		Sex:      *input.Sex,
		Birthday: *tm,
	}

	if err := s.storage.CreateActor(&actor); err != nil {
		if errors.Is(err, storage.ErrDuplicateName) {
			return ErrDuplicateName
		}
		return wrapper.Wrap(op, err)
	}
	return nil
}

func (s *ActorService) UpdateActor(ctx context.Context, input *models.UpdateActor) error {
	const op = "service.actor.UpdateActor"

	log := s.log.With(slog.String("op", op))

	_, err := s.storage.GetActor(input.Name)
	if err != nil {
		return ErrActorNotFound
	}

	actor := entities.UpdateActor{
		Name: input.Name,
	}

	if input.Sex != "" {
		if ok := validator.PermittedValue(input.Sex, validator.ValidSex...); !ok {
			log.Error("incorrect sex param")
			return wrapper.Wrap(op, ErrInvalidCredentials)
		}
		actor.Sex = &input.Sex
	}

	if input.Birthday.String() != "" {
		tm, err := validator.ParseTime(&input.Birthday)
		if err != nil {
			log.Error("invalid time param")
			return wrapper.Wrap(op, ErrInvalidCredentials)
		}
		actor.Birthday = tm
	}

	if err := s.storage.UpdateActor(&actor); err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}
