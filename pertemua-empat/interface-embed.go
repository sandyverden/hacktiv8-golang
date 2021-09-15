package main

import (
	"fmt"
	"math"
)

type calculate interface {
	calculate2d
	calculate3d
}

type calculate2d interface {
	luas() float64
}

type calculate3d interface {
	volume() float64
}

type kubus struct {
	sisi float64
}

func (k kubus) luas() float64 {
	return (k.sisi * k.sisi) * 8
}

func (k kubus) volume() float64 {
	//return k.sisi * k.sisi * k.sisi
	return math.Pow(k.sisi, 3)
}

func main() {
	var bangunDatar calculate = kubus{10}
	fmt.Printf("Luas Kubus : %.2f\n", bangunDatar.luas())
	fmt.Printf("Luas Kubus : %.2f\n", bangunDatar.volume())

}
