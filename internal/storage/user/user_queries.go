package user

const createUserQuery = `INSERT INTO users (email, password, role)
	VALUES ($1, $2, $3)`
