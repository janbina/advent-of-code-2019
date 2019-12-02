package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadNumbers(filename string) []int {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, ToInt(scanner.Text()))
	}
	return numbers
}

func ReadNumbers2(filename string) []int {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		for _, num := range nums {
			numbers = append(numbers, ToInt(num))
		}
	}
	return numbers
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}
