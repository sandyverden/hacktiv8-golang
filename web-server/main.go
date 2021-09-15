package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var studentData = []student{
	{
		ID:    "1",
		Name:  "Ali",
		Grade: 9,
	},
	{
		ID:    "2",
		Name:  "Lia",
		Grade: 10,
	},
	{
		ID:    "3",
		Name:  "Didi",
		Grade: 11,
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/golang", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := map[string]string{
			"FirstName": "Ali",
			"LastName":  "Lia",
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index HTML")
}

func users(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result, err := json.Marshal(studentData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		fmt.Println(id)
		//result, err := json.Marshal(studentData)

		for _, v := range studentData {
			if v.ID == id {
				result, err := json.Marshal(v)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
