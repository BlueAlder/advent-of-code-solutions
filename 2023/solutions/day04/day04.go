// Solution for day04 of the Advent of Code Challenge 2023
package day04

import (
	_ "embed"
	"regexp"
	"slices"
	"strconv"
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

func part1(inputData string) (total int) {
	cards := strings.Split(inputData, "\n")

	for _, card := range cards {
		winning, picked := getWinningAndPickedNumbers(card)
		matches := calcCardMatches(winning, picked)
		value := 1 << matches >> 1
		total += value
	}
	return
}

func part2(inputData string) (total int) {
	cards := strings.Split(inputData, "\n")
	cardCount := make([]int, len(cards))

	for idx, card := range cards {
		winning, picked := getWinningAndPickedNumbers(card)
		matches := calcCardMatches(winning, picked)
		copiesOfCurrent := cardCount[idx]
		for i := 1; i <= matches; i++ {
			cardCount[idx+i] += copiesOfCurrent + 1
		}
	}
	total = len(cards)
	for _, v := range cardCount {
		total += v
	}
	return
}

func calcCardMatches(winning []int, picked []int) (matches int) {
	for _, win := range winning {
		if slices.Contains(picked, win) {
			matches++
		}
	}
	return
}

func getWinningAndPickedNumbers(line string) ([]int, []int) {
	reNum := regexp.MustCompile(`\d+`)
	parts := strings.Split(line, "|")

	winningStr := reNum.FindAllString(parts[0], -1)[1:]
	winning, _ := util.MapSliceWithError(winningStr, strconv.Atoi)

	pickedStr := reNum.FindAllString(parts[1], -1)
	picked, _ := util.MapSliceWithError(pickedStr, strconv.Atoi)

	return winning, picked
}
