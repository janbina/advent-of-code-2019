package main

import (
	"fmt"
	"github.com/janbina/advent-of-code-2019/common"
	"github.com/janbina/advent-of-code-2019/utils"
	"strings"
)

func main() {
	part1()
	part2()
}

func getInput() []int64 {
	lines := utils.ReadLines("input.txt")
	strings := strings.Split(lines[0], ",")
	return utils.StringsToInts64(strings)
}

func part1() {
	ints := getInput()
	mem := utils.IntSliceToMap(ints)

	ints[1] = 12
	ints[2] = 2

	common.RunIntcodeSimple(mem, []int64{})

	fmt.Println(mem[0])
}

func part2() {
	ints := getInput()

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			ints[1] = int64(noun)
			ints[2] = int64(verb)
			mem := utils.IntSliceToMap(ints)
			common.RunIntcodeSimple(mem, []int64{})
			if mem[0] == 19690720 {
				fmt.Println("Noun =", noun, "verb =", verb, "answer =", noun*100+verb)
				return
			}
		}
	}
}
