package main

import "fmt"

func main() {
	defer fmt.Println("Finish")
	fmt.Println("Start-1")
	//return

	orderSomeFood("nasi goreng")
	orderSomeFood("mie")
}

func orderSomeFood(food string) {
	defer fmt.Println("Finish")

	if food == "mie" {
		fmt.Println("--------------------")
		fmt.Println("And memesan", food)
		fmt.Println("--------------------")
		return
	}
}
