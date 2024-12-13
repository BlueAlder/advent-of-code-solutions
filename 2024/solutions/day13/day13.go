// Solution for day13 of the Advent of Code Challenge 2024
package day13

import (
	_ "embed"
	"regexp"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/common/utils"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		util.LogFatal("invalid part number")
		return -1
	}
}

type button struct {
	dx int
	dy int
}

type game struct {
	buttonA button
	buttonB button
	prize   struct {
		x int
		y int
	}
}

func parseInput(input string) []game {
	games := []game{}
	gameLines := strings.Split(input, "\n\n")
	digitRe := regexp.MustCompile(`\d+`)

	for _, line := range gameLines {
		lines := strings.Split(line, "\n")
		buttonA := util.MapSlice(digitRe.FindAllString(lines[0], 2), util.MustAtoi)
		buttonB := util.MapSlice(digitRe.FindAllString(lines[1], 2), util.MustAtoi)
		prize := util.MapSlice(digitRe.FindAllString(lines[2], 2), util.MustAtoi)
		games = append(games, game{
			buttonA: button{dx: buttonA[0], dy: buttonA[1]},
			buttonB: button{dx: buttonB[0], dy: buttonB[1]},
			prize:   struct{ x, y int }{x: prize[0], y: prize[1]},
		})
	}
	return games
}

func part1(inputData string) int {
	games := parseInput(inputData)

	total := 0
	for _, g := range games {
		aPresses, bPresses := g.findNumberOfPresses(false)
		if aPresses == -1 || bPresses == -1 {
			continue
		}
		total += (aPresses * 3) + bPresses
	}
	return total
}

func part2(inputData string) int {
	games := parseInput(inputData)
	total := 0
	for _, g := range games {
		aPresses, bPresses := g.findNumberOfPresses(true)
		if aPresses == -1 || bPresses == -1 {
			continue
		}
		total += (aPresses * 3) + bPresses
	}
	return total
}

func (g game) findNumberOfPresses(p2 bool) (int, int) {
	if p2 {
		g.prize.x += 10000000000000
		g.prize.y += 10000000000000
	}
	aPresses := float64((g.prize.x*g.buttonB.dy)-(g.buttonB.dx*g.prize.y)) / float64(((g.buttonA.dx * g.buttonB.dy) - (g.buttonB.dx * g.buttonA.dy)))
	if float64(int(aPresses)) != aPresses {
		return -1, -1
	}
	bPresses := float64(g.prize.y-(g.buttonA.dy*int(aPresses))) / float64(g.buttonB.dy)
	if float64(int(bPresses)) != bPresses {
		return -1, -1
	}

	return int(aPresses), int(bPresses)
}
