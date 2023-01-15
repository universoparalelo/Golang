package main

import "strings"

func validationUserInput(firstName string, lastName string, email string, userTicket int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	isValidTicketsNumber := userTicket > 0 && remainingTickets >= uint8(userTicket)

	return isValidName, isValidEmail, isValidTicketsNumber
}
