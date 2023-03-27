// Every code must belong to package
package main

import "fmt"

// A program can only have 1 "main" function, because we only have 1 entrypoint
func main() {
	// var confName = "Go Conf"
	confName := "Go Conf" // Syntatix suger with variable, not apply to constant
	const confTickets int = 50
	remainingTickets := 50

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
	var userName string
	var userTickets int

	userName = "Test"
	userTickets = 100

	fmt.Printf("User %v have booked %v tickets", userName, userTickets)
}
