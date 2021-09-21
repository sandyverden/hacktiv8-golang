package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "go-swagger/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Order struct {
	OrderId      string    `json:"order_id" example:"ABC"`
	Customername string    `json:"customer_name" example:"Ali"`
	OrderAt      time.Time `json:"order_at" example:"2006-01-02T15:04:05Z07:00"`
	Items        Item      `json:"items"`
}

type Item struct {
	ItemId      string `json:"item_id" example:"123"`
	Description string `json:"description" example:"Sepatu Nike"`
	Quantity    int    `json:"quantity" example:"1"`
}

var orders []Order
var prevOrderId = 0

// CreateOrder godoc
// @Summary Create a new order
// @Description Create new order with input on payload
// @Tags orders
// @Accept json
// @Product json
// @Param order body Order true "Create Order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderId++
	order.OrderId = strconv.Itoa(prevOrderId)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder godoc
// @Summary Get all order
// @Description Get all  order with input on payload
// @Tags orders
// @Accept json
// @Product json
// @Success 200 {object} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder godoc
// @Summary Get all order
// @Description Get all  order with input on payload
// @Tags orders
// @Accept json
// @Product json
// @Success 200 {object} Order
// @Router /orders/{orderId} [get]
// @Param orderId path int true "Id Order"
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderId == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderId == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderId == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// @title Kampus Merdeka API Docs
// @version 1.1
// @description This is a sample server for API Kampus Merdeka
// @contact.name sandyverden@gmail.com
// @host localhost:8083
// @bßßßßasePath /

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server Running on port 8083")
	http.ListenAndServe(":8083", router)

}
