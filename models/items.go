package models

// Item ...
type Item struct {
	ID          uint
	ItemOrderID uint

	ItemCategory   ItemCategory
	ItemCategoryID uint

	Name         string  `sql:"unique;not null"`
	UnitPrice    float64 `sql:"not null"`
	ItemPhotoURI string  `sql:"default:'Image Not Available'"`
	Description  string  `sql:"default:'No Description Available'"`
}
