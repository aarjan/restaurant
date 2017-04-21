package models

// ItemCategory ...
type ItemCategory struct {
	ID    uint
	Name  string `sql:"unique;not null"`
	Extra string `sql:"default:'Not Available'"`
}
