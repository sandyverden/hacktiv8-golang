package main

import (
	"fmt"
	"reflect"
)

type person struct {
}

func main() {

	var message string = "hello world"

	reflectVal := reflect.ValueOf(message)
	fmt.Println("Type Data:", reflectVal.Type())
	fmt.Println("Value:", reflectVal.String())
	fmt.Println(reflect.TypeOf(message))

	person := person{}
	reflectStruct := reflect.ValueOf(person)
	fmt.Println(reflect.TypeOf(person))
	fmt.Println(reflectStruct.Kind())
}
