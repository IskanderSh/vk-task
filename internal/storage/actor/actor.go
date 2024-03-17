package actor

import (
	"database/sql"
	"errors"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/lib/error/wrapper"
	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

var (
	ErrDuplicateName = errors.New("actors: duplicate name")
)

func (s *Storage) CreateActor(actor *entities.CreateActor) error {
	const op = "storage.actor.CreateActor"

	_, err := s.db.Exec(createActorQuery, actor.Name, actor.Sex, actor.Birthday)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			if pqError.Code == "23505" {
				return ErrDuplicateName
			}
		}

		return wrapper.Wrap(op, err)
	}

	return nil
}
