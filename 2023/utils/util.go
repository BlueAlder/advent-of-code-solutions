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

func MapSliceWithError[T, U any](slice []T, f func(T) (U, error)) ([]U, error) {
	mapped := make([]U, len(slice))
	for i := range slice {
		res, err := f(slice[i])
		if err != nil {
			return nil, err
		}
		mapped[i] = res
	}
	return mapped, nil
}

func ReduceSlice[T, U any](slice []T, f func(T, U) U) U {
	var accumulator U
	for i := range slice {
		accumulator = f(slice[i], accumulator)
	}
	return accumulator
}

func IsAsciiNumber[K rune | byte](c K) bool {
	return c >= 48 && c <= 57
}
