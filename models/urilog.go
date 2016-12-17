package models

import (
	"database/sql"
	"time"
)

const (
	// ref time: Mon Jan 2 15:04:05 -0700 MST 2006
	timestamptzFormat = "2006-01-02T15:04:05.999999Z"

	queryAllDescending = "SELECT * FROM log ORDER BY log_id DESC"
	queryInsert        = "INSERT INTO log (log_uri, log_description) VALUES($1, $2)"
)

type UriLog struct {
	Timestamp   time.Time
	Uri         string
	Description sql.NullString
	Id          int
}

func AllLogs(db *sql.DB) ([]*UriLog, error) {

	rows, err := db.Query(queryAllDescending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := make([]*UriLog, 0)
	for rows.Next() {
		log := new(UriLog)
		rawtime := make([]byte, 0)
		// Place timestamp in byte buffer
		err := rows.Scan(&log.Id, &log.Uri, &log.Description, &rawtime)
		if err != nil {
			return nil, err
		}
		// Parse timestamp into time.Time type
		t, err := time.Parse(timestamptzFormat, string(rawtime))
		if err != nil {
			return nil, err
		}
		log.Timestamp = t
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
