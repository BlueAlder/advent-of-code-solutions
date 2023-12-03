// Solution for day03 of the Advent of Code Challenge 2023
package day03

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

//go:embed input.txt
var input string

func Solve() {
	ans1 := part1()
	ans2 := part2()
	fmt.Printf("Part 1: %d\n", ans1)
	fmt.Printf("Part 2: %d\n", ans2)
}

type Set map[int]struct{}

func part2() (total int) {

	lines := strings.Split(input, "\n")
	rowLength := len(lines[0])
	data := strings.Join(lines, "")
	// 1. Find idx of all *
	reStar := regexp.MustCompile(`\*`)
	stars := reStar.FindAllStringIndex(data, -1)

	// 2. Find all numbers
	reDigit := regexp.MustCompile(`\d+`)
	resIdx := reDigit.FindAllStringIndex(data, -1)
	resString := reDigit.FindAllString(data, -1)
	resInts, err := util.MapSliceWithError(resString, strconv.Atoi)
	if err != nil {
		util.LogFatal("error converting to ints %w", err)
	}

	// 3. For each * check surrounding for part
	for _, star := range stars {
		surround := [8]int{
			-rowLength - 1,
			-rowLength,
			-rowLength + 1,
			-1, 1,
			rowLength - 1,
			rowLength,
			rowLength + 1,
		}

		parts := make(Set)
		for _, dx := range surround {
			c := star[0] + dx
			if c < 0 || c > len(data) {
				continue
			}
			// Check digits
			for idx, digitIdx := range resIdx {
				if c >= digitIdx[0] && c < digitIdx[1] {
					parts[idx] = struct{}{}
				}
			}
		}
		// 4. If 2, lookup corresponding digits and multiply them

		ratio := 1
		if len(parts) == 2 {
			for k := range parts {
				ratio *= resInts[k]
			}
			total += ratio
		}
	}
	return
	// 5. Sum total
}

func part1() (total int) {
	plane := strings.Split(input, "\n")
	// strings.ReplaceAll(input, "\n", "")
	symbolMap := generateSymbolMap(plane)
	var validValues []string

	for y, line := range plane {
		reDigit := regexp.MustCompile(`\d+`)
		resIdx := reDigit.FindAllStringIndex(line, -1)
		resString := reDigit.FindAllString(line, -1)
		for i := range resIdx {
			if checkValidNumber(resIdx[i][0], y, resIdx[i][1]-resIdx[i][0], symbolMap) {
				validValues = append(validValues, resString[i])
			}
		}
	}
	ints, err := util.MapSliceWithError(validValues, strconv.Atoi)
	if err != nil {
		panic(fmt.Errorf("could not convert values: %w", err))
	}
	total = util.ReduceSlice(ints, func(el int, total int) int { return el + total })
	return
}

func checkValidNumber(x int, y int, length int, symbolMap [][]bool) bool {
	for iy := y - 1; iy < y+2; iy++ {
		if iy < 0 || iy >= len(symbolMap) {
			continue
		}
		for ix := x - 1; ix < x+length+1; ix++ {
			if ix < 0 || ix >= len(symbolMap[0]) {
				continue
			}
			if symbolMap[iy][ix] {
				return true
			}
		}
	}
	return false
}

func generateSymbolMap(lines []string) [][]bool {
	symbolMap := make([][]bool, len(lines))
	for i := range symbolMap {
		symbolMap[i] = make([]bool, len(lines[0]))
	}

	for y, line := range lines {
		for x, c := range line {
			if !(string(c) == "." || util.IsAsciiNumber(c)) {
				symbolMap[y][x] = true
			}
		}
	}
	return symbolMap
}
