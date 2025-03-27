package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getValue("Revenue")
	expenses = getValue("Expenses")
	taxRate = getValue("Tax Rate")

	ebt := expenseBeforeTax(revenue, expenses)
	profit := profit(ebt, taxRate)
	ratio := ratio(ebt, profit)

	outputPrint(ebt, profit, ratio)
}

func getValue(text string) (val float64) {
	fmt.Printf("%v : ", text)
	fmt.Scan(&val)
	return val
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
