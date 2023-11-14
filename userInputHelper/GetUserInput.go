package userinputhelper

import "fmt"

func GetUserInput() (string, string, int) {
	var firstName string
	var lastName string
	var bookedTicket int
	//pointer--storing the userName value in the address of userName Variable
	fmt.Printf("Enter Your First Name :\n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter Your Last Name :\n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter amount of Tickets you want to Book \n")
	fmt.Scan(&bookedTicket)
	return firstName, lastName, bookedTicket
}
