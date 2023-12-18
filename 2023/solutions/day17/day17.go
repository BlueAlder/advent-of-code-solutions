// Solution for day17 of the Advent of Code Challenge 2023
package day17

import (
	_ "embed"
	"image"
	"math"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
	"github.com/gammazero/deque"
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
	rows := strings.Split(inputData, "\n")

	visted := make(map[image.Point]int)
	var stack deque.Deque[image.Point]
	start := image.Point{X: 0, Y: 0}
	stack.PushFront(start)
	consecutative := 0
	min := math.MaxInt

	for stack.Len() > 0 {
		i := stack.PopFront()

	}

	// DFS?

	return 0
}

func part2(inputData string) int {
	return 0
}
