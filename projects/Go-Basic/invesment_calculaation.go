package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investment float64
	var years float64
	expectedReturn := 2.5
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investment)
	// fmt.Print("Expected Return Rate: ")
	// fmt.Scan(&expectedReturn)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := investmentCalc(investment, expectedReturn, years)

	// futureValue := investment * math.Pow((1+expectedReturn/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Real value: %.1f\n", futureRealValue)
	fmt.Print(formatedFV, formatedRFV)

	// fmt.Println("Future value", futureValue)
	// fmt.Println((futureRealValue))
	// fmt.Printf("Future value: %.3f\nFuture Real value: %v\n", futureValue, futureRealValue)
	// fmt.Printf(`Future value: %.3f
	// Future Real value: %v`, futureValue, futureRealValue)
}

// func investmentCalc(investment, expectedReturn, years float64) (float64, float64) {
// 	fv := investment * math.Pow((1+expectedReturn/100), years)
// 	frv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, frv
// }

func investmentCalc(investment, expectedReturn, years float64) (fv float64, frv float64) {
	fv = investment * math.Pow((1+expectedReturn/100), years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
