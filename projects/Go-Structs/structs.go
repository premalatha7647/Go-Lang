package main

import (
	"fmt"
)

type user struct {
	firstName string,
	lastName string,
	bithDate string
	createdAt time.Time
}

func main() {

	firstName := getValue("Enter your first name: ")
	lastName := getValue("Enter your last name: ")
	birthDate := getValue("Enter your bithdate (MM/DD/YYYY): ")
	fmt.Print(firstName, lastName, birthDate)
}

func getValue(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
