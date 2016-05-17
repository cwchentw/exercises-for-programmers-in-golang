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

	fmt.Print("What is the length of the room in feet? ")
	length, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("What is the width of the room in feet? ")
	width, err := Input(scanner)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("You entered dimension of %d feet by %d feet.\n",
		length, width)

	const factor = 0.09290304
	fmt.Printf("The area is\n%d square feet\n%.3f square meters\n",
		length * width, float32(length) * float32(width) * factor)
}
