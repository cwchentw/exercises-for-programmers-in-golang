package main

import (
	"fmt"
	"log"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Print("What is the password? ")

	password, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalln("Failed entering password")
	}

	if string(password) != "abc$123" {
		fmt.Println("\nI don't know you.")
	} else {
		fmt.Println("\nWelcome!")
	}
}
