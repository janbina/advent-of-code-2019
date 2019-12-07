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

	in := make(chan int)
	out := make(chan int)
	done := make(chan struct{})
	go common.RunIntcode(ints, in, out, done)
	in <- 1

	fmt.Println("Part 1")
	for v := range out {
		fmt.Println(v)
	}
	<-done
}

func part2() {
	ints := getInput()

	in := make(chan int)
	out := make(chan int)
	done := make(chan struct{})
	go common.RunIntcode(ints, in, out, done)
	in <- 5

	fmt.Println("Part 2")
	for v := range out {
		fmt.Println(v)
	}
	<-done
}
