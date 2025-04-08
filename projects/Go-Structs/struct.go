package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	userFirstName := getValue("Enter your first name: ")
	userLastName := getValue("Enter your last name: ")
	userBirthDate := getValue("Enter your bithdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Print(err)
		return
	}
	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserData()
	appUser.OutputUserDetails()

}

//	func outputUserDetails(u *user) {
//		// fmt.Println((*u).firstName, (*u).lastName, (*u).bithDate, (*u).createdAt)
//		fmt.Println(u.firstName, u.lastName, u.bithDate, u.createdAt)
//	}
func getValue(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}
