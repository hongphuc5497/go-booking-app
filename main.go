// Every code must belong to package
package main

import "fmt"

// A program can only have 1 "main" function, because we only have 1 entrypoint
func main() {
	var confName = "Go Conf"
	const confTickets = 50
	var remainingTickets = 50
 	
	fmt.Println("Welcome to", confName, "booking application")
 	fmt.Println("We have a total of", confTickets, "tickets and", remainingTickets, "are still available")
 	fmt.Println("Get your tickers here to attend")
}
