package main

import (
	"github.com/mtso/ulog/models"
	"database/sql"
	"log"
	"net/http"
	"fmt"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/log", env.retrieveLogs)
	log.Print("listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
	fmt.Println("GET /log")
	for _, log := range logs {
		fmt.Fprintf(w, "log_id=%v log_timestamp=%s\nlog_description=\"%s\"\nlog_uri=\"%s\"\n\n", log.Id, log.Timestamp, log.Description.String, log.Uri)
	}
}