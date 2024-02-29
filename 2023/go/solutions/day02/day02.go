// Solution for day02 of the Advent of Code Challenge 2023
package day02

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	p1, p2 := solve(input)
	if part == 1 {
		return p1
	} else if part == 2 {
		return p2
	} else {
		util.LogFatal("invalid part")
		return -1
	}
}

func solve(inputData string) (int, int) {
	max := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	total := 0
	powerTotal := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		line := scanner.Text()
		id, err := extractGameID(line)
		if err != nil {
			util.LogFatal("unable to extract game id")
		}

		possible, power, err := checkGamePossible(line, max)
		if err != nil {
			util.LogFatal("invalid game: %s", line)
		}
		if possible {
			total += id
		}
		powerTotal += power
	}
	return total, powerTotal
}

func extractGameID(line string) (int, error) {
	res := strings.Split(line, ":")[0]
	id := strings.Split(res, " ")[1]
	return strconv.Atoi(id)
}

func checkGamePossible(line string, max map[string]int) (possible bool, power int, err error) {
	possible = true
	minCubes := map[string]int{"blue": 0, "green": 0, "red": 0}

	gamesLine := strings.SplitAfter(line, ": ")[1]
	games := strings.Split(gamesLine, "; ")
	for _, game := range games {
		pulls := strings.Split(game, ", ")
		for _, pull := range pulls {
			vals := strings.Split(pull, " ")
			num, err := strconv.Atoi(vals[0])

			if err != nil {
				return false, 0, err
			}
			val, exists := max[vals[1]]
			if !exists || val < num {
				possible = false
			}

			if currMax := minCubes[vals[1]]; currMax < num {
				minCubes[vals[1]] = num
			}
		}
	}
	power = minCubes["blue"] * minCubes["green"] * minCubes["red"]
	return
}
