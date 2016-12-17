package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	queryTableCreate = `CREATE TABLE log (
		log_id bigserial PRIMARY KEY,
		log_uri text NOT NULL,
		log_description text,
		log_timestamp timestamp WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
	);`
	queryExists = `SELECT EXISTS ( 
		SELECT 1
		FROM information_schema.tables 
		WHERE table_name='log'
	);`
)

func InitDB(databaseType, databaseUrl string) (*sql.DB, error) {
	db, err:= sql.Open(databaseType, databaseUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Create log table if it doesn't exist
	if _, error := db.Query(queryExists); error != nil {
		db.Exec(queryTableCreate)
	}

	return db, nil
}