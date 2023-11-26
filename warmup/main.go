package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	rawInput, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read file")
	}

	input := strings.Split(string(rawInput), "\n\n")

	var calories []int
	for _, g := range input {
		sum := 0
		for _, x := range strings.Split(g, "\n") {
			i, err := strconv.Atoi(x)

			if err != nil {
				// ... handle error
				panic(err)
			}

			sum += i
		}
		calories = append(calories, sum)
	}

	sort.Ints(calories)
	fmt.Println(calories)
	fmt.Println("Part 1: ", calories[len(calories)-1])

	fmt.Println(calories[len(calories)-3:])
	// Part 2
	sum := 0
	for _, v := range calories[len(calories)-3:] {
		sum += v
	}

	fmt.Println("Part 2: ", sum)
}
