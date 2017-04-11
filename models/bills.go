package models

// Bill ...
type Bill struct {
	ID         int
	CustomerID int
	TableID    int
	EmployeeID string
	Delivery   bool
	PaymentID  int
	TotalPrice int
	Note       string
	VatID      int
	ServiceID  int
}
