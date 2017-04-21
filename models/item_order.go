package models

import (
	"time"
)

// ItemOrder ...
type ItemOrder struct {
	ID     uint
	BillID uint

	Items []Item

	Quantity       int       `sql:"not null;default:1"`
	OrderPrice     float64   `sql:"not null"`
	OrderTimeStamp time.Time `sql:"default:NOW()"`
}
