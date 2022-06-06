package main

import (
	"fmt"
	"strings"
	// "github.com/mattmuroya/go-booking-app/common"
)

// defining vars outside main function makes them "package level variables"
// go will infer var types, but you can also declare them manually
const conferenceName string = "Go Conference"
const conferenceTickets uint = 50

var remainingTickets uint = 50

// arrays in go have type of: fixed (maximum) size and variable type, ie [50]string
// a Slice is an abstraction of an Array with dynamic size, ie []string
var bookings []string

// bookings := []string{} // alt var syntax for implied type with empty slice initial assignment
// this alt syntax cannot be used at the package level

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidTicketAmount := ValidateUserInput(email, userTickets)

		if !isValidEmail {
			fmt.Println("Invalid email address. Please try again.")
			continue // skip to next iteration of main loop
		} else if !isValidTicketAmount {
			fmt.Printf("Requested number of tickets (%v) exceeds total available (%v). Please try again.\n",
				userTickets, remainingTickets)
			continue // skip to next iteration of main loop
		}

		bookTickets(userTickets, firstName, lastName, email)

		fmt.Printf("Current bookings: %v\n", getFirstNames())

		if remainingTickets == 0 {
			// end the program
			fmt.Printf("%v is booked out; please come back next year!\n", conferenceName)
			break
		}
	}
}

func greetUsers() {
	// %v subs the value of a variable; can also use %T to sub the type
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("%v of %v tickets are still available.\n", remainingTickets, conferenceTickets)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// - prompt for user input and assign it to firstName
	// - & indicates a pointer, which is a variable that
	//   points to the memory address of the value of firstName
	// - Scan receives the memory address, not the value of
	//   firstName (which is currently string zero value "")
	// fmt.Println(remainingTickets) // returns ""
	// fmt.Println(&remainingTickets) // returns 0xc0000b2008

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func ValidateUserInput(email string, userTickets uint) (bool, bool) {
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets <= remainingTickets
	// can have multiple return values, typed individually in first line of function
	return isValidEmail, isValidTicketAmount
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// you can set a value to a specific Array index like this:
	// bookings[0] = firstName + " " + lastName
	// but can also just append() the value to a Slick
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you, %v %v! You have have booked %v tickets. You will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	// range iterates over elements in different data structures
	// for Arrays/Slices, range provides the index and value for each element
	// underscore = blank identifier for the index, which is not used in our loop
	for _, booking := range bookings {
		// strings.Fields(string) takes the string and splits it at white space;
		// strings.Fields("Sundar Pichai" => ["Sundar", "Pichai"])
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
