package models

// Position ...
type Position struct {
	ID          uint
	EmployeeID  uint
	Name        string `sql:"unique;not null"`
	Description string `sql:"Description not Available"`
}
