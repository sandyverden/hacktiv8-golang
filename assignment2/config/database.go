package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	connection := "root:@tcp(localhost:3306)/kampus_merdeka?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic("Failed connect to database")
	}
	return db
}
