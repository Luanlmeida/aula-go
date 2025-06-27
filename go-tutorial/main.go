package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // Criar variavel e cetar ela
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // u = valor não pode ser negativo

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avalible", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Luan"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
