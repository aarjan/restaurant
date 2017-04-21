package models

import (
	"time"
)

// Customer ...
type Customer struct {
	ID          uint
	Bills       []Bill
	Memberships []Membership

	FirstName     string
	MiddleName    string
	LastName      string
	ContactNumber string
	Address       string
	Gender        string
	Age           uint `sql:"check(age>5)"`
	Email         string
	JoinDate      time.Time `sql:"default:NOW()"`
}
