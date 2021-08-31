package main

import "fmt"

func main() {

	nilai := 80
	if nilai >= 90 {
		fmt.Println("Nilai dalam hurup: A")
	} else if nilai >= 80 {
		fmt.Println("Nilai dalam hurup: B")
	} else {
		fmt.Println("Nilai dalam hurup: C")
	}

	if grade := 85; grade >= 90 {
		fmt.Println(grade, "get A")
	} else if grade >= 80 {
		fmt.Println(grade, "Get B")
	}

	num := 3 //delapan

	switch num {
	case 1, 2, 3, 4:
		fmt.Println("satu")
	case 5, 6, 7, 8:
		fmt.Println("Dua")
	default:
		fmt.Println("Lebih dari Dua")
	}

	rank := 80 //delapan
	switch {
	case rank >= 89:
		fmt.Println("satu")
	case rank >= 80:
		fmt.Println("Dua")
	default:
		fmt.Println("Lebih dari Dua")
	}
}
