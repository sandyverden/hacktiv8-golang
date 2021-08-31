package main

import "fmt"

type Student struct {
	name  string
	grade int
}

func main() {

	var sandi Student
	sandi.name = "sandi"
	sandi.grade = 1

	fmt.Println("Name: ", sandi.name, "\nGrade:", sandi.grade)

	var budi = Student{
		name:  "Budi",
		grade: 1,
	}

	fmt.Println("Name: ", budi.name, "\nGrade:", budi.grade)

	// struct as slice

	var allPerson = []Student{
		{name: "Ali", grade: 1},
		{name: "Sandi", grade: 2},
	}

	fmt.Println(allPerson)
}
