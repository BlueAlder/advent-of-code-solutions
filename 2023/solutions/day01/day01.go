// Solution for day01 of the Advent of Code Challenge 2023
package day01

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	part1 := calculateCalibration(false)
	part2 := calculateCalibration(true)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

var DIGIT_WORDS = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func isAsciiNumber[K rune | byte](c K) bool {
	return c >= 48 && c <= 57
}

func startsWithDigitWord(str string) int {
	for idx, word := range DIGIT_WORDS {
		if strings.HasPrefix(str, word) {
			return idx
		}
	}
	return -1
}

func endsWithDigitWord(str string) int {
	for idx, word := range DIGIT_WORDS {
		if strings.HasSuffix(str, word) {
			return idx
		}
	}
	return -1
}

func calculateCalibration(part2 bool) (total int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		val := ""

		// Starting number
		for idx, c := range line {
			if isAsciiNumber(c) {
				val = string(c)
				break
			}

			if part2 {
				if dig := startsWithDigitWord(line[idx:]); dig != -1 {
					val = strconv.Itoa(dig)
					break
				}
			}
		}

		// Last number
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if isAsciiNumber(c) {
				val += string(c)
				break
			}

			if part2 {
				if dig := endsWithDigitWord(line[:i+1]); dig != -1 {
					val += strconv.Itoa(dig)
					break
				}
			}
		}

		parsedVal, _ := strconv.Atoi(val)
		total += parsedVal
	}
	return
}
