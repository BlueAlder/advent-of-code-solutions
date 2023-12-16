// Solution for day16 of the Advent of Code Challenge 2023
package day16

import (
	_ "embed"
	"image"
	"strings"
	"sync"

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

type Ray struct {
	X  int
	Y  int
	dx int
	dy int
}

func (r *Ray) Move() {
	r.X += r.dx
	r.Y += r.dy
}

func isInBounds(r *Ray, maxX, maxY int) bool {
	return r.X >= 0 && r.X < maxX && r.Y >= 0 && r.Y < maxY
}

func part1(inputData string) int {
	rows := strings.Split(inputData, "\n")
	initRay := &Ray{X: -1, Y: 0, dx: 1, dy: 0}
	return calculateActivated(rows, initRay)
}

func part2(inputData string) int {
	rows := strings.Split(inputData, "\n")
	max := 0
	for y := range rows {
		r1 := &Ray{
			X:  -1,
			Y:  y,
			dx: 1,
			dy: 0,
		}
		r2 := &Ray{
			X:  len(rows[0]),
			Y:  y,
			dx: -1,
			dy: 0,
		}
		r1m := calculateActivated(rows, r1)
		r2m := calculateActivated(rows, r2)
		max = util.Max(util.Max(max, r2m), r1m)
	}

	for x := range rows[0] {
		r1 := &Ray{
			X:  x,
			Y:  -1,
			dx: 0,
			dy: 1,
		}
		r2 := &Ray{
			X:  x,
			Y:  len(rows),
			dx: 0,
			dy: -1,
		}
		r1m := calculateActivated(rows, r1)
		r2m := calculateActivated(rows, r2)
		max = util.Max(util.Max(max, r2m), r1m)
	}
	return max
}
func calculateActivated(rows []string, initRay *Ray) int {

	seenMutex := &sync.RWMutex{}
	seen := make(sets.Set[Ray])

	vistedChan := make(chan image.Point)
	var visted []image.Point

	var wg sync.WaitGroup
	var followRay func(*Ray)
	followRay = func(r *Ray) {
		defer wg.Done()

		for {
			r.Move()
			seenMutex.RLock()
			notSeen := !seen.Has(*r)
			seenMutex.RUnlock()

			if isInBounds(r, len(rows[0]), len(rows)) && notSeen {
				seenMutex.Lock()
				seen.Add(*r)
				seenMutex.Unlock()

				vistedChan <- image.Point{X: r.X, Y: r.Y}

				block := rows[r.Y][r.X]
				if block == '|' && r.dy == 0 {
					r.dx = 0
					r.dy = 1
					newR := *r
					newR.dy = -r.dy
					wg.Add(1)

					go followRay(&newR)
				} else if block == '-' && r.dx == 0 {
					r.dy = 0
					r.dx = 1
					newR := *r
					newR.dx = -r.dx
					wg.Add(1)

					go followRay(&newR)
				} else if block == '/' {
					r.dx, r.dy = -r.dy, -r.dx
				} else if block == '\\' {
					r.dx, r.dy = r.dy, r.dx
				}
			} else {
				return
			}
		}

	}
	wg.Add(1)

	go followRay(initRay)
	go func() {
		wg.Wait()
		close(vistedChan)
	}()

	for p := range vistedChan {
		visted = append(visted, p)
	}
	return len(visted)
}
