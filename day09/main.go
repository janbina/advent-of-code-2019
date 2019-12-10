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

func getInput() map[int64]int64 {
	lines := utils.ReadLines("input.txt")
	strings := strings.Split(lines[0], ",")
	ints := utils.StringsToInts64(strings)
	return utils.IntSliceToMap(ints)
}

func part1() {
	ints := getInput()

	out := common.RunIntcodeSimple(ints, []int64{1})
	fmt.Println(out[0])
}

func part2() {
	ints := getInput()

	out := common.RunIntcodeSimple(ints, []int64{2})
	fmt.Println(out[0])
}
