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

	fmt.Print("How many euros are you exchanging? ")
	euros, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting euro amount")
	}

	fmt.Print("What is the exchange rate? ")
	rate, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting exchange rate")
	}

	fmt.Printf("%.2f euros at an exchange rate of %.2f is\n%.2f U.S. dollars\n",
		euros, rate, euros * rate / 100.0)
}
