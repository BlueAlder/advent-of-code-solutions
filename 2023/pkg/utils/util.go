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

func HammingDistanceString(s1 string, s2 string) (distance int) {
	if len(s1) != len(s2) {
		return -1
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return
}
