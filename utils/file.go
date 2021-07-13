package utils

import (
	"bufio"
	"os"
)

// Readline reads all lines of a file and returns an error if there is any
func Readline(fp string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
