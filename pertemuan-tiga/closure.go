package main

import (
	"fmt"
)

func main() {
	getMinMax := func(n []int) (int, int) {
		min, max := n[0], n[0]

		for _, v := range n {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
		}
		return min, max
	}

	numbers := []int{2, 5, 2, 6, 7, 3, 2}
	min, max := getMinMax(numbers)
	fmt.Println(min, max)

	newNumber := func(n int) []int {
		var result []int

		for _, v := range numbers {
			if v > n {
				continue
			}
			result = append(result, v)
		}
		return result
	}

	fmt.Println(newNumber(5))

	lenNumber, number := filterMax(numbers, 5)
	fmt.Printf("Jumlah Number: %v\nList Number %v\n", lenNumber, number())

}

func filterMax(n []int, max int) (int, func() []int) {
	var result []int
	for _, v := range n {
		if v > max {
			continue
		}
		result = append(result, v)
	}
	return len(result), func() []int {
		return result
	}
}
