package models

import (
	"database/sql"
	"time"
)

type UriLog struct {
	Timestamp Time
	Uri string
	Description sql.NullString
	Id int
}

func AllLogs(db *sql.DB) ([]*Log, error) {
	rows, err := db.Query("SELECT * FROM log")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := make([]*Log, 0)
	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&log.Id, &log.Uri, &log.Description, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}