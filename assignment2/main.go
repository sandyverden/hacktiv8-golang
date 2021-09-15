package main

import (
	"assignment2/config"
	"assignment2/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.InitDB()
	DBConn := &controllers.DBConn{DB: db}
	router := gin.Default()

	router.GET("/orders", func(c *gin.Context) {
		c.String(http.StatusOK, "Orders OK")
	})

	router.POST("/orders", DBConn.CreateOrders)

	router.Run(":8082")
}
