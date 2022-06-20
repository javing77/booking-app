package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		userName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTicket, userName, lastName, email)
			go sendEmail(userTicket, userName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The names of the bookings are : %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come the next year ")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Please enter a valid Name")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid Email")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of ticket that you enter is invalid")
			}
		}

	}

}

func bookTicket(userTicket uint, userName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	var userData = UserData{
		firstName:       userName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v \n", bookings)

	fmt.Printf("Thank you %v, %v for bookings %v tickets. You will recive email confirmitation at  %v  \n", userName, userTicket, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func getUserInput() (string, string, string, uint) {
	var userName string
	var lastName string
	var email string
	var userTicket uint
	// Ask user name

	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTicket)

	return userName, lastName, email, userTicket

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to  %v out conference bookings application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get yours tickes here to attend")
}

func sendEmail(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket: \n  %v  \n to email address %v \n", ticket, email)
	fmt.Println("########################")
}
