package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitDB(databaseType, databaseUrl string) (*sql.DB, error) {
	db, err:= sql.Open(databaseType, databaseUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}