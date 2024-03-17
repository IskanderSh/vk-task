package services

import (
	"context"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	"github.com/IskanderSh/vk-task/internal/lib/validation"
)

func (s *ActorService) AddActor(ctx context.Context, input *models.Actor) error {
	const op = "service.AddActor"

	if ok := validator.PermittedValue(*input.Sex, validator.ValidSex...); !ok {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	tm, err := validator.ParseTime(input.Birthday)
	if err != nil {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	actor := entities.CreateActor{
		Name:     *input.Name,
		Sex:      *input.Sex,
		Birthday: *tm,
	}

	if err := s.storage.CreateActor(&actor); err != nil {
		return wrapper.Wrap(op, err)
	}
	return nil
}
