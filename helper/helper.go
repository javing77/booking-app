package helper

import "strings"

func ValidateUserInput(userName, lastName, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}
