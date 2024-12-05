// Solution for day03 of the Advent of Code Challenge 2024
package day03

import (
	_ "embed"
	"regexp"

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

func part1(inputData string) int {
	return extractMuls(inputData)
}

func part2(inputData string) int {
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	remainingString := inputData
	total := 0
	for len(remainingString) > 0 {
		dontIndex := dontRe.FindStringIndex(remainingString)
		if dontIndex == nil {
			total += extractMuls(remainingString)
			break
		}
		slice := remainingString[:dontIndex[0]]
		remainingString = remainingString[dontIndex[1]:]
		total += extractMuls(slice)
		doIndex := doRe.FindStringIndex(remainingString)
		if doIndex == nil {
			break
		}
		remainingString = remainingString[doIndex[1]:]
	}

	return total
}

func extractMuls(input string) int {
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := mulRe.FindAllString(input, -1)

	digitsRe := regexp.MustCompile(`\d+`)
	total := 0
	for _, match := range matches {
		digitMatches := digitsRe.FindAllString(match, -1)
		digits := util.MapSlice(digitMatches, util.MustAtoi)
		total += digits[0] * digits[1]
	}
	return total
}
