package greetUser

import "fmt"

var AuthorName = "created by Sandeep"

func GreetUser(conferenceName string) {
	fmt.Printf("\nWelcome to %v Booking Application\n", conferenceName)
}
