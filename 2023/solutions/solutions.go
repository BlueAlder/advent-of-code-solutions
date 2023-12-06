package solutions

import (
	"fmt"
	"time"

	"github.com/BlueAlder/advent-of-code-solutions/solutions/day01"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day02"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day03"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day04"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day05"
	"github.com/BlueAlder/advent-of-code-solutions/solutions/day06"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var slns = map[int]interface{}{
	1: day01.Solve,
	2: day02.Solve,
	3: day03.Solve,
	4: day04.Solve,
	5: day05.Solve,
	6: day06.Solve,
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
