package controllers

import (
	"assignment2/config"
	"assignment2/structs"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) CreateOrders(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var order structs.Orders
	json.NewDecoder(c.Request.Body).Decode(&order)
	json.NewEncoder(c.Writer).Encode(order)
	orderID := config.InitID()
	Conn.DB.Exec("insert into orders (order_id, customer_name, ordered_at) values (?, ?,?)", orderID,
		order.CustomerName, order.OrderAt)
	Conn.DB.Exec("insert into items (item_code, description, quantity, order_id) values (?, ?, ?, ?)", order.Items[0].ItemCode, order.Items[0].Description, order.Items[0].Quantity, orderID)
}

func (Conn *DBConn) GetOrders(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var order structs.Orders
	Conn.DB.Raw("select customer_name, ordered_at, item_code, description, quantity from orders, items limit 1").Scan(&order)
	log.Println(order)
	json.NewEncoder(c.Writer).Encode(order)
	return
}

func (Conn *DBConn) UpdateOrder(c *gin.Context) {
	//var order structs.Orders
	//var result gin.H

	//order.order = c.PostForm("order")
}

func (Conn *DBConn) DeleteOrders(c *gin.Context) {
	//var order structs.Orders
	//var result gin.H

	//order.order = c.PostForm("order")
}
