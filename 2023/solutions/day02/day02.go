// Solution for day02 of the Advent of Code Challenge 2023
package day02

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

//go:embed input.txt
var input string

func Solve() {
	part1()
}

func part1() {
	max := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	total := 0
	powerTotal := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		id, err := extractGameID(line)
		if err != nil {
			util.LogFatal("unable to extract game id")
		}

		possible, err := checkGamePossible(line, max)
		if err != nil {
			util.LogFatal("invalid game: %s", line)
		}
		if possible {
			total += id
		}

		power, err := calcPowerGame(line)
		if err != nil {
			util.LogFatal("invalid game: %s", line)
		}
		powerTotal += power
	}
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", powerTotal)
}

func extractGameID(line string) (int, error) {
	res := strings.Split(line, ":")[0]
	id := strings.Split(res, " ")[1]
	return strconv.Atoi(id)
}

func checkGamePossible(line string, max map[string]int) (bool, error) {

	gamesLine := strings.SplitAfter(line, ": ")[1]
	games := strings.Split(gamesLine, "; ")
	for _, game := range games {
		pulls := strings.Split(game, ", ")
		for _, pull := range pulls {
			vals := strings.Split(pull, " ")
			num, err := strconv.Atoi(vals[0])
			if err != nil {
				return false, err
			}
			val, exists := max[vals[1]]
			if !exists || val < num {
				return false, nil
			}
		}
	}
	return true, nil
}

func calcPowerGame(line string) (int, error) {
	// Assume blue, green, red are the only cubes
	minCubes := map[string]int{"blue": 0, "green": 0, "red": 0}
	gamesLine := strings.SplitAfter(line, ": ")[1]
	games := strings.Split(gamesLine, "; ")
	for _, game := range games {
		pulls := strings.Split(game, ", ")
		for _, pull := range pulls {
			vals := strings.Split(pull, " ")
			num, err := strconv.Atoi(vals[0])
			if err != nil {
				return 0, err
			}

			if currMax := minCubes[vals[1]]; currMax < num {
				minCubes[vals[1]] = num
			}
		}
	}
	return minCubes["blue"] * minCubes["green"] * minCubes["red"], nil
}
