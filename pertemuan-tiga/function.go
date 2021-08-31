package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	name := []string{"Kampus", "Merdeka"}
	printMessage(name)

	randomNum := randomWithRange(1, 10)
	fmt.Println("Random Number:", randomNum)

	luasLingkaran, kelilingLingkarang := circleCalculate(float64(randomNum))
	fmt.Printf("Luas Lingkaran: %.2f \nKeliling Lingkarang: %.2f \n", luasLingkaran, kelilingLingkarang)

	luasLingkaranPre, kelilingLingkarangPre := circleCalculatePre(float64(randomNum))
	fmt.Printf("Pre Luas Lingkaran: %.2f \nPre Keliling Lingkarang: %.2f \n", luasLingkaranPre, kelilingLingkarangPre)

	average := calculateAvg(10, 20, 30, 40, 50)
	fmt.Println("Nilai rata2 =", average)
}

func printMessage(input []string) {
	name := strings.Join(input, " ")
	//fmt.Println("hello world", input)
	fmt.Println("Hello world", name)
}

func randomWithRange(min, max int) int {
	randomNum := rand.Int()%(max-min+1) + min
	//fmt.Println(randomNum)
	return randomNum
}

func circleCalculate(d float64) (float64, float64) {
	luas := math.Pi * math.Pow(d/2, 2)
	keliling := math.Pi * d

	return luas, keliling
}

func circleCalculatePre(d float64) (luas, keliling float64) {
	luas = math.Pi * math.Pow(d/2, 2)
	keliling = math.Pi * d

	return
}

func calculateAvg(input ...int) int {
	result := 0
	for _, value := range input {
		result += value
	}
	avgResult := result / len(input)
	return avgResult
}
