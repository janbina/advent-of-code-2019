package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	Check(scanner.Err())

	return lines
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func ToInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	Check(err)
	return result
}
