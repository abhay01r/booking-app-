package main

import "strings"

func validateUserInput(firstName string, lasttName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lasttName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && remainingTickets >= userTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
