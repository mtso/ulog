package models

import (
	"database/sql"
	"time"
)

const (
	queryAllDescending = "SELECT * FROM log ORDER BY log_id DESC"
	queryInsert        = "INSERT INTO log (log_uri, log_description) VALUES($1, $2)"
)

type UriLog struct {
	// Must use pointer for time.Time
	// From: https://golang.org/pkg/database/sql/#Rows.Scan
	// ```
	// Source values of type time.Time may be scanned into values
	// of type *time.Time, *interface{}, *string, or *[]byte. 
	// ```
	Timestamp   *time.Time
	Uri         string
	Description sql.NullString
	Id          int
}

func AllLogs(db *sql.DB) ([]*UriLog, error) {

	// Query log table
	rows, err := db.Query(queryAllDescending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map rows into UriLog objects
	logs := make([]*UriLog, 0)
	for rows.Next() {
		log := new(UriLog)
		err := rows.Scan(&log.Id, &log.Uri, &log.Description, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}

func CreateLog(db *sql.DB, uri string, description string) (int, error) {

	result, err := db.Exec(queryInsert, uri, description)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
