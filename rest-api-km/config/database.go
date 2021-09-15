package config

import (
	"fmt"
	"rest-api-km/structs"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	connection := "root:@tcp(localhost:3306)/golang?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic("Failed connect to database")
	}
	db.AutoMigrate(structs.Person{})
	return db
}
