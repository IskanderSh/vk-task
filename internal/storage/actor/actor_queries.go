package actor

const createActorQuery = `INSERT INTO actors (name, sex, birthday) 
			VALUES ($1, $2, $3)`

const getActorQuery = `SELECT * FROM actors WHERE name=$1`

const updateActorQuery = `UPDATE actors SET`

const deleteActorQuery = `DELETE FROM actors WHERE name=$1`
