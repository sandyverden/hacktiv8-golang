package main

import (
	"encoding/json"
	"log"
	"net/url"
)

type User struct {
	FullName string `json:"name"`
	Age      int
	//Properti Property
}

type Property struct {
	friend string
}

func main() {
	urlString := "https://golang.com/search?keyword=golang&page=2"

	URL, err := url.Parse(urlString)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Full URL:", URL)
	log.Println("Schema:", URL.Scheme)
	log.Println("Host:", URL.Host)
	log.Println("Path:", URL.Path)
	log.Println("Query String:", URL.Query())

	// JSON
	// {"name":"ali", "age":21}

	// JSON to Object (Struct) -> Unmarshal
	// 1. JSON -> body json || file
	// 2. Transform string -> byte
	// 3. Unmarshal to ....

	jsonString := `{"name":"ali", "age":12 }`
	jsonData := []byte(jsonString)
	//log.Println(jsonData)
	var user User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.FullName, user.Age)

	// JSON to Map [Key:Value]
	var userMap map[string]interface{}
	err = json.Unmarshal(jsonData, &userMap)
	log.Println(userMap["name"])

	// JSON to interface
	var UserInterface interface{}
	err = json.Unmarshal(jsonData, &UserInterface)
	log.Println(UserInterface.(map[string]interface{})["name"])

	// JSON Array
	jsonStringArray := `[
							{"name":"ali", "age":12},
							{"name":"lia", "age":13},
							{"name":"ilo", "age":14}
						]`

	jsonStringArrayData := []byte(jsonStringArray)
	var userStructArray []User
	err = json.Unmarshal(jsonStringArrayData, &userStructArray)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(userStructArray[2].FullName, userStructArray[2].Age)

	// Object to String JSON
	var userObj = []User{
		{"Ali", 21},
		{"lia", 22},
	}

	jsonDataByte, err := json.Marshal(&userObj)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(jsonDataByte))
}
