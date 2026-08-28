package main

import (
	"fmt"
	"io"
	"os"
)

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("unable to open file - '%s'. %w", filename, err)
	}
	defer file.Close()

	bytes := make([]byte, 2048)
	var read int
	for {
		count, err := file.Read(bytes)
		read += count
		if err != nil {
			if err != io.EOF {
				return 0, fmt.Errorf("unable to read file: %w", err)
			}
			break
		}
	}

	return read, nil
}
