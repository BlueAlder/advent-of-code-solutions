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
