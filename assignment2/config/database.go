package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

func InitDB() *gorm.DB {
	connection := "root:@tcp(localhost:3306)/orders_by?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic("Failed connect to database")
	}
	return db
}

func InitID() xid.ID {
	Guid := xid.New()
	return Guid
}
