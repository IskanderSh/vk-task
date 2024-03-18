package film

const createFilmQuery = `INSERT INTO films (name, description, date, rating, actors) 
			VALUES ($1, $2, $3, $4, $5)`

const getFilmQuery = `SELECT * FROM films WHERE name=$1`

const updateFilmQuery = `UPDATE films SET`

const deleteFilmQuery = `DELETE FROM films WHERE name=$1`
