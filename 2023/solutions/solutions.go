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
	if v, ext := slns[day]; ext {
		v.(func())()
		return
	}

	util.LogFatal("day does not exist in function map")
}
