package main

import "fmt"

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

const conferenceName string = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []UserData = make([]UserData, 0)

func main() {

	greetUsers()

	for {
		userData := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userData)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userData)
			printFirstNames()

			// end program
			if remainingTickets == 0 {
				fmt.Printf("The %v is now sold out. Please come back next year!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is invalid. The first name field requires 2 or more characters. The last name field requires 2 or more characters.")
			}
			if !isValidEmail {
				fmt.Println("Your email address is invalid. The email address field requires an @ symbol.")
			}
			if !isValidTicketNumber {
				fmt.Println("The entered number of tickets is invalid. The number of tickets must be greater than zero and less than or equal to the total number of remaining tickets.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickts are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInput() UserData {
	var userData = UserData{}

	// ask user for first name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&userData.firstName)

	// ask user for last name
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&userData.lastName)

	// ask user for email
	fmt.Println("Please enter your email address: ")
	fmt.Scan(&userData.email)

	// ask user for number of tickets
	fmt.Println("How many tickets would you like to purchase? ")
	fmt.Scan(&userData.tickets)

	return userData
}

func bookTicket(userData UserData) {
	remainingTickets = remainingTickets - userData.tickets
	bookings = append(bookings, userData)

	// Thank you message
	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n",
		userData.firstName, userData.lastName, userData.tickets, userData.email)
	fmt.Printf("%v are remaining for the %v.\n", remainingTickets, conferenceName)
}

func printFirstNames() {
	fmt.Printf("The first names of our bookings are: ")

	for index, booking := range bookings {
		fmt.Printf("%v", booking.firstName)
		if index < len(bookings)-1 {
			fmt.Printf(", ")
		}
	}

	fmt.Printf("\n")
}
