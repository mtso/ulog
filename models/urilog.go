package models

import (
	"database/sql"
	"time"
)

type UriLog struct {
	Time Date
	Uri string
	Description string
	Tags []string
}