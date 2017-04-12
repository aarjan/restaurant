package models

// Table ...
type Table struct {
	TableNo  uint `gorm:"primary_key:true"`
	Capacity int
	Status   bool
}
