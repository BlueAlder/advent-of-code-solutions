package util

import (
	"fmt"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read file")
		return nil, fmt.Errorf("unable to read file: %w", err)
	}
	return file, nil
}

func ReverseString(input string) (result string) {
	for _, v := range input {
		result = string(v) + result

	}
	return
}

func IsAsciiNumber[K rune | byte](c K) bool {
	return c >= 48 && c <= 57
}
