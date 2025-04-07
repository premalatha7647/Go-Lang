package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	bithDate  string
	createdAt time.Time
}

// struct methods -- called receiver
func (u user) outputUserDetails() {
	// fmt.Println((*u).firstName, (*u).lastName, (*u).bithDate, (*u).createdAt)
	fmt.Println(u.firstName, u.lastName, u.bithDate, u.createdAt)
}

func (u *user) clearUserData() {
	u.firstName = ""
	u.lastName = ""
}

//	func (u user) clearUserData() {
//		u.firstName = ""
//		u.lastName = ""
//	}
func main() {

	userFirstName := getValue("Enter your first name: ")
	userLastName := getValue("Enter your last name: ")
	userBirthDate := getValue("Enter your bithdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		bithDate:  userBirthDate,
		createdAt: time.Now(),
	}
	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserData()
	appUser.outputUserDetails()

}

//	func outputUserDetails(u *user) {
//		// fmt.Println((*u).firstName, (*u).lastName, (*u).bithDate, (*u).createdAt)
//		fmt.Println(u.firstName, u.lastName, u.bithDate, u.createdAt)
//	}
func getValue(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)
	return value
}
