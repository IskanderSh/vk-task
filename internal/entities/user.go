package entities

type CreateUser struct {
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
