package main

import (
	"fmt"
	"sync"
)

func printNumber(mode string, wg *sync.WaitGroup) {
	for i := 0; i < 100000000000; i++ {
		//fmt.Println(mode, i)
	}
	fmt.Println(mode)
	wg.Done()
}

func printNumberBiasa(mode string) {
	for i := 0; i < 100000000000; i++ {
		//fmt.Println(mode, i)
	}
	fmt.Println(mode)
}

func main() {
	//printNumber("Biasa")
	//time.Sleep(1 * time.Second)

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go printNumber("Goroutine", &wg)
	//wg.Wait()

	printNumberBiasa("biasa")
}
