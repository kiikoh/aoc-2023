package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read file")
	}

	input := strings.Split(string(rawInput), "\n")

	part1(input)
	part2(input)

}

func part1(input []string) {
	sum := 0
	for _, line := range input {
		firstDigit := -1
		lastDigit := -1

		for j := 0; j < len(line); j++ {
			digit, err := strconv.Atoi(string(line[j]))

			if err == nil {

				if firstDigit == -1 {
					firstDigit = digit
				}

				lastDigit = digit
			}

		}

		calibrationValue, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))

		if err != nil {
			panic("Can't parse result")
		}

		sum += calibrationValue

	}

	fmt.Println("Part 1 Solution:", sum)
}

func part2(input []string) {

	spelledOut := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	for _, line := range input {
		firstDigit := -1
		lastDigit := -1

		for j := 0; j < len(line); j++ {
			digit, err := strconv.Atoi(string(line[j]))

			if err == nil {

				if firstDigit == -1 {
					firstDigit = digit
				}

				lastDigit = digit
			} else {
				for v, word := range spelledOut {
					if j+len(word) < len(line)+1 && word == line[j:j+len(word)] {
						// we detected a number
						if firstDigit == -1 {
							firstDigit = v + 1
						}

						lastDigit = v + 1
					}
				}
			}

		}

		calibrationValue, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))

		if err != nil {
			panic("Can't parse result")
		}

		sum += calibrationValue

	}

	fmt.Println("Part 2 Solution:", sum)
}
