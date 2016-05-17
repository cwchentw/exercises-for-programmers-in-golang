package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun: ")

	var noun string
	if scanner.Scan() {
		noun = scanner.Text()
	}

	fmt.Print("Enter a verb: ")

	var verb string
	if scanner.Scan() {
		verb = scanner.Text()
	}

	fmt.Print("Enter an adjective: ")

	var adj string
	if scanner.Scan() {
		adj = scanner.Text()
	}

	fmt.Print("Enter an adverb: ")

	var adv string
	if scanner.Scan() {
		adv = scanner.Text()
	}

	fmt.Printf("Do you %s your %s %s %s? That's hilarious\n",
		verb, adj, noun, adv)
}
