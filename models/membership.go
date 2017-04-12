package models

// Membership ...
type Membership struct {
	CustomerID uint

	MembershipType uint `gorm:"primary_key:true"`
	Discount       float64
	Description    string
}
