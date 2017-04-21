package models

import "time"

// Bill ...
type Bill struct {
	ID        uint
	CreatedAt time.Time `sql:"default:NOW()"`
	UpdatedAt time.Time
	// One-one relationships
	// It embeds all the field of the respective tables including id

	Table   Table
	TableID uint

	Employee   Employee
	EmployeeID uint

	VAT   VAT
	VATID uint

	ServiceCh   ServiceCh
	ServiceChID uint

	// Many-to-one relationships

	ItemOrders []ItemOrder

	Payments []Payment

	// One-to-many relationships
	CustomerID uint

	Delivery    bool    `gorm:"default:false"`
	TotalPrice  float64 `sql:"not null"`
	Description string  `sql:"default:'Description Not Available'"`
}
