package user

import (
	"errors"
	"fmt"
	"time"
)

// initializing a custom type user struct
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct type that builds up on the user struct type by embedding the user in it
type Admin struct {
	email    string
	password string
	User
}

// method for the struct to output data
func (user *User) OutputUserDetails() { // this is a special kind of argument called the receiver argument pointing at the user struct
	// ...
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// method for the struct to clear the user name
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

// utility function to create a struct (constructor intialization)
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}

	return &User{ // using a pointer to prevent a copy from being made
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
