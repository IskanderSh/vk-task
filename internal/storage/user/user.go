package user

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
	ErrDuplicateEmail = errors.New("users: duplicate email")
)

func (s *Storage) CreateUser(user *entities.CreateUser) error {
	const op = "storage.user.CreateUser"

	_, err := s.db.Exec(createUserQuery, user.Email, user.Password, user.Role)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			if pqError.Code == "23505" {
				return ErrDuplicateEmail
			}
		}

		return wrapper.Wrap(op, err)
	}

	return nil
}
