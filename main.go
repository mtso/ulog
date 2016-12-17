package main

import (
	"ulog/models"
	"database/sql"
	"log"
	"net/http"
)

const (
	db_source = "user=kingcandy password=cupcakes dbname=urilog sslmode=disable"
)

type Env struct {
	db *sql.DB
}

func main() {
	psql, err := models.InitDB(db_source)
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: psql}

	http.HandleFunc("/log", )
}

func (env *Env) retrieveLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	logs, error := models.AllLogs(env.db)
	if err != nil {
		http.Err(w, http.StatusText(500), 500)
		return
	}
	for _, log := range logs {
		fmt.Fprintf(w, "%s %s: %s; \"%s...\"", log.Id, log.Timestamp, log.Description, log.Uri[:15])
	}
}