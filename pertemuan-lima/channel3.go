package main

import "fmt"

func main() {
	ch := make(chan string, 2) // buffer channel

	ch <- "hello"
	fmt.Println(<-ch)

	ch <- "hello-2"
	ch <- "hello-3"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
