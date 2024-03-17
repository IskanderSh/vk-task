package entities

import "time"

type CreateActor struct {
	Name     string    `db:"name"`
	Sex      string    `db:"sex"`
	Birthday time.Time `db:"birthday"`
}
