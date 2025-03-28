package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var balanceFile = "balance.txt"

func main() {

	var accountBalance, err = readBalanceFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		panic("cant continue, sorry!")
	}
	fmt.Println("Welcome to bank")
	// for i := 0; i < 2; i++ {
	for {
		fmt.Println("1: Check balance")
		fmt.Println("2: Deposit money")
		fmt.Println("3: Withdraw money")
		fmt.Println("4: Exit")
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Updated Balance:", accountBalance)
			writeBalanceFile(accountBalance)

		case 3:
			var withdraw float64
			fmt.Print("Enter the amout to withdraw:")
			fmt.Scan(&withdraw)
			if withdraw > accountBalance {
				fmt.Println("Cant withdraw!")
				// return
				continue
			}
			accountBalance -= withdraw
			fmt.Println("Updated balance:", accountBalance)
			writeBalanceFile(accountBalance)
		default:
			fmt.Print("GoodBye!")
			fmt.Print("Thank you for banking wit us!")
			return

		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("Your deposit")
		// 	fmt.Scan(&depositAmount)
		// 	accountBalance += depositAmount
		// 	fmt.Println("Updated Balance:", accountBalance)
		// } else if choice == 3 {
		// 	var withdraw float64
		// 	fmt.Print("Enter the amout to withdraw:")
		// 	fmt.Scan(&withdraw)
		// 	if withdraw > accountBalance {
		// 		fmt.Println("Cant withdraw!")
		// 		// return
		// 		continue
		// 	}
		// 	accountBalance -= withdraw
		// 	fmt.Println("Updated balance:", accountBalance)
		// } else {
		// 	fmt.Print("GoodBye!")
		// 	// return
		// 	break
		// }
	}

}

func writeBalanceFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	err := os.WriteFile("balance.txt", []byte(balanceTxt), 0644)
	fmt.Println(err)
}
func readBalanceFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("failed to find the balance file")

	}
	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil {
		return 1000, errors.New("failed to parse the value")
	}
	return balance, nil
}
