package main

import (
	"fmt"

	"unsafemango.com/structs/user"
)

func main() {
	generic()
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// instance of a struct
	// struct literal
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		// exit out of the application
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "password")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data
	appUser.OutputUserDetails()
	appUser.ClearUserName()

	appUser.OutputUserDetails()
}

func getUserData(promtpText string) string {
	fmt.Print(promtpText)

	var value string
	fmt.Scanln(&value)
	return value
}
