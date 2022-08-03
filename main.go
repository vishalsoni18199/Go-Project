package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Total number of Tickets..
const conferenceTicket int = 50

var remainingTicket uint = 50
var confName = "CST"
var Booking = make([]UserData, 0)

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//Greeting the user and aware them how many tickets are remaining
	greetUsers()

	for {

		fName, lName, email, userTicket := getUserInput()

		//Validate the User Input
		isValidName, isValidEmail, isValidTicketNumber := validUserInput(fName, lName, email, userTicket)

		if isValidEmail && isValidName && isValidTicketNumber {

			//Booking Ticket for Conference...
			bookTicket(remainingTicket, userTicket, fName, email, lName)
			wg.Add(2)
			go sendTicket(userTicket, email)

			//Printing FirstName after booking
			firstName := printFirstName()
			fmt.Printf("The first name of the bookings are: %v\n", firstName)

			if remainingTicket == 0 {
				fmt.Printf("Webinar Ticket Sold out, Please Come back next time..")
				//break
			}
		} else {
			if !isValidName {
				fmt.Printf("Please enter correct name\n")
			}
			if !isValidEmail {
				fmt.Printf("Please enter correct email\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Please enter correct number of Tickets\n")
			}
		}
		fmt.Printf("####################################\n")

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Print("\n")
	fmt.Printf("\nwelcome to our %v conference\n", confName)
	fmt.Printf("Total Slots for Conference is %v\n", conferenceTicket)
	fmt.Printf("Book your Ticket now,\nwe have %v Ticket is available\n", remainingTicket)
}

func printFirstName() []string {
	fNames := []string{}
	for _, Bookings := range Booking {

		fNames = append(fNames, Bookings.FirstName)
	}
	return fNames
}

func validUserInput(fName string, lName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket
	return isValidEmail, isValidName, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var userTicket uint

	var fName string
	var lName string
	var email string
	//Ask the name of the user and number of Tickets
	fmt.Printf("Enter your first name\n")
	fmt.Scan(&fName)

	fmt.Printf("Enter your last name\n")
	fmt.Scan(&lName)

	//userName = fName + " " + lName

	fmt.Printf("Enter the email id\n")
	fmt.Scan(&email)

	fmt.Printf("enter the number of Tickets\n")
	fmt.Scan(&userTicket)

	return fName, lName, email, userTicket
}

func bookTicket(remainingTicket uint, userTicket uint, fName string, email string, lName string) {
	remainingTicket = remainingTicket - userTicket

	//create a map for a user
	var userData = UserData{
		FirstName:       fName,
		LastName:        lName,
		Email:           email,
		NumberOfTickets: userTicket,
	}

	Booking = append(Booking, userData)
	fmt.Printf("List of bookings is %v\n", Booking)

	fmt.Printf("Thank you!! %v for registering %v tickets\n", fName, userTicket)
	fmt.Printf("You will receive a confirmation mail at %v\n", email)
	fmt.Printf("%v Tickets are remaining for %v Conference\n", remainingTicket, confName)

}

func sendTicket(userTicket uint, email string) {
	time.Sleep(20 * time.Second)
	fmt.Println("#####################\n")
	fmt.Printf("%v, Tickets is sent to %v email id\n", userTicket, email)
	fmt.Println("#####################\n")
	wg.Done()
}
