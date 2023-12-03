package util

import (
	"fmt"
	"os"
)

var (
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
)

// logFatal prints a message in red and then runs os.Exit(1)
func LogFatal(s string, a ...any) {
	fmt.Printf(colorRed+s+"\n", a...)
	os.Exit(1)
}

func LogGood(s string, a ...any) {
	fmt.Printf(colorGreen+s, a...)
}
