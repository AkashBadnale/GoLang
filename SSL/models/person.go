package models

import "time"

type Person struct {
	ID         uint
	FirstName  string
	MiddleName string
	LastName   string
	DOB        time.Time
	Mobile     int
	Email      string
	created    time.Time
}
