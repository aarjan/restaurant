package models

import (
	"time"
)

// ItemOrder ...
type ItemOrder struct {
	ID     uint
	BillID uint

	Items []Item

	Quantity       int
	OrderPrice     int
	OrderTimeStamp time.Time
}
