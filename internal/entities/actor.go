package entities

import "time"

type Actor struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Sex       string    `db:"sex"`
	Birthday  time.Time `db:"birthday"`
	CreatedAt time.Time `db:"created_at"`
}
