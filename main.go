package main

import (
	"github.com/mtso/ulog/models"
	"database/sql"
	"log"
	"net/http"
	"fmt"
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

	http.HandleFunc("/log", env.retrieveLogs)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) retrieveLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	logs, error := models.AllLogs(env.db)
	if error != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(error)
		return
	}
	for _, log := range logs {
		// fmt.Fprintf(w, "%v: %s \"%s...\"\n", log.Id, log.Description.String[:22], log.Uri[:15])
		fmt.Fprintf(w, "%v %s: %s; \"%s...\"\n", log.Id, log.Timestamp, log.Description.String, log.Uri[:15])
	}
}