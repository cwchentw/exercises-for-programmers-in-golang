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

func F2C(temp float64) float64 {
	return (temp - 32.0) * 5.0 / 9.0
}

func C2F(temp float64) float64 {
	return (temp * 9.0 / 5.0) + 32.0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	var choice string
	if scanner.Scan() {
		choice = scanner.Text()

		if choice != "C" && choice != "F" {
			log.Fatalln("Wrong choice")
		}
	}

	if choice == "C" {
		fmt.Print("Please enter the temperature in Fahrenheit: ")
		temp, err := Input(scanner)
		if err != nil {
			log.Fatalln("Failed getting the temperature")
		}
		fmt.Printf("The temperature in Celsius is %.0f.\n", F2C(temp))
	} else if choice == "F" {
		fmt.Print("Please enter the temperature in Celsius: ")
		temp, err := Input(scanner)
		if err != nil {
			log.Fatalln("Failed getting the temperature")
		}
		fmt.Printf("The temperature in Fahrenheit is %.0f.\n", C2F(temp))
	}
}
