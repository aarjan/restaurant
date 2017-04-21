package models

// ServiceCh ...
type ServiceCh struct {
	ID    uint
	Name  string  `sql:"unique"`
	Value float64 `sql:"default:10.00"`
}
