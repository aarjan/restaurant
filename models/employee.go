package models

// Employee ...
type Employee struct {
	ID uint

	Positions []Position

	FirstName     string `sql:"not null"`
	MiddleName    string
	LastName      string `sql:"not null"`
	DOB           string `sql:"not null;date"`
	JoinDate      string `sql:"date;default:'DATE()'"`
	Salary        uint   `sql:"not null"`
	ContactNumber string `sql:"unique;not null"`
	Gender        string `sql:"not null"`
	Address       string `sql:"default:'No Address Available'"`
	Photo         string `sql:"default:'No Photo Available'"`
	Email         string `sql:"default:'No Email Available'"`
}
