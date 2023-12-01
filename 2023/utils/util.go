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

func MapSlice[T, U any](slice []T, f func(T) U) []U {
	mapped := make([]U, len(slice))
	for i := range slice {
		mapped[i] = f(slice[i])
	}
	return mapped
}
