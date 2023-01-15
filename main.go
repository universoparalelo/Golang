package main

import (
	"fmt"
	"sync"
	"time"
)

// Package level variables
const conferenceTickets int8 = 50

var conferenceName = "Go conference"
var remainingTickets uint8 = 50

var booking = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	numberTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTicket := userInput()

	isValidName, isValidEmail, isValidTicketsNumber := validationUserInput(firstName, lastName, email, userTicket)

	if isValidEmail && isValidName && isValidTicketsNumber {

		bookingTickets(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendingEmail(userTicket, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noRemainingTickets := remainingTickets == 0
		if noRemainingTickets {
			// end of program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First o last name is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email doesn't contain an @ sign.")
		}
		if !isValidTicketsNumber {
			fmt.Println("The number of tickets is invalid.")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, values := range booking {
		// solution for a slice of maps
		firstNames = append(firstNames, values.firstName)
	}
	return firstNames
}

func userInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int

	// ask the user their name and amount of tickets booked
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets booked: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookingTickets(userTicket int, firstName string, lastName string, email string) {

	remainingTickets -= uint8(userTicket)

	// creating a varible to store the struct for an user
	userData := UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		numberTickets: uint(userTicket),
	}

	booking = append(booking, userData)
	fmt.Printf("List of booking: %v\n", booking)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll received a confirmation on you email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v remaining tickets for %v.\n", remainingTickets, conferenceName)
}

func sendingEmail(userTicket int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("Sending the ticket: \n%v \nto the email %v\n", ticket, email)
	fmt.Println("-------------------------------------------------------")

	wg.Done()
}

// Seccion comentarios que eran parte del codigo pero ya no
// "strings"
// var booking [50]string
// fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)
// isValidCity = city == "Singapore" || city == "London"
// continue
// solution for a slice of strings
// var name = strings.Fields(values)
// firstNames = append(firstNames, name[0])
// booking[0] = firstName + " " + lastName
// fmt.Printf("The whole array: %v\n", booking)
// fmt.Printf("The first user: %v\n", booking[0])
// fmt.Printf("Slice type: %T\n", booking)
// fmt.Printf("Slice length: %v\n", len(booking))
// var booking = make([]map[string]string, 0)
// firstNames = append(firstNames, values["firstName"])

// creating a map for a user
// userData := make(map[string]string)
// userData["firstName"] = firstName
// userData["lastName"] = lastName
// userData["email"] = email
// userData["numberTickets"] = strconv.FormatUint(uint64(userTicket), 10)
