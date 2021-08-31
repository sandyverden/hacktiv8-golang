package main

import "fmt"

func main() {
	x := []int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(x)

	// nil slice
	var y []int
	fmt.Println(y == nil)

	// fruits slice
	fruits := []string{
		"mangga",
		"pisang",
	}

	fmt.Printf("Element: %v, len: %v, cap: %v \n ", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "anggur")
	fmt.Printf("Element: %v, len: %v, cap: %v \n ", fruits, len(fruits), cap(fruits))

	fruits = append(fruits, "jambu")
	fmt.Printf("Element: %v, len: %v, cap: %v \n ", fruits, len(fruits), cap(fruits))

	// slice menambah cap (kapaasitas) menggunakan pangkat 2 (2, 4, 8, 16, 32)
	fruits = append(fruits, "jeruk")
	fmt.Printf("Element: %v, len: %v, cap: %v \n ", fruits, len(fruits), cap(fruits))

	// create slice using make function
	animals := make([]string, 5)
	fmt.Println(animals)

	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	// copy slice
	fruits8 := make([]string, 3)
	numCopy := copy(fruits8, fruits)
	fmt.Println(fruits8, numCopy)

	// slicing slice
	alphabet := []string{"a", "b", "c", "d", "e"}
	fmt.Println(alphabet)

	newAlphabet := alphabet[1:3]
	fmt.Println(newAlphabet)

	// slice share memory
	newAlphabet[1] = "ba"
	fmt.Println(alphabet)

	newAlphabet = append(newAlphabet, "z")
	fmt.Println(newAlphabet)

	// create map type data composite
	aMap := map[string][]string{
		"fruits":  {"mangga", "jambu"},
		"animals": {"gajah", "kodok", "ikan"},
	}
	fmt.Println(aMap)

}
