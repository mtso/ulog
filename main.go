package main

import (
	"database/sql"
	"net/http"
	"fmt"
	"log"
	"os"

	"github.com/mtso/ulog/models"
)

const (
	// Localhost dummy db
	// Change these parameters to fit your postgresql instance
	db_local = "user=kingcandy password=cupcakes dbname=urilog sslmode=disable"
)

// Environment object holds a pointer to the database connection
type Env struct {
	db *sql.DB
}

func main() {
	// Determine database url
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = db_local
	}

	// Initialize database connection
	psql, err := models.InitDB("postgres", db_url)
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: psql}

	// Get port number
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // If port was not set as an environment variable
	}

	// Attach endpoint handler to /log route
	http.HandleFunc("/log", env.logEndpoint)

	// Log port number and start listening
	log.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Adds an endpoint function to the environment object
// Single /log endpoint for both GET and POST
func (env *Env) logEndpoint(w http.ResponseWriter, r *http.Request) {

	switch m := r.Method; {
	case m == "GET":
		// Execute on model
		logs, error := models.AllLogs(env.db)
		if error != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(error)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		// log.Println(r.URL.Host)
		// if r.URL.Host == "m-m.io" {
		// 	w.Header().Set("Access-Control-Allow-Origin", "*") //https://www.m-m.io")
		// }

		// Log and return results
		log.Println("GET /log")
		for _, log := range logs {
			fmt.Fprintf(w, "log_description=\"%s\"\nlog_id=%v log_timestamp=%s\nlog_uri=\"%s\"\n\n", log.Id, log.Timestamp, log.Description.String, log.Uri)
		}

	case m == "POST":
		// Parse parameter values
		uri := r.FormValue("uri")
		description := r.FormValue("description")

		// Execute on model
		rowsAffected, error := models.CreateLog(env.db, uri, description)
		if error != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(error)
			return
		}
		// Log and return results
		log.Println("POST /log")
		fmt.Fprintf(w, "Successfully created \"%s...\", %d row(s) affected.\n", description, rowsAffected)

	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
