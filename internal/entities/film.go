package entities

import "time"

type Film struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Date        time.Time `db:"date"`
	Rating      int       `db:"rating"`
	Actors      []string  `db:"actors"`
	CreatedAt   time.Time `db:"created_at"`
}

type UpdateFilm struct {
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	Date        *time.Time `db:"date"`
	Rating      *int       `db:"rating"`
	Actors      *[]string  `db:"actors"`
}
