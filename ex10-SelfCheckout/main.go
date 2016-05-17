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

	type item struct {
		price int
		quantity int
	}

	items := make([]item, 0)

	count := 1
	for {
		fmt.Printf("Enter the price of item %d: ", count)
		price, err := Input(scanner)
		if err != nil {
			break
		}

		fmt.Printf("Enter the quantity of item %d: ", count)
		quantity, err := Input(scanner)
		if err != nil {
			break
		}

		items = append(items, item{price, quantity})
		count++
	}

	subtotal := 0.0
	for _, item := range items {
		subtotal += float64(item.price) * float64(item.quantity)
	}

	const rate = 0.055
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax: $%.2f\n", subtotal * rate)
	fmt.Printf("Total: $%.2f\n", subtotal * (1 + rate))
}
