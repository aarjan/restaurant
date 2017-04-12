package models

// Employee ...
type Employee struct {
	ID uint

	Positions []Position

	FirstName string
	LastName  string
	DOB       string
	JoinDate  string
	Salary    uint
	Phone     uint
	Gender    string
	Address   string
	Photo     string
	Email     string
}
