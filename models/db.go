package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	queryTableCreateIfNotExist = `CREATE TABLE IF NOT EXISTS log (
		log_id bigserial PRIMARY KEY,
		log_uri text NOT NULL,
		log_description text,
		log_timestamp timestamp WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
	);`
)

func InitDB(databaseType, databaseUrl string) (*sql.DB, error) {
	// Open database type with url
	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Create `log` table if it doesn't exist
	if _, error := db.Exec(queryTableCreateIfNotExist); error != nil {
		return nil, error
	}

	return db, nil
}
