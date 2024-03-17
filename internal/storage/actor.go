package storage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type ActorStorage struct {
	db *sql.DB
}

func NewActorStorage(db *sql.DB) *ActorStorage {
	return &ActorStorage{
		db: db,
	}
}

var (
	ErrDuplicateName = errors.New("actors: duplicate name")
)

func (s *ActorStorage) CreateActor(name, sex string, birthday time.Time) error {
	const op = "storage.actor.CreateActor"

	_, err := s.db.Exec(createActorQuery, name, sex, birthday)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			if pqError.Code == "23505" {
				return ErrDuplicateName
			}
		}

		return err
	}

	return nil
}
