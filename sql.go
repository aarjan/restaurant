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

	membership := models.Membership{
		// get CustomerID
		MembershipType: "Exclusive",
		Discount:       0.2,
		Description:    "Full featured access",
	}

	customer1 := models.Customer{
		FirstName:     "Aarjan",
		LastName:      "Baskota",
		Address:       "Balkot",
		Age:           23,
		Email:         "baskotaaarjan@gmail.com",
		ContactNumber: "9849321326",
		Memberships:   []models.Membership{membership},
		Gender:        "Male",
	}

	itemCategory1 := models.ItemCategory{
		Name: "Snacks",
	}
	itemCategory2 := models.ItemCategory{
		Name: "Soop",
	}

	item1 := models.Item{
		Name:         "Sandwitch",
		Description:  "Cheese and Grilled",
		ItemCategory: itemCategory1,
		UnitPrice:    12.233,
	}

	item2 := models.Item{
		Name:         "Veg Soop",
		Description:  "With Cheese and Honey",
		ItemCategory: itemCategory2,
		UnitPrice:    34.23,
	}

	itemOrder := models.ItemOrder{
		Items:          []models.Item{item1, item2},
		Quantity:       3,
		OrderPrice:     234.2,
		OrderTimeStamp: time.Now(),
	}

	serviceCh := models.ServiceCh{
		Value: 23.23,
		Name:  "Service charge",
	}
	table1 := models.Table{
		Alias:    "Table 004",
		Capacity: 6,
		Status:   models.Ordered,
	}

	payment := models.Payment{
		PaymentType: "Visa DebitCard",
	}

	position := models.Position{
		Name:        "Cook",
		Description: "2nd Standard",
	}

	employee1 := models.Employee{
		FirstName:     "Ram",
		LastName:      "Thapa",
		Address:       "Dillibazar",
		DOB:           "1992-3-23",
		Email:         "",
		Gender:        "Male",
		ContactNumber: "9832899483",
		JoinDate:      "2073-4-31",
		Salary:        23000,
		Positions:     []models.Position{position},
	}

	employee2 := models.Employee{
		FirstName:     "Shiva",
		LastName:      "Thapa",
		Address:       "Bagbazar",
		DOB:           "1994-3-23",
		Email:         "",
		Gender:        "Male",
		ContactNumber: "9832379483",
		JoinDate:      "2071-3-31",
		Salary:        20000,
	}
	leave := models.Leave{
		EmployeeID: employee2.ID,
		Status:     true,

		Date: fmt.Sprint(time.Now().Date()),
	}

	vat := models.VAT{
		Value: 0.13,
		Name:  "Value Added Tax",
	}

	bill := models.Bill{
		CreatedAt: time.Now(),
		Delivery:  false,
		Employee:  employee1,
		ServiceCh: serviceCh,

		ItemOrders: []models.ItemOrder{itemOrder},
		Payments:   []models.Payment{payment},
		Table:      table1,
		VAT:        vat,
	}

	db.Create(&bill)

	err = db.Table("employees").Create(&employee2).Error
	checkError(err)

	err = db.Table("customers").Create(&customer1).Error
	checkError(err)

	err = db.Table("customers").Create(&employee2).Error
	checkError(err)

	db.Where("id= ?", 1).First(&employee2)
	err = db.Table("leaves").Create(&leave).Update("EmployeeID", employee2.ID).Error
	checkError(err)

	db.Where("id= ?", 1).First(&customer1)
	err = db.Model(&bill).Update("customer_id", customer1.ID).Error
	checkError(err)

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
