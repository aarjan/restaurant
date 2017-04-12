package main

import (
	"fmt"
	"log"

	"github.com/aarjan/restaurant/models"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/restaurant?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customer1 := models.Customer{
		FirstName: "Aarjan",
		LastName:  "Baskota",
		Address:   "Balkot",
		Email:     "baskotaaarjan@gmail.com",
		Phone:     9849321326,
		Gender:    "Male",
	}

	itemCategory1 := models.ItemCategory{
		ID:   777,
		Name: "Snacks",
	}
	itemCategory2 := models.ItemCategory{
		ID:   737,
		Name: "Soop",
	}

	item1 := models.Item{
		Name:           "Sandwitch",
		Description:    "Cheese and Grilled",
		ItemCategoryID: itemCategory1.ID,
	}

	item2 := models.Item{
		ID:             0x000002,
		Name:           "Veg Soop",
		Description:    "With Cheese and Honey",
		ItemCategoryID: itemCategory2.ID,
	}

	itemOrder := models.ItemOrder{
		Items:          []models.Item{item1, item2},
		Quantity:       3,
		OrderTimeStamp: time.Now(),
	}

	membership := models.Membership{
		// get CustomerID
		MembershipType: 2222,
		Discount:       0.2,
		Description:    "Exclusive",
	}
	serviceCh := models.ServiceCh{
		ID:           001,
		ServiceValue: 23.23,
	}
	table1 := models.Table{
		TableNo:  004,
		Capacity: 6,
		Status:   true,
	}

	payment := models.Payment{
		ID:          0,
		PaymentType: "DebitCard",
		Description: "Visa",
	}

	position := models.Position{
		ID:          0,
		EmployeeID:  0123,
		Name:        "Cook",
		Description: "2nd Standard",
	}

	employee1 := models.Employee{
		ID:        0123,
		FirstName: "Ram",
		LastName:  "Thapa",
		Address:   "Dillibazar",
		DOB:       "1992-3-23",
		Email:     "",
		Gender:    "Male",
		Phone:     9832899483,
		Photo:     "",
		JoinDate:  "2073-4-31",
		Salary:    23000,
	}

	employee2 := models.Employee{
		ID:        0134,
		FirstName: "Shiva",
		LastName:  "Thapa",
		Address:   "Bagbazar",
		DOB:       "1994-3-23",
		Email:     "",
		Gender:    "Male",
		Phone:     9832379483,
		Photo:     "",
		JoinDate:  "2071-3-31",
		Salary:    20000,
	}
	leave := models.Leave{
		ID:         001,
		EmployeeID: employee2.ID,
		Status:     true,

		Date: fmt.Sprint(time.Now().Date()),
	}

	vat := models.VAT{
		ID:    01,
		Value: 0.13,
	}

	bill := models.Bill{
		CreatedAt:   time.Now(),
		Delivery:    false,
		EmployeeID:  employee1.ID,
		CustomerID:  customer1.ID,
		ServiceChID: serviceCh.ID,
		ItemOrders:  []models.ItemOrder{itemOrder},
		Payments:    []models.Payment{payment},
		TableID:     table1.TableNo,
	}
	db.Model(&bill)
	// db.CreateTable(&position)
	// db.CreateTable(&leave)
	// db.CreateTable(&customer1)
	// db.CreateTable(&payment)
	// db.CreateTable(&serviceCh)
	// db.CreateTable(&membership)
	// db.CreateTable(&table1)
	// db.CreateTable(&item1)
	// db.CreateTable(&itemCategory1)
	// db.CreateTable(&itemOrder)
	// db.CreateTable(&bill)
	// db.CreateTable(&vat)

}
