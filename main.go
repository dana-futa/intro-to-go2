package main

import (
	"fmt"
	"strings"
)

func main() {
	// cannot use := syntax to assign value to a const
	// cannot use := syntax when specifying type of var

	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings [conferenceTickets]string // define an array
	// you can also do the following to initialize the array to an empty array:
	// var bookings = [conferenceTickets]string{}
	var bookings []string // define a slice aka dynamic array

	// check const/var type
	// fmt.Printf("conferenceName is a %T conferenceTickets is a %T remainingTickets is a %T\n", conferenceName, conferenceTickets, remainingTickets)

	// Welcome message
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickts are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	// can also do: for remainingTickets > 0 {}
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for first name
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		// ask user for last name
		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		// ask user for email
		fmt.Println("Please enter your email address: ")
		fmt.Scan(&email)

		// ask user for number of tickets
		fmt.Println("How many tickets would you like to purchase? ")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			// update ticket information
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			// Thank you message
			fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v are remaining for the %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			// _ acknowledge that there is a variable, but we don't want/need to use it.
			// In Go, you need to make unused variables explicit.
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of our bookings are %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The %v is now sold out. Please come back next year!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is invalid. The first name field requires 2 or more characters. The last name field requires 2 or more characters.")
			}
			if !isValidEmail {
				fmt.Println(("Your email address is invalid. The email address field requires an @ symbol."))
			}
			if !isValidTicketNumber {
				fmt.Println("The entered number of tickets is invalid. The number of tickets must be greater than zero and less than or equal to the total number of remaining tickets.")
				fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
			}

			fmt.Println("Please try again!")
		}
	}
}
