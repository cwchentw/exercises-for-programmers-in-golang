package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Input(scanner *bufio.Scanner) (int, error) {
	if scanner.Scan() {
		nStr := scanner.Text()
		n, err := strconv.Atoi(nStr)

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

	fmt.Print("How many people? ")
	people, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("How many pizzas do you have? ")
	pizza, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%d people with %d pizzas\n", people, pizza)
	fmt.Printf("Each person gets %d pieces of pizza.\n",
		pizza * 8 / people)
	fmt.Printf("There are %d leftover pieces.\n",
		pizza * 8 % people)
}
