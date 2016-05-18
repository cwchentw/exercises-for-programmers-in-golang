package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Input(scanner *bufio.Scanner) (float64, error) {
	if scanner.Scan() {
		nStr := scanner.Text()
		n, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			return 0, errors.New("Failed converting number")
		}

		return n, nil
	} else {
		return 0, errors.New("Failed to read text")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the order amount? ")
	amount, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting amount")
	}

	fmt.Print("What is the state? ")
	var state string
	if scanner.Scan() {
		state = scanner.Text()
	}

	const taxRate = 0.055
	if state == "WI" {
		fmt.Printf("The subtotal is $%.2f.\n", amount)
		fmt.Printf("The tax is $%.2f.\n", amount * taxRate)
		fmt.Printf("The total is $%.2f.\n", amount * (1.0 + taxRate))
	} else {
		fmt.Printf("The total is $%.2f.\n", amount)
	}
}
