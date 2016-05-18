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

	fmt.Print("Enter the principal: ")
	principal, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the principal")
	}

	fmt.Print("Enter the rate of interest: ")
	interest, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the interest")
	}

	fmt.Print("Enter the number of year: ")
	year, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the number of years")
	}

	fmt.Printf("After %d years at %.2f%%, the investment will be worth %.0f.\n",
		int(year), interest, principal * (1 + interest * year / 100.0))
}
