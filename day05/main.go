package main

import (
	"aoc19/common"
	"aoc19/utils"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func getInput() []int {
	lines := utils.ReadLines("input.txt")
	strings := strings.Split(lines[0], ",")
	return utils.StringsToInts(strings)
}

func part1() {
	ints := getInput()

	output := common.RunIntcode(ints, []int{1})
	fmt.Println(output)
}

func part2() {
	ints := getInput()

	output := common.RunIntcode(ints, []int{5})
	fmt.Println(output)
}
