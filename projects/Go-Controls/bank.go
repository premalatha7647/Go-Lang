package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to bank")
	fmt.Println("1: Check balance")
	fmt.Println("2: Deposit money")
	fmt.Println("3: Withdraw money")
	fmt.Println("4: Exit")
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Your deposit")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Updated Balance:", accountBalance)
	} else if choice == 3 {
		var withdraw float64
		fmt.Print("Enter the amout to withdraw:")
		fmt.Scan(&withdraw)
		if withdraw > accountBalance {
			fmt.Println("Cant withdraw!")
			return
		}
		accountBalance -= withdraw
		fmt.Println("Updated balance:", accountBalance)
	}
}
