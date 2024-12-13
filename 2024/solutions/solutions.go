package solutions

import (
	"fmt"
	"time"

	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day01"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day02"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day03"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day04"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day05"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day06"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day07"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day08"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day09"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day10"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day11"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day12"
	"github.com/BlueAlder/advent-of-code-solutions/2024/solutions/day13"
	util "github.com/BlueAlder/advent-of-code-solutions/common/utils"
)

var slns = map[int]interface{}{
	1:  day01.Solve,
	2:  day02.Solve,
	3:  day03.Solve,
	4:  day04.Solve,
	5:  day05.Solve,
	6:  day06.Solve,
	7:  day07.Solve,
	8:  day08.Solve,
	9:  day09.Solve,
	10: day10.Solve,
	11: day11.Solve,
	12: day12.Solve,
	13: day13.Solve,
}

const TARGET_TIME = 10 * time.Millisecond

func Run(day int, part int) {

	fmt.Printf("Running solution for day %d part %d\n", day, part)
	v, ext := slns[day]
	if !ext {
		util.LogWarn("Day %d has not been implemented yet", day)
		return
	}

	startTime := time.Now()
	answer := v.(func(int) int)(part)
	elapsedTime := time.Since(startTime)

	util.LogGood("Part %d: %d", part, answer)

	if elapsedTime > TARGET_TIME {
		util.LogWarn("Solution took %s to run.\n", elapsedTime)
	} else {
		fmt.Printf("Solution took "+util.ColorGreen+"%s"+util.ColorReset+" to run.\n\n", elapsedTime)
	}
	// submit(day, part, answer)

}

// func submit(day int, part int, answer int) {
// 	endpoint := fmt.Sprintf("https://adventofcode.com/2023/day/%d/answer", day)

// 	data := url.Values{}
// 	data.Add("level", strconv.Itoa(part))
// 	data.Add("answer", strconv.Itoa(answer))

// 	token := os.Getenv("AOC_SESSION")

// 	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
// 	_ = http.Cookie{
// 		Name:     "session",
// 		Value:    token,
// 		Path:     "/",
// 		MaxAge:   3600,
// 		Secure:   true,
// 		HttpOnly: true,
// 		SameSite: http.SameSiteLaxMode,
// 	}

// 	// req.AddCookie()
// 	if err != nil {
// 		util.LogFatal("Unable to submit solution", err)
// 	}

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		util.LogFatal("Unable to read body")
// 	}
// 	fmt.Printf(string(resBody))

// }
