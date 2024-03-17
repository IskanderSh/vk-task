package storage

import (
	"database/sql"
	"fmt"

	"github.com/IskanderSh/vk-task/internal/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewStorage(cfg *config.Storage) (*sql.DB, error) {
	const op = "storage.NewStorage"

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, op)
	}

	return db, nil
}
