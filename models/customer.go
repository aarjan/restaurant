package models

// Customer ...
type Customer struct {
	ID          uint
	Bills       []Bill
	Memberships []Membership

	FirstName string
	LastName  string
	Phone     int
	Address   string
	Gender    string
	Age       int
	Email     string
}
