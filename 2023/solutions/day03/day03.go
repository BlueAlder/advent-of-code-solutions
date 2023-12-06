// Solution for day03 of the Advent of Code Challenge 2023
package day03

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
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

type Set map[int]struct{}

func part1(inputData string) (total int) {
	lines := strings.Split(inputData, "\n")
	rowLength := len(lines[0])
	data := strings.Join(lines, "")

	// 1. Find idx of all symbols
	reSymbol := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)
	symbols := reSymbol.FindAllStringIndex(data, -1)

	// 2. Find all numbers, their indexes, and convert to ints
	partIdxs, partValues := findAllPossibleParts(data)

	// For each symbol add all adjacent parts to Set
	totalParts := make(Set)
	for _, star := range symbols {
		parts := getAdjacentPartIndexes(star[0], rowLength, partIdxs)
		for k := range parts {
			totalParts[k] = struct{}{}
		}
	}
	// Sum all valid parts
	for partIdx := range totalParts {
		total += partValues[partIdx]
	}
	return total
}

func part2(inputData string) (total int) {
	lines := strings.Split(inputData, "\n")
	rowLength := len(lines[0])
	data := strings.Join(lines, "")

	// 1. Find idx of all * (gears)
	reStar := regexp.MustCompile(`\*`)
	stars := reStar.FindAllStringIndex(data, -1)

	// 2. Find all possible parts and convert to ints
	partIdxs, partValues := findAllPossibleParts(data)

	// 3. For each gear check surrounding for part
	for _, star := range stars {
		parts := getAdjacentPartIndexes(star[0], rowLength, partIdxs)
		// 4. If exactly 2, lookup corresponding parts and multiply them for ratio
		if len(parts) == 2 {
			ratio := 1
			for k := range parts {
				ratio *= partValues[k]
			}
			// 5. Sum total
			total += ratio
		}
	}
	return
}

func findAllPossibleParts(data string) (partIndexes [][]int, partValues []int) {
	reDigit := regexp.MustCompile(`\d+`)
	partIndexes = reDigit.FindAllStringIndex(data, -1)

	resString := reDigit.FindAllString(data, -1)
	partValues, err := util.MapSliceWithError(resString, strconv.Atoi)
	if err != nil {
		util.LogFatal("error converting to ints: %w", err)
	}
	return
}

func getAdjacentPartIndexes(symbolIdx int, rowLength int, partIndexes [][]int) Set {
	surround := [8]int{
		-rowLength - 1,
		-rowLength,
		-rowLength + 1,
		-1, 1,
		rowLength - 1,
		rowLength,
		rowLength + 1,
	}
	// Use a set here to combat duplicates
	parts := make(Set)
	for _, dx := range surround {
		c := symbolIdx + dx
		// Check if we are in a number
		for idx, digitIdx := range partIndexes {
			if c >= digitIdx[0] && c < digitIdx[1] {
				parts[idx] = struct{}{}
			}
		}
	}
	return parts
}
