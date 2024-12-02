// Solution for day18 of the Advent of Code Challenge 2023
package day18

import (
	_ "embed"
	"image"
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

type Direction struct {
	dx int
	dy int
}

var direction = map[string]Direction{
	"R": {dx: 1, dy: 0},
	"L": {dx: -1, dy: 0},
	"D": {dx: 0, dy: 1},
	"U": {dx: 0, dy: -1},
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")

	dug := 0
	var vertices []image.Point
	curr := image.Point{X: 0, Y: 0}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := direction[parts[0]]
		num, _ := strconv.Atoi(parts[1])
		for i := 0; i < num; i++ {
			curr.X += dir.dx
			curr.Y += dir.dy
			dug++
		}
		vertices = append(vertices, curr)
	}
	s := shoelaceTheorem(vertices)
	i := s + (dug / 2) + 1

	return i
}

func part2(inputData string) int {
	lines := strings.Split(inputData, "\n")

	dirLookup := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}

	dug := 0
	var vertices []image.Point
	curr := image.Point{X: 0, Y: 0}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := direction[dirLookup[string(parts[2][len(parts[2])-2])]]
		num, _ := strconv.ParseInt(parts[2][2:len(parts[2])-2], 16, 0)
		curr.X += dir.dx * int(num)
		curr.Y += dir.dy * int(num)
		vertices = append(vertices, curr)
		dug += int(num)
	}

	s := shoelaceTheorem(vertices)
	// Picks theorem https://en.wikipedia.org/wiki/Pick%27s_theorem
	i := s + (dug / 2) + 1
	return i
}

func shoelaceTheorem(verts []image.Point) int {
	a := 0
	for i := 0; i < len(verts); i++ {
		a += verts[i].X * (verts[mod(i+1, len(verts))].Y - verts[mod(i-1, len(verts))].Y)
	}
	res := a / 2
	return res
}

func mod(a, b int) int {
	return (a%b + b) % b
}
