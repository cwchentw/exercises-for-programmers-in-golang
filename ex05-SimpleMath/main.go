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
		n, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return 0, errors.New("Failed converting number")
		}

		return n, nil
	} else {
		return 0, errors.New("Failed reading from stdin")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the first number? ")
	num1, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("What is the second number? ")
	num2, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%d + %d = %d\n", num1, num2, num1 + num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1 - num2)
	fmt.Printf("%d * %d = %d\n", num1, num2, num1 * num2)
	fmt.Printf("%d / %d = %d\n", num1, num2, num1 / num2)
}
