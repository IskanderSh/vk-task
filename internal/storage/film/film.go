package film

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
	ErrDuplicateName = errors.New("films: duplicate name")
)

func (s *Storage) CreateFilm(film *entities.Film) error {
	const op = "storage.film.CreateFilm"

	_, err := s.db.Exec(createFilmQuery, film.Name, film.Description, film.Date, film.Rating, film.Actors)
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

func (s *Storage) GetFilm(name string) (*entities.Film, error) {
	const op = "storage.film.GetFilm"

	var film entities.Film

	row := s.db.QueryRow(getFilmQuery, name)
	if err := row.Scan(&film.ID, &film.Name, &film.Description,
		&film.Date, &film.Rating, &film.Actors, &film.CreatedAt); err != nil {
		return nil, wrapper.Wrap(op, err)
	}

	return &film, nil
}

func (s *Storage) UpdateFilm(film *entities.UpdateFilm) error {
	const op = "storage.film.UpdateFilm"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if film.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *film.Description)
		argId++
	}

	if film.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *film.Date)
		argId++
	}

	if film.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *film.Rating)
		argId++
	}

	if film.Actors != nil {
		setValues = append(setValues, fmt.Sprintf("actors=$%d", argId))
		args = append(args, *film.Actors)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("%s %s WHERE name=$%d", updateFilmQuery, setQuery, argId)
	args = append(args, film.Name)
	fmt.Println(query)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}

func (s *Storage) DeleteFilm(name string) error {
	const op = "storage.film.DeleteFilm"

	_, err := s.db.Exec(deleteFilmQuery, name)
	if err != nil {
		return wrapper.Wrap(op, err)
	}

	return nil
}
