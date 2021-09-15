package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	printMessage := func(name string) {
		data := fmt.Sprintf("Hello %s", name)
		ch <- data // write data to channel
	}

	go printMessage("Andika")
	go printMessage("Ali")
	go printMessage("Lia")

	//message := <-ch // read data from channel
	fmt.Println(<-ch)

	message1 := <-ch // read data from channel
	fmt.Println(message1)

	message2 := <-ch // read data from channel
	fmt.Println(message2)

}
