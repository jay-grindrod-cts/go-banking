package database

import (
	"fmt"
	"os"
)

func AddUser(name string, email string) {
	result, err := db.Exec("INSERT INTO users (email, balance, name) VALUES (?, 0, ?)", email, name)
	if err == nil {
		result.LastInsertId()
		fmt.Printf("Success! Your account is set up with the email: %v \n", email)
	} else {
		fmt.Printf("Error adding user to DB. Erorr: %v", err)
		fmt.Println()
		os.Exit(1)
	}
}

func DepositMoney(depositAmount float64, email string) {
	result, err := db.Exec("UPDATE users SET balance=balance+? WHERE email=?", depositAmount, email)
	if err == nil {
		result.LastInsertId()
	} else {
		fmt.Printf("Error adding the Dollar to your Account. Erorr: %v", err)
		fmt.Println()
		os.Exit(1)
	}
}

func WithdrawMoney(withdrawAmount float64, email string) {
	result, err := db.Exec("UPDATE users SET balance=balance-? WHERE email=?", withdrawAmount, email)
	if err == nil {
		result.LastInsertId()
	} else {
		fmt.Printf("Error retrieving the Dollar from your Account. Erorr: %v", err)
		fmt.Println()
		os.Exit(1)
	}
}