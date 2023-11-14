package main

import (
	"fmt"
	greetUser "nana-booking-app/greetUser"
	userinputhelper "nana-booking-app/userInputHelper"
	"strings"
	"time"
)

var remainingTickets = 50
var bookedTicket int

func main() {
	var conferenceName = "Go Coonference"
	const conferenceTicket = 50
	var bookings []string
	var isBookMore string
	var firstName string
	var lastName string

	//we can even call variables from other package
	fmt.Printf(greetUser.AuthorName)
	//calling GreenUser from package helper by importing with the module name nana-booking-app/helper where helper is package name
	greetUser.GreetUser(conferenceName)
	fmt.Printf("We have total %v tickets available for %s \n", remainingTickets, conferenceName)
	fmt.Println("Get Your Tickets by filling the Form")

	firstName, lastName, bookedTicket = userinputhelper.GetUserInput()
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Welcome to %s, %s \n", conferenceName, bookings)
	isSuccess := generateTicket(bookedTicket, firstName, lastName)
	acknowledegBookingStatus(isSuccess)
	if remainingTickets != 0 {
		fmt.Printf("Do You want to Book more ticket \n")
		fmt.Scan(&isBookMore)
	}
	switch strings.ToLower(isBookMore) {
	case "yes":
		{
			firstName, lastName, bookedTicket = userinputhelper.GetUserInput()
			isSuccess = generateTicket(bookedTicket, firstName, lastName)
			acknowledegBookingStatus(isSuccess)
		}
	case "no":
		{
			fmt.Printf("Thank You!!!!")
			return
		}

	}

}

func generateTicket(bookedTicket int, firstName string, lastName string) bool {

	for i := 0; i < 25; i++ {
		fmt.Printf("#")
	}
	isTicketInputSucess := checkTicketInput(bookedTicket, remainingTickets)
	if isTicketInputSucess {
		fmt.Printf("\n %v tickets for %v %v is being booked\n", bookedTicket, firstName, lastName)
		//here the main thread is blocked
		time.Sleep(5 * time.Second)
		return true
	}
	return false

}

func checkTicketInput(bookedTicket int, remainingTickets int) bool {
	if bookedTicket > remainingTickets {
		fmt.Printf("\nERROR !!!! We have only %v  Tickets remaining and you can not book %v \n", remainingTickets, bookedTicket)
		return false
	}
	if remainingTickets == 0 {
		fmt.Println("\nConference Tickets are Sold Out!!, Come Back Next Year")
		return false
	}
	if bookedTicket == 0 {
		fmt.Printf("\nYou can not book %v Tickets. At least Book 1 Ticket", bookedTicket)
		return false
	}
	if bookedTicket < 0 {
		fmt.Printf("\nInvalid Number of ticket %v", bookedTicket)
		return false
	}
	return true
}

func acknowledegBookingStatus(isSuccess bool) {
	if isSuccess == true {
		fmt.Printf("Congratulation!!! You have booked total %v Tickets \n", bookedTicket)
		remainingTickets = remainingTickets - bookedTicket
		fmt.Printf("Total number of Tickets remaining : %v \n", remainingTickets)
	} else {
		fmt.Println("\n Ticket Can not be Booked")
	}
}
