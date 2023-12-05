package main

import (
	"fmt"
	"math"
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

	var cards []Card
	for _, line := range input {
		nums := strings.Split(line, ": ")
		numsSplit := strings.Split(nums[1], " | ")
		winning := strings.Split(numsSplit[0], " ")
		yours := strings.Split(numsSplit[1], " ")

		cards = append(cards, Card{convertStringsToInts(winning), convertStringsToInts(yours), 1})
	}

	part1(cards)
	part2(cards)
}

func convertStringsToInts(strs []string) []int {
	var ints []int

	for _, val := range strs {
		num, err := strconv.Atoi(val)
		if err != nil {
			// panic("cant convert")
			continue
		}

		ints = append(ints, num)
	}

	return ints
}

type Card struct {
	winningNumbers []int
	yourNumbers    []int
	copies         int
}

func part1(input []Card) {

	sum := 0
	for _, card := range input {
		matchingAmount := 0
		for _, yourNum := range card.yourNumbers {
			if containsValue(card.winningNumbers, yourNum) {
				matchingAmount++
			}
		}
		sum += getScoreByMatches(matchingAmount)
	}

	fmt.Println("Part 1:", sum)
}

func containsValue(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func getScoreByMatches(matches int) int {
	return int(math.Pow(2, float64(matches)-1))
}

func part2(input []Card) {
	for i, card := range input {
		matchingAmount := 0
		for _, yourNum := range card.yourNumbers {
			if containsValue(card.winningNumbers, yourNum) {
				matchingAmount++
			}
		}

		for j := 1; j < matchingAmount+1; j++ {
			input[i+j].copies += card.copies
		}
	}

	sum := 0
	for _, card := range input {
		fmt.Println(card.copies)
		sum += card.copies
	}

	fmt.Println("Part 2:", sum)
}
