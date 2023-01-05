package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isEmailValid, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isEmailValid && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("All Tickets are booked out.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is invalid")
			}
			if !isEmailValid {
				fmt.Println("Email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is invalid")
			}
		}

	}
}

func greetUsers(conferenceName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are available.\n", confTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isEmailValid, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name:\n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name:\n")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address:\n")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets:\n")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		numOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
