package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Count(number int) int { //this function count the number of digits there are in the number passed like parameter
	count := 0
	for number != 0 {
		number = number / 10
		count++
	}
	return count
}

func TranformInLetter(num, length int) string { //this function transforms the number (int) into letters
	var word string
	negative := false
	letters := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	if num < 0 {
		num = num * (-1)
		negative = true
	}
	for length > 0 {
		ris := float64(num)

		for ris >= 10 {
			ris = ris / 10
		}

		if length != Count(num) {
			word = word + letters[0]
		} else {
			a := int(ris)
			word = word + letters[a]
			num = num - a*(int(math.Pow(10, float64(length-1))))
		}
		length--
	}

	if negative {
		word = "(Negative)" + word
	}
	return word
}

func main() {
	numeric_numbers := []int{}
	result := []string{}
	args := os.Args[1:] //put parameters into string array (args)

	//convert strings to int
	for j := 0; j < len(args); j++ {
		number, _ := strconv.Atoi(args[j])
		numeric_numbers = append(numeric_numbers, number)
	}

	//transform numbers in letters
	for i := 0; i < len(numeric_numbers); i++ {
		result = append(result, TranformInLetter(numeric_numbers[i], Count(numeric_numbers[i])))
	}

	//I chose to print like this because is like the print in the instructions sheet
	fmt.Print(result[0])
	for i := 1; i < len(numeric_numbers); i++ {
		fmt.Print(", ", result[i])
	}
	fmt.Println("")
}
