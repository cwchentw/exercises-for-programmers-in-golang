package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your age? ")
	var age int
	if scanner.Scan() {
		nStr := scanner.Text()

		n, err := strconv.Atoi(nStr)
		if err != nil {
			log.Fatalln("Failed converting number")
		}
		age = n
	} else {
		log.Fatalln("Failed getting data from stdin")
	}

	if age < 16 {
		fmt.Println("You are not old enough to legally drive.")
	} else {
		fmt.Println("You are old enough to legally drive.")
	}
}
