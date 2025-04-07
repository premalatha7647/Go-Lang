package main

import "fmt"

func main() {

	age := 24
	var agePointer *int
	fmt.Println("age", age)
	agePointer = &age
	fmt.Println("age", agePointer, *agePointer)
	getAdultYears(&age)

	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
