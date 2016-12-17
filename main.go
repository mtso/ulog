package main

import (
	"ulog/models"
	"database/sql"
)

type Env struct {
	db *sql.DB
}