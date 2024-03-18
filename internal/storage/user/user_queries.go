package user

const createUserQuery = `INSERT INTO users (email, password, role)
	VALUES ($1, $2, $3)`

const getUserQuery = `SELECT * FROM users WHERE email = $1`
