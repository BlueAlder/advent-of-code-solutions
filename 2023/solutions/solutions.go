package solutions

import (
	"fmt"

	"github.com/BlueAlder/advent-of-code-solutions/solutions/day01"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day02"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day03"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var slns = map[int]interface{}{
	1: day01.Solve,
	2: day02.Solve,
	3: day03.Solve,
}

func Run(day int) {
	fmt.Printf("Running solution for day %d\n", day)
	v, ext := slns[day]
	if !ext {
		util.LogFatal("day does not exist in function map")
	}
	p1, p2 := v.(func() (int, int))()
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)

}
