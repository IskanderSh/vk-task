package storage

const createActorQuery = `INSERT INTO actors (name, sex, birthday) 
			VALUES ($1, $2, $3) RETURNING id`
