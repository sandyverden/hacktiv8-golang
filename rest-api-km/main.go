package main

import (
	"net/http"
	"rest-api-km/config"
	"rest-api-km/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.InitDB()
	DBConn := &controllers.DBConn{DB: db}
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})

	router.POST("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World POST\n")
	})

	// CRUD

	// Create
	router.POST("/person", DBConn.CreatePerson)
	// Read -- localhost:8081/persin/1 -- use Param
	router.GET("/person/:id", DBConn.GetPersonByID)
	router.GET("/persons", DBConn.GetPersons)
	// Update -- localhost:8081/person?id=1 -- Use Query
	router.PUT("/person", DBConn.UpdatePerson)
	// Delete
	router.DELETE("/persin/:id", DBConn.DeletePerson)

	router.Run(":8081")
}
