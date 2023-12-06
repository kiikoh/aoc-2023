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

	seeds := strListToInts(strings.Split(strings.Split(input[0], ": ")[1], " "))

	almanac := Almanac{seeds, []Section{}}

	secs := strings.Split(string(rawInput), "\n\n")

	for _, sec := range secs[1:] {
		section := Section{[]Map{}}
		for _, secMaps := range strings.Split(sec, "\n")[1:] {

			mapData := strListToInts(strings.Split(secMaps, " "))
			section.maps = append(section.maps, Map{mapData[0], mapData[1], mapData[2]})
		}
		almanac.sections = append(almanac.sections, section)
	}

	part1(almanac)
	part2(almanac)
}

func strListToInts(strs []string) []int {
	var ints []int

	for _, str := range strs {

		numValue, err := strconv.Atoi(str)

		if err != nil {
			panic("failed conversion")
		}

		ints = append(ints, numValue)
	}
	return ints
}

type Almanac struct {
	seeds    []int
	sections []Section
}

type Section struct {
	maps []Map
}

type Map struct {
	destStart int
	srcStart  int
	length    int
}

func part1(alm Almanac) {
	results := []int{}
	for _, seed := range alm.seeds {
		res := getLocation(alm.sections, 0, seed)
		results = append(results, res)
	}

	m := 999999999999999999
	for _, res := range results {
		m = min(res, m)
	}

	fmt.Println("Part 1:", m)
}

func part2(alm Almanac) {
	location := 0
	for {
		// we need to go backwards
		s := getSeed(alm.sections, len(alm.sections)-1, location)

		for i := 0; i < len(alm.seeds); i += 2 {
			if alm.seeds[i] < s && s < alm.seeds[i]+alm.seeds[i+1] {
				fmt.Println("Part 2:", location)
				return
			}
		}

		location++
	}

}

func getLocation(secs []Section, secIndex int, input int) int {

	if secIndex == len(secs) {
		return input
	}

	for _, m := range secs[secIndex].maps {
		if m.srcStart <= input && input < m.srcStart+m.length {
			return getLocation(secs, secIndex+1, input+m.destStart-m.srcStart)
		}
	}

	return getLocation(secs, secIndex+1, input)
}

func getSeed(secs []Section, secIndex int, input int) int {
	if secIndex == -1 {
		return input
	}

	for _, m := range secs[secIndex].maps {
		if m.destStart <= input && input < m.destStart+m.length {
			return getSeed(secs, secIndex-1, input-m.destStart+m.srcStart)
		}
	}

	return getSeed(secs, secIndex-1, input)
}
