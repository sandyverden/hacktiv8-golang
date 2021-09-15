package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	go func() {
		for {
			name := <-ch
			fmt.Println("Name ", name)
		}
	}()

	for i := 0; i < 100; i++ {
		data := fmt.Sprintf("Murid %v", i)
		ch <- data

	}
}
