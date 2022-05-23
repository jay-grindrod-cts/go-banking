package main

import (
	"banking-app/database"
	"fmt"
)

func main() {
	var name string
	var email string	

	database.DbConnection()

	fmt.Println()
	fmt.Println("Welcome to the Go Banking Application!")
	fmt.Println("------------------------------------")

	fmt.Println("Firstly, what is your name?")
	fmt.Scan(&name)
	fmt.Println("Next, what is your email")
	fmt.Scan(&email)

	// Checking DB for User
	database.CurrentUserCheck(email, name)	
}
