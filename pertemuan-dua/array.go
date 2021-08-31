package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 10
	x[3] = 30
	fmt.Println(x)

	// literal
	y := [3]int{1, 2, 3}
	fmt.Println(y)

	y[1] = 10
	fmt.Println(y)

	// style horizontal
	fruits := [5]string{"mangga", "apel", 3: "anggur", "pisang"}
	fruits[2] = "jeruk"
	fmt.Println(fruits)

	// style vertical
	buahBuahan := [5]string{
		"mangga",
		"apel",
		"jerik",
		"anggur",
		"pisang",
	}
	fmt.Println(buahBuahan)

	// jika jumlah index array belum diketahui, gunakan `...`
	animals := [...]string{
		"gajah",
		"kuda",
	}
	fmt.Println(animals)

}
