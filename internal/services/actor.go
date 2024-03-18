package services

import (
	"context"
	"log/slog"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	"github.com/IskanderSh/vk-task/internal/lib/validation"
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
		return wrapper.Wrap(op, err)
	}
	return nil
}
