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

	ints[1] = 12
	ints[2] = 2

	done := make(chan struct{})
	go common.RunIntcode(ints, nil, nil, done)
	<-done

	fmt.Println("Output =", ints[0])
}

func part2() {
	ints := getInput()

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			mem := utils.CopyInts(ints)
			mem[1] = noun
			mem[2] = verb
			done := make(chan struct{})
			go common.RunIntcode(mem, nil, nil, done)
			<-done
			if mem[0] == 19690720 {
				fmt.Println("Noun =", noun, "verb =", verb, "answer =", noun*100+verb)
				return
			}
		}
	}
}
