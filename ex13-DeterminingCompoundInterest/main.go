package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

	fmt.Print("What is the principal amount? ")
	principal, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the principal")
	}

	fmt.Print("What is the rate? ")
	rate, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the rate")
	}

	fmt.Print("What is the number of years? ")
	year, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the number of years")
	}

	fmt.Print(`What is the number of times the interest
is compounded per year? `)
	perYear, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the number of times per year")
	}

	amount := principal * math.Pow((1 + rate / 100.0 / perYear), year * perYear)
	fmt.Printf("$%d invested at %.2f%% for %d years\ncompunded %d times per year is $%.2f.\n",
		int(principal), rate, int(year), int(perYear), amount)
}
