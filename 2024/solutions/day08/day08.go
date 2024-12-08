// Solution for day08 of the Advent of Code Challenge 2024
package day08

import (
	_ "embed"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/common/sets"
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

type point struct {
	x int
	y int
}

func parseInput(input string) (map[rune][]point, []string) {
	lines := strings.Split(input, "\n")
	points := make(map[rune][]point)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				points[char] = append(points[char], point{x, y})
			}
		}
	}
	return points, lines
}

func part1(inputData string) int {
	antennas, lines := parseInput(inputData)
	antinodes := make(sets.Set[point])
	width := len(lines[0])
	height := len(lines)
	for _, points := range antennas {

		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				dx := points[i].x - points[j].x
				dy := points[i].y - points[j].y
				anti1 := point{points[i].x - dx, points[i].y - dy}
				if anti1.x >= 0 && anti1.x < width && anti1.y >= 0 && anti1.y < height {
					antinodes.Add(anti1)
				}
				anti2 := point{points[i].x + (2 * dx), points[i].y + (2 * dy)}
				if anti2.x >= 0 && anti2.x < width && anti2.y >= 0 && anti2.y < height {
					antinodes.Add(anti2)
				}
			}
		}
	}
	return len(antinodes)
}

func part2(inputData string) int {
	return 0
}
