package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	port   = "8080"
	dbhost = "localhost"
	dbport = 5432
	dbuser = "postgres"
	dbpass = "postgres"
	dbname = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)

	db, err := ConnectDB(psqlconn)
	if err != nil {
		log.Println("error connecting to database:", err)
		return
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/create", Create(db)).Methods("POST")
	router.HandleFunc("/read/{id}", Read(db)).Methods("GET")

	log.Printf("listening on %s...\n", port)
	http.ListenAndServe(":"+port, router)
}
