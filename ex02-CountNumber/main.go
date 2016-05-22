package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the input string? ")
	var str string
	if scanner.Scan() {
		str = scanner.Text()
	}

	fmt.Printf("%s has %d characters.\n", str, utf8.RuneCountInString(str))
}
