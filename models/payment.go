package models

// Payment ...
type Payment struct {
	ID     uint
	BillID uint

	PaymentType string
	Description string
}
