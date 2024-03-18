package entities

import (
	"time"

	"github.com/IskanderSh/vk-task/internal/generated/models"
	"github.com/go-openapi/strfmt"
)

type Actor struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Sex       string    `db:"sex"`
	Birthday  time.Time `db:"birthday"`
	CreatedAt time.Time `db:"created_at"`
}

type UpdateActor struct {
	Name     string
	Sex      *string
	Birthday *time.Time
}

func ConvertToInput(a Actor) models.Actor {
	tm := strfmt.Date(a.Birthday)
	return models.Actor{
		Name:     &a.Name,
		Sex:      &a.Sex,
		Birthday: &tm,
	}
}
