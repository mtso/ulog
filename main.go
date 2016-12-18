package main

import (
	"database/sql"
	"net/http"
	"encoding/json"
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

// PostResponse
type PostResponse struct {
	Success bool `json:"success"`
	RowsAffected int `json:"rows_affected"`
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
			return
		}

		// Allow anyone to GET log data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.URL)

		// Log and return json results
		log.Println("GET /log")
		jsonLogs, error := json.Marshal(logs)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonLogs)

	case m == "POST":
		// Parse parameter values
		uri := r.FormValue("uri")
		description := r.FormValue("description")

		// Execute on model
		rowsAffected, error := models.CreateLog(env.db, uri, description)
		if error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		// Log and return results
		log.Println("POST /log")
		jsonResponse := &PostResponse {
			Success: true,
			RowsAffected: rowsAffected,
		}

		// Create a json response
		success, error := json.Marshal(jsonResponse)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
		}
		w.Write(success)

	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
