package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investment float64
	var years float64
	var expectedReturn float64
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investment)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturn)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	futureValue := investment * math.Pow((1+expectedReturn/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println((futureRealValue))
}
