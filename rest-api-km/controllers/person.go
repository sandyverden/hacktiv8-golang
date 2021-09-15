package controllers

import (
	"net/http"
	"rest-api-km/structs"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) CreatePerson(c *gin.Context) {
	var person structs.Person
	var result gin.H

	person.FirstName = c.PostForm("first_name")
	person.LastName = c.PostForm("last_name")

	err := Conn.DB.Create(&person).Error
	if err != nil {

	}

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
	return
}

func (Conn *DBConn) GetPersonByID(c *gin.Context) {
	var person structs.Person
	var result gin.H
	id := c.Param("id")
	err := Conn.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) GetPersons(c *gin.Context) {

}

func (Conn *DBConn) UpdatePerson(c *gin.Context) {
	id := c.Query("id")

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data Not found",
		}
	}

	newPerson.FirstName = firstName
	newPerson.LastName = lastName

	err = Conn.DB.Model(&person).Update(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update Failed",
		}
	}

	result = gin.H{
		"result": "Update Data Success",
	}

	c.JSON(http.StatusOK, result)
	return
}

func (Conn *DBConn) DeletePerson(c *gin.Context) {

	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Delete Failed",
		}
		c.JSON(http.StatusOK, result)
	}
}
