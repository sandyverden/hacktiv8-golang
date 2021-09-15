package main

import "fmt"

func main() {
	ch := make(chan string)

	names := []string{
		"Ali",
		"Lia",
		"Sandi",
	}
	for _, name := range names {
		go func(name string) {
			data := fmt.Sprintf("Hello %s", name)
			ch <- data
		}(name)
	}

	for i := 0; i < len(names); i++ {
		printMessage(ch)
	}
}

func printMessage(ch <-chan string) {
	fmt.Println(<-ch)
}
