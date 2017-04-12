package models

// Item ...
type Item struct {
	ID          uint
	ItemOrderID uint

	ItemCategory   ItemCategory
	ItemCategoryID uint

	Name        string
	UnitPrice   int
	ItemPhoto   string
	Description string
}
