package main

import (
	"log"

	"fmt"

	"github.com/aarjan/restaurant/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@/restaurant?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vats := []models.VAT{}
	// db.Save(&vat)
	db.Find(&vats)
	fmt.Println(vats)
}
