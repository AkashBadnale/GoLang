package main

import (
	"database/sql"
	"log"
)

// Person - database model
type Person struct {
	ID   string
	Name string
	Age  int
}

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("err sql.Open:", err)
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Println("db ping err:", err)
		return nil, err
	}

	log.Println("successfully connected to db...")

	query := `
            CREATE TABLE IF NOT EXISTS Persons (
                id SERIAL NOT NULL,
                name TEXT NOT NULL,
                age INTEGER NOT NULL,
                PRIMARY KEY (id)
            );`

	if _, err := db.Exec(query); err != nil {
		log.Println("error creating table:", err)
		return nil, err
	}

	return db, nil
}
