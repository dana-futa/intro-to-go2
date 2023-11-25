package main

import "strings"

// other good file names include: helper.go or shared.go

func validateUserInput(userData UserData) (bool, bool, bool) {
	isValidName := len(userData.firstName) >= 2 && len(userData.lastName) >= 2
	isValidEmail := strings.Contains(userData.email, "@")
	isValidTicketNumber := userData.tickets > 0 && userData.tickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
