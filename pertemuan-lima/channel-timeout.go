package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	ch := make(chan int)
	go sendData(ch)
	retrieveData(ch)
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		exeTime := rand.Int()%10 + 1
		fmt.Println("Time Execution", exeTime)
	}
}
