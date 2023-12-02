package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func main() {

	rawInput, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Could not read file")
	}

	input := strings.Split(string(rawInput), "\n")

	games := []Game{}
	for gameId, line := range input {
		game := Game{gameId + 1, []Round{}}

		roundInput := strings.Join(strings.Split(line, ": ")[1:], "")

		for _, roundData := range strings.Split(roundInput, "; ") {
			round := Round{0, 0, 0}

			for _, colorCount := range strings.Split(roundData, ", ") {
				split := strings.Split(colorCount, " ")
				count, err := strconv.Atoi(split[0])
				color := split[1]

				if err != nil {
					panic("Couldnt parse count")
				}

				switch color {
				case "red":
					round.red = count
				case "green":
					round.green = count
				case "blue":
					round.blue = count
				}
			}

			game.rounds = append(game.rounds, round)

		}
		games = append(games, game)
	}

	// fmt.Println(games)

	part1(games)
	part2(games)
}

func part1(games []Game) {
	idSum := 0
	for _, game := range games {
		// is the game valid

		valid := true
		for _, round := range game.rounds {
			if round.red > 12 || round.green > 13 || round.blue > 14 {
				// fmt.Println(game.id, "is invalid")
				valid = false
			}
		}

		if valid {
			idSum += game.id
		}
	}

	fmt.Println("Part 1:", idSum)
}

func part2(games []Game) {
	powerSum := 0
	for _, game := range games {

		minR := 0
		minG := 0
		minB := 0

		for _, round := range game.rounds {
			if minR < round.red {
				minR = round.red
			}

			if minG < round.green {
				minG = round.green
			}

			if minB < round.blue {
				minB = round.blue
			}
		}
		// fmt.Println(minR * minG * minB)

		powerSum += minR * minG * minB

	}

	fmt.Println("Part 2:", powerSum)
}
