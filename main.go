// Every code must belong to package
package main

import "fmt"

// A program can only have 1 "main" function, because we only have 1 entrypoint
func main() {
	// var confName = "Go Conf"
	confName := "Go Conf" // Syntatix suger with variable, not apply to constant
	const confTickets int = 50
	var remainingTickets uint = 50

	// String Interpolation ---
	// Solution 1.
	// fmt.Println("Welcome to", confName, "booking application")
 	// fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
 	// fmt.Println("Get your tickers here to attend")

	// Solution 2
	fmt.Printf("confName is %T, confTickets is %T, remainingTickets is %T\n", confName, confTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", confName)
 	fmt.Printf("We have a total of %v tickets and %v are still available \n", confTickets, remainingTickets)
 	fmt.Println("Get your tickers here to attend")

	// Data Type ---
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Arrays & Slices
	// var bookings [50]string
	// var bookings = []string{}
	// bookings := []string{}
	var bookings []string

	// User Input & Pointer
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your emai: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	//* Booking logic in application ---
	// remainingTickets = remainingTickets - userTickets
	remainingTickets -= userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	//* ------

	fmt.Printf("Thank you %v %v have booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)

	fmt.Printf("All bookings: %v", bookings)
}
