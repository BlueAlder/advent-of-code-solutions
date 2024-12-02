package solutions

import (
	"fmt"
	"time"

	util "github.com/BlueAlder/advent-of-code-solutions/common/utils"
)

var slns = map[int]interface{}{}

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