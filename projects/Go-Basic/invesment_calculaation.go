package main

import (
	"fmt"
	"math"
)

func main() {
	var invesment = 1000
	var expectedReturn = 5.5
	var years = 10

	var futureValue = float64(invesment) * math.Pow((1+expectedReturn/100), float64(years))
	fmt.Print(futureValue)
}
