package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	// remainingTickets := 50 // alt var syntax, doesn't work for const
	
	// arrays in go have type of: fixed (maximum) size and variable type, ie [50]string
	// a Slice is an abstraction of an Array with dynamic size, ie []string
	var bookings []string
	// bookings := []string{} // alt syntax for implied type with empty slice assignment

	// %v subs the value of a variable; can also use %T to sub the type
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("%v of %v tickets are still available.\n", remainingTickets, conferenceTickets)

	for {
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

		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email address. Please try again.")
			continue // skip to next iteration of main loop
		}
	
		fmt.Print("Please enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Requested number of tickets (%v) exceeds total available (%v). Please try again.\n",
				userTickets, remainingTickets)
			continue // skip to next iteration of main loop
		}
	
		remainingTickets = remainingTickets - userTickets
	
		// you can set a value to a specific Array index like this:
		// bookings[0] = firstName + " " + lastName 
		// but can also just append() the value to a Slick
		bookings = append(bookings, firstName + " " + lastName)
	
		// fmt.Printf("bookings: %v\n", bookings)
		// fmt.Printf("bookings[0]: %v\n", bookings[0])
		// fmt.Printf("bookings type: %T\n", bookings)
		// // prints whole Array including all the empty spaces
		// // use Slice for dynamic size
		// fmt.Printf("bookings length: %v\n", len(bookings))
	
		fmt.Printf("Thank you, %v %v! You have have booked %v tickets. You will receive a confimration email at %v.\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
		firstNames := []string{}
		// range iterates over elements in different data structures
		// for Arrays/Slices, range provides the index and value for each element
		// underscore = blank identifier for the index, which is not used in our loop
		for _, booking := range bookings  {
			// strings.Fields(string) takes the string and splits it at white space;
			// strings.Fields("Sundar Pichai" => ["Sundar", "Pichai"])
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Current bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Printf("%v is booked out; please come back next year!\n", conferenceName)
			break
		}
	}
}