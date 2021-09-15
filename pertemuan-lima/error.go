package main

import (
	"fmt"
	"strconv"
)

func main() {
	strNum := "1b"

	var err error
	var num int

	num, err = strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Input harus string number", err.Error())
	} else {
		fmt.Println("input anda", num)
	}

}
