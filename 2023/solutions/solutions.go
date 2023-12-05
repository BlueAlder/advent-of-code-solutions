package solutions

import (
	"fmt"
	"time"

	"github.com/BlueAlder/advent-of-code-solutions/solutions/day01"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day02"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day03"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day04"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var slns = map[int]interface{}{
	1: day01.Solve,
	2: day02.Solve,
	3: day03.Solve,
	4: day04.Solve,
}

func Run(day int, part int) {

	fmt.Printf("Running solution for day %d part %d\n", day, part)
	v, ext := slns[day]
	if !ext {
		util.LogFatal("day does not exist in function map")
	}

	startTime := time.Now()
	answer := v.(func(int) int)(part)
	elapsedTime := time.Since(startTime)

	util.LogGood("Part %d: %d", part, answer)
	fmt.Printf("Solution took %s to run.\n\n", elapsedTime)

}
