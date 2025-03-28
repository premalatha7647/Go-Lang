package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var err error
	revenue, err := getValue("Revenue")
	// if err != nil {
	// 	fmt.Println("ERROR:", err)
	// 	return
	// }
	expenses, err := getValue("Expenses")
	// if err != nil {
	// 	fmt.Println("ERROR", err)
	// 	return
	// }
	taxRate, err := getValue("Tax Rate")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	ebt := expenseBeforeTax(revenue, expenses)
	profit := profit(ebt, taxRate)
	ratio := ratio(ebt, profit)

	outputPrint(ebt, profit, ratio)
	data := fmt.Sprintf("EBT: %f\nProfit: %f\nratio: %f", ebt, profit, ratio)
	os.WriteFile("investment.txt", []byte(data), 0644)
}

func getValue(text string) (float64, error) {
	fmt.Printf("%v : ", text)
	var val float64
	fmt.Scan(&val)
	// if val <= 0 {
	// 	panic("Invalid input")
	// }
	// return val
	if val <= 0 {
		return 0, errors.New("invalid input")
	}
	return val, nil
}

func expenseBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}
func profit(ebt, taxRate float64) (profit float64) {
	profit = ebt * (1 - taxRate/100)
	return profit
}

func ratio(ebt, profit float64) (ratio float64) {
	ratio = ebt / profit
	return ratio
}

func outputPrint(ebt float64, profit float64, ratio float64) {
	formatedEBT := fmt.Sprintf("Earning before tax: %.1f\n", ebt)
	formatedProfit := fmt.Sprintf("Earning after tax: %.1f\n", profit)
	formatedRatio := fmt.Sprintf("Tax ratio %f\n", ratio)
	fmt.Print(formatedEBT, formatedProfit, formatedRatio)
}
