package main

import (
	"fmt"
	"math"
)

type calculate interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.diameter/2, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func main() {

	var bangunDatar calculate = lingkaran{10.0}
	fmt.Println(bangunDatar.luas())
	fmt.Println(bangunDatar.keliling())

	bangunDatar = persegi{10.0}
	fmt.Println(bangunDatar.luas())
	fmt.Println(bangunDatar.keliling())

}
