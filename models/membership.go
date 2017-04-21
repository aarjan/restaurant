package models

// Membership ...
type Membership struct {
	ID         uint
	CustomerID uint

	MembershipType string  `sql:"unique;not null"`
	Discount       float64 `sql:"not null;default:0"`
	Description    string  `sql:"default:'Description not available.'"`
}
