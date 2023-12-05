package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	rawInput, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read file")
	}

	input := strings.Split(string(rawInput), "\n")

	// add a . to the end of each line to make parsing easier
	var grid = []string{}
	for _, line := range input {
		grid = append(grid, line+".")
	}

	fmt.Println(len(grid[0]))

	part1(grid)
	part2(grid)
}

func part1(input []string) {

	sum := 0
	for i, line := range input {
		partNumber := ""
		for j, character := range line {

			number, err := strconv.Atoi(string(character))

			if err != nil {
				// problem parsing the digit,
				// if there is a number stored in the buffer,
				// end the buffer and add to sum if it touches a symbol

				if len(partNumber) > 0 {
					partNumberParsed, err := strconv.Atoi(string(partNumber))

					if err != nil {
						panic("failed to parse")
					}

					// determine if it neighbors a symbol
					// fmt.Println("Begin search around", partNumberParsed)
					if searchArea(input, i, j, len(partNumber)) {
						fmt.Println("Valid", partNumberParsed)
						sum += partNumberParsed
					} else {
						fmt.Println("Invalid", partNumberParsed)
					}

					partNumber = ""
				}

			} else {
				// we have a number
				partNumber += strconv.Itoa(number)
			}

		}
	}

	fmt.Println("Part 1:", sum)
}

func searchArea(grid []string, y int, x int, width int) bool {
	partPattern := regexp.MustCompile("[^0-9.]")

	// fmt.Println("Starting at", x, y)
	for i := -width - 1; i < 1; i++ {
		for j := -1; j <= 1; j++ {
			// fmt.Println(i, j)
			if 0 > x+i || x+i > len(grid[0])-1 || 0 > y+j || y+j > len(grid)-1 {
				continue
			}

			if partPattern.MatchString(string(grid[y+j][x+i])) {
				return true
			}
		}
	}

	return false
}

type location struct {
	x int
	y int
}

func part2(input []string) {

	gearLocMap := make(map[location]int)

	sum := 0
	for i, line := range input {
		partNumber := ""
		for j, character := range line {

			number, err := strconv.Atoi(string(character))

			if err != nil {
				// problem parsing the digit,
				// if there is a number stored in the buffer,
				// end the buffer and add to sum if it touches a symbol

				if len(partNumber) > 0 {
					partNumberParsed, err := strconv.Atoi(string(partNumber))

					if err != nil {
						panic("failed to parse")
					}

					// determine if it neighbors a symbol
					// fmt.Println("Begin search around", partNumberParsed)

					gearLoc := searchAreaForGears(input, i, j, len(partNumber))

					// found a gear
					if gearLoc.x != -1 && gearLoc.y != -1 {
						fmt.Println(gearLocMap[gearLoc])
						// add to gear map
						if gearLocMap[gearLoc] == 0 {
							gearLocMap[gearLoc] = partNumberParsed
						} else {
							sum += partNumberParsed * gearLocMap[gearLoc]
							fmt.Println("Gear connected to", partNumberParsed, "and", gearLocMap[gearLoc], "@", gearLoc.x, gearLoc.y)
						}
					}

					partNumber = ""
				}

			} else {
				// we have a number
				partNumber += strconv.Itoa(number)
			}

		}
	}

	fmt.Println("Part 2:", sum)
}

func searchAreaForGears(grid []string, y int, x int, width int) location {
	// fmt.Println("Starting at", x, y)
	for i := -width - 1; i < 1; i++ {
		for j := -1; j <= 1; j++ {
			// fmt.Println(i, j)
			if 0 > x+i || x+i > len(grid[0])-1 || 0 > y+j || y+j > len(grid)-1 {
				continue
			}

			if string(grid[y+j][x+i]) == "*" {
				return location{y + j, x + i}
			}
		}
	}

	return location{-1, -1}
}
