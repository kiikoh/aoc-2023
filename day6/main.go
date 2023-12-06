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

	millis := strListToInts(strings.Split(input[0], " "))
	minDists := strListToInts(strings.Split(input[1], " "))

	part1(millis, minDists)
	part2(concatNumbers(millis), concatNumbers(minDists))
}

func concatNumbers(nums []int) int {
	strNum := ""

	for _, num := range nums {
		strNum += strconv.Itoa(num)
	}

	num, _ := strconv.Atoi(strNum)
	return num
}

func strListToInts(strs []string) []int {
	var ints []int

	for _, str := range strs {

		numValue, err := strconv.Atoi(str)

		if err != nil {
			continue
		}

		ints = append(ints, numValue)
	}
	return ints
}

func part1(millis []int, minDists []int) {

	product := 1
	for i, mm := range millis {
		minDist := minDists[i]

		numTimes := 0
		for chargeTime := 0; chargeTime < mm; chargeTime++ {
			if chargeTime*(mm-chargeTime) > minDist {
				numTimes++
			}

		}
		product *= numTimes
	}
	fmt.Println("Part 1:", product)
}

func part2(mm int, minDist int) {

	numTimes := 0
	for chargeTime := 0; chargeTime < mm; chargeTime++ {
		if chargeTime*(mm-chargeTime) > minDist {
			numTimes++
		}
	}
	fmt.Println("Part 2:", numTimes)
}
