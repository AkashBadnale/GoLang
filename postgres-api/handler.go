package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := CreateRequest{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println("decode err:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			err := "name can not be empty"
			log.Println(err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}

		query := `insert into "persons"("name", "age") values($1, $2)`
		_, err = db.Exec(query, req.Name, req.Age)
		if err != nil {
			log.Println("error creating person:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("person created")
		w.WriteHeader(200)
	}
}

func Read(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			err := "id can not be empty"
			log.Println(err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}

		var person Person
		userSql := `SELECT * FROM persons WHERE id = $1`
		err := db.QueryRow(userSql, id).Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			log.Println("error executing query:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := ReadResponse{}
		res.Success = false
		res.User.ID = person.ID
		res.User.Name = person.Name
		res.User.Age = person.Age

		jsnStr, err := json.Marshal(res)
		if err != nil {
			log.Println("error unmarshalling ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsnStr)
	}
}
