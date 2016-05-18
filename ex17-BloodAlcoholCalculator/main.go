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

type Sex int
const (
	M Sex = iota
	F
)

func CalcBAC(alcohol float64, weight float64, hour float64, sex Sex) float64 {
	var distro float64
	if sex == M {
		distro = 0.73
	} else {
		distro = 0.66
	}

	bac := (alcohol * 5.14 / weight * distro) - 0.15 * hour

	if bac > 0 {
		return bac
	} else {
		return 0.0
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the gender (M/F): ")
	var sex Sex
	if scanner.Scan() {
		_sex := scanner.Text()

		if _sex == "M" || _sex == "m" {
			sex = M
		} else if _sex == "F" || _sex == "f" {
			sex = F
		} else {
			log.Fatalln("Unknown sex")
		}
	} else {
		log.Fatalln("Failed getting the gender")
	}

	fmt.Print("Enter the total alcohol consumed, in ounces: ")
	alcohol, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting alcohol amount")
	}

	fmt.Print("Enter the body weight, in pounds: ")
	weight, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting body weight")
	}

	fmt.Print("Enter the number of hours since the last drink: ")
	hour, err := Input(scanner)
	if err != nil {
		log.Fatalln("Failed getting the number of hours")
	}

	bac := CalcBAC(alcohol, weight, hour, sex)

	fmt.Printf("Your BAC is %.2f\n", bac)

	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive")
	}
}
