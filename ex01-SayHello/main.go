package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is your name? ")

	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Printf("Hello, %s, nice to meet you!\n", name)
}
