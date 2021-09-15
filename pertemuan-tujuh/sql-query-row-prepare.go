package main

import "fmt"

func SqlQueryRowPrepare() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	statement, err := db.Prepare("Select id, name, age, grade from studen where id=?", studentID)
}
