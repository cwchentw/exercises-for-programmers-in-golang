package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

	fmt.Print("Please input length: ")
	length, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("Please input width: ")
	width, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	gallon := int(math.Ceil(float64(length) * float64(width) / float64(350)))
	fmt.Printf("You will need to purchase %d gallons of\n", gallon)
	fmt.Printf("paint to cover %d square feet.\n", length * width)
}
