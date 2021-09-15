package main

import "fmt"

type person struct {
	name  string
	grade int
}

func main() {
	var unknowntype interface{}
	fmt.Println(unknowntype)

	unknowntype = "testaja"
	fmt.Println(unknowntype)

	unknowntype = 2
	fmt.Println(unknowntype)

	data := map[string]interface{}{
		"name":   "sandi",
		"bahasa": "Go",
		"grade":  10,
	}

	fmt.Println(data)

	// casting
	var unknownStruct interface{} = &person{name: "sandi", grade: 10}
	fmt.Println(unknownStruct.(*person).name)

	// casting slice
	var unknownSlice interface{} = []string{"a", "b", "c", "d"}
	fmt.Println(unknownSlice.([]string)[1])
}
