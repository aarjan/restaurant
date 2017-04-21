package models

import "database/sql/driver"

// StatusType ...
type StatusType string

const (
	Ordered   StatusType = "order"
	Empty     StatusType = "empty"
	UnOrdered StatusType = "unorder"
	Served    StatusType = "served"
)

// Table ...
type Table struct {
	ID       uint
	Alias    string     `gorm:"unique"`
	Capacity int        `sql:"not null"`
	Status   StatusType `sql:"not null;type:ENUM('order','empty','unorder','served')"`
}

// Scan implemented from Scanner interface
func (s *StatusType) Scan(value interface{}) error {
	*s = StatusType(value.([]byte))
	return nil
}

// Value implemented from Valuer interface
func (s StatusType) Value() (driver.Value, error) {
	return string(s), nil
}
