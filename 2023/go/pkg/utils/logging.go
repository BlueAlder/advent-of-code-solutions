package util

import (
	"fmt"
	"os"
)

var (
	ColorGreen = "\033[32m"
	ColorRed   = "\033[31m"
	ColorReset = "\033[0m"
)

// logFatal prints a message in red and then runs os.Exit(1)
func LogFatal(s string, a ...any) {
	fmt.Printf(ColorRed+s+ColorReset+"\n", a...)
	os.Exit(1)
}

func LogWarn(s string, a ...any) {
	fmt.Printf(ColorRed+s+ColorReset+"\n", a...)
}

func LogGood(s string, a ...any) {
	fmt.Printf(ColorGreen+s+ColorReset+"\n", a...)
}
