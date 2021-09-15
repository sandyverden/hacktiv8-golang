package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input tidak boleh kosong")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi kesalahan", r)
	} else {
		fmt.Println("quit")
	}
}

func main() {
	defer catch()
	input := ""

	_, err := validate(input)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hello", input)
	}

	if valid, err := validate(input); valid {
		fmt.Println("hello", input)
	} else {
		panic(err.Error())
	}
	fmt.Scanln()
}
