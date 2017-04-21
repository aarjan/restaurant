package models

// VAT ...
type VAT struct {
	ID    uint    `json:"id"`
	Value float64 `sql:"default:13" json:"value"`
	Name  string  `sql:"unique;not null" json:"name"`
}
