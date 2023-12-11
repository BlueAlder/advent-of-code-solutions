package util

import (
	"fmt"
	"os"
)

var (
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
	colorReset = "\033[0m"
)

// logFatal prints a message in red and then runs os.Exit(1)
func LogFatal(s string, a ...any) {
	fmt.Printf(colorRed+s+colorReset+"\n", a...)
	os.Exit(1)
}

func LogWarn(s string, a ...any) {
	fmt.Printf(colorRed+s+colorReset+"\n", a...)
}

func LogGood(s string, a ...any) {
	fmt.Printf(colorGreen+s+colorReset+"\n", a...)
}
