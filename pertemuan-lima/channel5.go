package main

import "fmt"

func main() {
	numbers := []int{1, 434, 4, 5, 5, 5, 6, 6, 424, 5, 65, 7, 4324, 6, 5345, 65, 5}

	maxCh := make(chan int)
	avgCh := make(chan float64)

	go getMax(numbers, maxCh)
	go getAvg(numbers, avgCh)

	select {
	case max := <-maxCh:
		fmt.Println("max number: ", max)
	case avg := <-avgCh:
		fmt.Printf("Avg number: %.2f", avg)
	}

}

func getMax(numbers []int, ch chan int) {
	maxNum := numbers[0]

	for _, v := range numbers {
		if maxNum < v {
			maxNum = v
		}
	}

	ch <- maxNum
}

func getAvg(numbers []int, ch chan float64) {
	total := 0
	for _, v := range numbers {
		total += v
	}

	avgNum := float64(total) / float64(len(numbers))

	ch <- avgNum
}
