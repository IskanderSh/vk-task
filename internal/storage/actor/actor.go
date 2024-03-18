package actor

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (s *Storage) CreateActor(actor *entities.Actor) error {
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

func (s *Storage) GetActor(name string) (*entities.Actor, error) {
	const op = "storage.actor.GetActor"

	var actor entities.Actor

	row := s.db.QueryRow(getActorQuery, name)
	if err := row.Scan(&actor.ID, &actor.Name, &actor.Sex, &actor.Birthday, &actor.CreatedAt); err != nil {
		return nil, wrapper.Wrap(op, err)
	}

	return &actor, nil
}

func (s *Storage) UpdateActor(actor *entities.UpdateActor) error {
	const op = "storage.actor.UpdateActor"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if actor.Sex != nil {
		setValues = append(setValues, fmt.Sprintf("sex=$%d", argId))
		args = append(args, *actor.Sex)
		argId++
	}

	if actor.Birthday != nil {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, *actor.Birthday)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("%s %s WHERE name=$%d", updateActorQuery, setQuery, argId)
	args = append(args, actor.Name)
	fmt.Println(query)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}

func (s *Storage) DeleteActor(name string) error {
	const op = "storage.actor.DeleteActor"

	_, err := s.db.Exec(deleteActorQuery, name)
	if err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}
