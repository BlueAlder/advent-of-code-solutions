// Solution for day14 of the Advent of Code Challenge 2024
package day14

import (
	_ "embed"
	"fmt"
	"image"
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

type robot struct {
	position image.Point
	dx       int
	dy       int
}

func parseInput(input string) []*robot {
	robots := []*robot{}
	digitRe := regexp.MustCompile(`-*\d+`)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		digitsStr := digitRe.FindAllString(line, 4)
		digits := util.MapSlice(digitsStr, util.MustAtoi)
		robots = append(robots, &robot{
			position: image.Pt(digits[0], digits[1]),
			dx:       digits[2],
			dy:       digits[3],
		})
	}
	return robots
}

func part1(inputData string) int {
	robots := parseInput(inputData)
	gridWidth := 101
	gridHeight := 103
	seconds := 100

	// printGrid(robots, gridWidth, gridHeight)

	for _, r := range robots {
		r.move(seconds, gridWidth, gridHeight)
	}
	// Count how many in each quadrant
	total := 1

	quadrants := make(map[int]int)
	for _, r := range robots {
		if r.position.X < gridWidth/2 && r.position.Y < gridHeight/2 {
			quadrants[0]++
		} else if r.position.X > gridWidth/2 && r.position.Y < gridHeight/2 {
			quadrants[1]++
		} else if r.position.X > gridWidth/2 && r.position.Y > gridHeight/2 {
			quadrants[2]++
		} else if r.position.X < gridWidth/2 && r.position.Y > gridHeight/2 {
			quadrants[3]++
		}
	}

	for _, v := range quadrants {
		total *= v
	}

	// printGrid(robots, gridWidth, gridHeight)
	return total
}

func part2(inputData string) int {
	robots := parseInput(inputData)
	gridWidth := 101
	gridHeight := 103

	// printGrid(robots, gridWidth, gridHeight)
	second := 0
	for {
		second++
		for _, r := range robots {
			r.move(1, gridWidth, gridHeight)
		}
		if !checkOverlap(robots) {
			break
		}
	}
	// printGrid(robots, gridWidth, gridHeight)
	return second
}

func (r *robot) move(seconds int, width int, height int) {
	newW := util.Mod((r.position.X + (r.dx * seconds)), width)
	newH := util.Mod((r.position.Y + (r.dy * seconds)), height)
	r.position = image.Pt(newW, newH)
}

func checkOverlap(robots []*robot) bool {
	positions := make(map[image.Point]int)
	for _, r := range robots {
		if _, ok := positions[r.position]; ok {
			return true
		}
		positions[r.position] = 1
	}
	return false
}

func printGrid(robots []*robot, width int, height int) {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, r := range robots {
		if grid[r.position.Y][r.position.X] == '.' {
			grid[r.position.Y][r.position.X] = '1'
		} else {
			grid[r.position.Y][r.position.X] += 1
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}
