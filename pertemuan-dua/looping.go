package main

import "fmt"

func main() {
	// for standard
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------------------")
	// for argument
	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("---------------------------")
	// for ...
	x := 0
	for {
		x++
		if x == 5 {
			continue
		}
		if x == 10 {
			break
		}
		fmt.Println(x)

	}

	alphabet := []string{"a", "b", "c", "d", "e", "f"}
	for i = 0; i < len(alphabet); i++ {
		fmt.Printf("alphabet ke %v adalah %v \n", i+1, alphabet[i])
	}
	fmt.Println("---------------------------")

	for index, huruf := range alphabet {
		fmt.Printf("alphabet ke %v adalah %v \n", index+1, huruf)
	}

	mapA := map[string]string{
		"fruit":  "mangga",
		"animal": "ular",
	}

	for k, v := range mapA {
		fmt.Printf("%v contohnya %v \n", k, v)
	}
}
