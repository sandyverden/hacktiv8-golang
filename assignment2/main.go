package main

import (
	"assignment2/config"
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.InitDB()
	DBConn := &controllers.DBConn{DB: db}
	router := gin.Default()

	router.GET("/orders", DBConn.GetOrders)

	router.POST("/orders", DBConn.CreateOrders)

	router.PUT("/orders", DBConn.UpdateOrder)
	router.DELETE("/orders/:id", DBConn.DeleteOrders)

	router.Run(":8082")

}
