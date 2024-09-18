package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	// ... do something awesome with that gathered data

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promtpText string) string {
	fmt.Print(promtpText)

	var value string
	fmt.Scan(&value)
	return value
}
