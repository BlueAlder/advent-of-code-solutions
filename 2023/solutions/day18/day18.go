// Solution for day18 of the Advent of Code Challenge 2023
package day18

import (
	_ "embed"
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/pkg/sets"
	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
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

	dug := make(sets.Set[image.Point])
	var vertices []image.Point
	curr := image.Point{X: 0, Y: 0}
	dug.Add(curr)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := direction[parts[0]]
		num, _ := strconv.Atoi(parts[1])
		for i := 0; i < num; i++ {
			curr.X += dir.dx
			curr.Y += dir.dy
			dug.Add(curr)
		}
		vertices = append(vertices, curr)
	}
	fmt.Println(len(dug))
	// fmt.Println(vertices)
	// area := shoelaceTheorem(vertices)
	area := 0
	area = floodFill(dug, image.Point{X: 1, Y: 1})
	// mnX, mxX, mnY, mxY := findBounds(dug)

	// for y := mnY; y <= mxY; y++ {
	// 	inside := false
	// 	line := ""
	// 	for x := mnX; x <= mxX; x++ {
	// 		if dug.Has(image.Point{X: x, Y: y}) {
	// 			inside = !inside
	// 			line += "#"
	// 		} else if inside {
	// 			area += 1
	// 			line += "#"
	// 		} else {
	// 			line += "."

	// 		}
	// 	}
	// 	fmt.Println(line)
	// }

	return area + len(dug)
}

func part2(inputData string) int {
	return 0
}

func floodFill(dug sets.Set[image.Point], start image.Point) (area int) {
	var queue []image.Point
	queue = append(queue, start)
	visted := make(sets.Set[image.Point])
	var inside []image.Point
	dirs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if !dug.Has(n) && !visted.Has(n) {
			inside = append(inside, n)
			visted.Add(n)

			area += 1
			for _, dir := range dirs {
				p := image.Point{X: n.X + dir[0], Y: n.Y + dir[1]}
				if !(visted.Has(p) || dug.Has(p)) {
					queue = append(queue, p)
				}
			}
		}
	}
	fmt.Println(inside)
	return
}

func shoelaceTheorem(verts []image.Point) int {
	a1 := 0
	a2 := 0
	for i := 0; i < len(verts); i++ {
		fmt.Println(verts[i].X * verts[(i+1)%(len(verts))].Y)
		a1 += verts[i].X * verts[(i+1)%(len(verts))].Y
		a2 += verts[i].Y * verts[(i+1)%(len(verts))].X
	}
	return util.Abs(a1-a2) / 2
}

func findBounds(dug sets.Set[image.Point]) (minX, maxX, minY, maxY int) {
	minX, minY = math.MaxInt, math.MaxInt
	maxX, maxY = -math.MaxInt, -math.MaxInt

	for v := range dug {
		if v.X < minX {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}
	return
}
