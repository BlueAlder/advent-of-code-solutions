// Solution for day04 of the Advent of Code Challenge 2024
package day04

import (
	_ "embed"
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

type direction struct {
	dx int
	dy int
}

type point struct {
	x int
	y int
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func part1(inputData string) int {
	grid := parseInput(inputData)
	directions := []direction{
		{dx: 1, dy: 1},
		{dx: 0, dy: 1},
		{dx: -1, dy: 1},
		{dx: -1, dy: 0},
		{dx: -1, dy: -1},
		{dx: 0, dy: -1},
		{dx: 1, dy: -1},
		{dx: 1, dy: 0},
	}
	return len(findWordInGrid(grid, directions, "XMAS"))
}

func part2(inputData string) int {
	grid := parseInput(inputData)
	directions := []direction{
		{dx: 1, dy: 1},
		{dx: -1, dy: 1},
		{dx: -1, dy: -1},
		{dx: 1, dy: -1},
	}
	words := findWordInGrid(grid, directions, "MAS")
	// Check how many share a centre point
	midPointCount := map[point]int{}
	for _, word := range words {
		midPoint := word[len(word)/2]
		midPointCount[midPoint]++
	}

	total := 0
	for _, v := range midPointCount {
		if v > 1 {
			total++
		}
	}
	return total
}

func findWordInGrid(grid []string, dirs []direction, word string) [][]point {
	allPoints := [][]point{}
	width := len(grid[0])
	length := len(grid)

	for y, row := range grid {
		for x, char := range row {
			if char != rune(word[0]) {
				continue
			}
		DirectionLoop:
			for _, dir := range dirs {
				points := []point{}
				for i, c := range word {
					newX := x + i*dir.dx
					newY := y + i*dir.dy
					if newX < 0 || newX >= width || newY < 0 || newY >= length {
						continue DirectionLoop
					}
					if grid[newY][newX] != byte(c) {
						continue DirectionLoop
					}
					points = append(points, point{x: newX, y: newY})
				}
				allPoints = append(allPoints, points)
			}
		}
	}
	return allPoints
}
