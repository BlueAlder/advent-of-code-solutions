// Solution for day01 of the Advent of Code Challenge 2023
package day01

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

//go:embed input.txt
var input string

func Solve() (int, int) {
	return calculateCalibration(false), calculateCalibration(true)
}

func isAsciiNumber[K rune | byte](c K) bool {
	return c >= 48 && c <= 57
}

func startsWithDigitWord(str string, words []string) int {
	for idx, word := range words {
		if strings.HasPrefix(str, word) {
			return idx
		}
	}
	return -1
}

func getFirstDigit(line string, digitWords []string, part2 bool) (val string) {
	for idx, c := range line {
		if isAsciiNumber(c) {
			val = string(c)
			return
		}

		if part2 {
			if dig := startsWithDigitWord(line[idx:], digitWords); dig != -1 {
				val = strconv.Itoa(dig)
				return
			}
		}
	}
	return
}

func calculateCalibration(part2 bool) int {
	var DIGIT_WORDS = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	total := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// Get first digit from line
		line := scanner.Text()
		val1 := getFirstDigit(line, DIGIT_WORDS, part2)

		// Reverse line and word digits and check again for the other side
		rev := util.ReverseString(line)
		revd_digits := util.MapSlice(DIGIT_WORDS, util.ReverseString)
		val2 := getFirstDigit(rev, revd_digits, part2)

		// Add and total
		parsedVal, _ := strconv.Atoi(val1 + val2)
		total += parsedVal
	}
	return total
}
