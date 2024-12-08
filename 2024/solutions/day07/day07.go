// Solution for day07 of the Advent of Code Challenge 2024
package day07

import (
	_ "embed"
	"fmt"
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

type equation struct {
	goal    int
	numbers []int
}

func parseInput(input string) []equation {
	lines := strings.Split(input, "\n")
	var equations []equation
	max := 0
	for _, line := range lines {
		components := strings.Split(line, ": ")
		goal := util.MustAtoi(components[0])
		numberString := strings.Fields(components[1])
		numbers := util.MapSlice(numberString, util.MustAtoi)
		if len(numbers) > max {
			max = len(numbers)
		}
		equations = append(equations, equation{goal, numbers})
	}
	return equations
}

func part1(inputData string) int {
	equations := parseInput(inputData)
	sum := 0

	for _, equation := range equations {
		if canFormGoal(equation, 1, equation.numbers[0], false) {
			sum += equation.goal
		}
	}
	return sum
}

func part2(inputData string) int {
	equations := parseInput(inputData)
	sum := 0
	for _, equation := range equations {
		if canFormGoal(equation, 1, equation.numbers[0], true) {
			sum += equation.goal
		}
	}
	return sum
}

func canFormGoal(equation equation, index int, current int, concat bool) bool {
	if index == len(equation.numbers) {
		return current == equation.goal
	}
	if current > equation.goal {
		return false
	}

	return canFormGoal(equation, index+1, current+equation.numbers[index], concat) ||
		canFormGoal(equation, index+1, current*equation.numbers[index], concat) ||
		(concat && canFormGoal(equation, index+1, util.MustAtoi(fmt.Sprintf("%d%d", current, equation.numbers[index])), concat))
}
