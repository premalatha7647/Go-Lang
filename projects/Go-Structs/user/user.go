package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	bithDate  string
	createdAt time.Time
}

// struct methods -- called receiver
func (u *User) OutputUserDetails() {
	// fmt.Println((*u).firstName, (*u).lastName, (*u).bithDate, (*u).createdAt)
	fmt.Println(u.firstName, u.lastName, u.bithDate, u.createdAt)
}

func (u *User) ClearUserData() {
	u.firstName = ""
	u.lastName = ""
}

//	func (u User) clearUserData() {
//		u.firstName = ""
//		u.lastName = ""
//	}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("firstName,lastName and birthdate are required")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		bithDate:  userBirthDate,
		createdAt: time.Now(),
	}, nil
}
