package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
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

	fmt.Print("What is your current age? ")
	age, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("At what age would you like to retire? ")
	retirement, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("You have %d years left until you can retire.\n",
		retirement - age)

	currentYear := time.Now().Year()
	fmt.Printf("It's %d, so you can retire in %d\n",
		currentYear, currentYear + retirement - age)
}
