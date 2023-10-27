package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Were you born before Y2K?")
	year := getBirthYear()
	isBornBeforeY2K := checkIfBornBeforeY2K(year)

	if isBornBeforeY2K {
		fmt.Println("Yes, you were born before Y2K.")
	} else {
		fmt.Println("No, you were born in or after the year 2000.")
	}
}

func getBirthYear() int {
	for {
		fmt.Print("Enter your birth year: ")
		var birthYearStr string
		_, err := fmt.Scanln(&birthYearStr)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid birth year.")
		} else {
			birthYear, err := strconv.Atoi(birthYearStr)
			if err != nil {
				fmt.Println("Invalid input, please enter a valid number.")
			} else {
				return birthYear
			}
		}
	}
}

func CheckIfBornBeforeY2K(year int) bool {
	if year < 2000 {
		return true
	}
	return false
}
