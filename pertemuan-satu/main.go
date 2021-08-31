package main

import "fmt"

const (
	PHI      = 3.14
	LANGUAGE = "ID"
)

var (
	kendaraan = "motor"
	tahun     = 2014
)

func main() {
	fmt.Println("Hello World - 2")
	fmt.Println("Hello World - 1")

	var name string = "Sandi"
	fmt.Println(name)

	var tanggal = 24
	fmt.Println(tanggal)

	// variable dengan :=
	umur := 30
	fmt.Println(umur)

	// variable dengan menggunakan var multi value
	var kota, universitas string = "Jakarta", "UNTAR"
	fmt.Println(kota, universitas)

	// variable dengan menggunakan := multi value
	_, lastName, year := "Sandi", "Budiman", 2021

	fmt.Println(lastName, year)

	hacktiv := new(string)
	*hacktiv = "belajar golang"
	fmt.Println(*hacktiv)

	var formatInt int = 10
	var formatStr string = "Fundamental GO"
	var formatBool bool = true
	var formatFloat float32 = 10.5888

	fmt.Printf("formatted integer: %v \n Formated string: %v \n Formated bool: %v \n Formatted float32: %.1f\n", formatInt, formatStr, formatBool, formatFloat)

	fmt.Println(PHI, LANGUAGE)

	fmt.Println(kendaraan, tahun)

	weirdInt := 1_234
	fmt.Println(weirdInt)
}
