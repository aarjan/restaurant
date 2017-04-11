package models

// ItemOrder ...
type ItemOrder struct {
	ID             int
	ItemID         int
	BillID         int
	Quantity       int
	OrderPrice     int
	OrderTimeStamp int
}
