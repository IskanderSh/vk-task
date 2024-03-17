package services

import (
	"context"
	"time"

	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	"github.com/go-openapi/strfmt"
)

func (s *ActorService) AddActor(ctx context.Context, actor *models.Actor) error {
	const op = "service.AddActor"

	if ok := validateSex(*actor.Sex); !ok {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	tm, err := validateTime(actor.Birthday)
	if err != nil {
		return wrapper.Wrap(op, ErrInvalidCredentials)
	}

	if err := s.storage.CreateActor(*actor.Name, *actor.Sex, *tm); err != nil {
		return wrapper.Wrap(op, err)
	}
	return nil
}

func validateSex(sex string) bool {
	for _, value := range validSex {
		if sex == value {
			return true
		}
	}

	return false
}

func validateTime(check *strfmt.Date) (*time.Time, error) {
	now := time.Now()

	strTime := check.String()

	tm, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		return nil, err
	}

	if tm.After(now) {
		return nil, ErrIncorrectTime
	}

	return &tm, nil
}
