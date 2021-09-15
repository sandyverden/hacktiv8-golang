package main

import "fmt"

//sq.Query

func SqlQuery() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	age := 30
	row, err := db.Query("select id, name, grade, age from Student where age =?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer row.Close()

	var result []Student
	for row.Next() {
		var student = Student{}
		err := row.Scan(&student.id, &student.name, &student.age, &student.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, student)
	}

	if err = row.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}
