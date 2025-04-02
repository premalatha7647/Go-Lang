package main

import "fmt"

func main() {

	age := 24
	var agePointer *int
	fmt.Println("age", age)
	agePointer = &age
	fmt.Println("age", agePointer, *agePointer)
	adultAge := getAdultYears(agePointer)

	fmt.Println(adultAge)
}

func getAdultYears(age *int) int {
	return *age - 18
}
