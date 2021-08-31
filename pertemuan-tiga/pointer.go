package main

import "fmt"

func main() {

	var numberA int = 4
	var newA int = numberA
	var numberB *int = &numberA

	fmt.Println("NumberA value:", numberA, "\nNumberA pointer address:", &numberA)
	fmt.Println("NumberA value:", newA, "\nNumberA pointer address:", &newA)
	fmt.Println("NumberA value:", *numberB, "\nNumberA pointer address:", numberB)

	numberA = 5

	fmt.Println("NumberA value:", numberA, "\nNumberA pointer address:", &numberA)
	fmt.Println("NumberA value:", newA, "\nNumberA pointer address:", &newA)
	fmt.Println("NumberA value:", *numberB, "\nNumberA pointer address:", numberB)

	number := 10
	fmt.Println(number)

	change(&number, 100)
	fmt.Println(number)
}

func change(number *int, newNumber int) {
	*number = newNumber
}
