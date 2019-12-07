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
	program := getInput()

	fmt.Println(findMaxSignal(program, []int{0, 1, 2, 3, 4}))
}

func part2() {
	program := getInput()

	fmt.Println(findMaxSignal(program, []int{5, 6, 7, 8, 9}))
}

func findMaxSignal(program []int, settings []int) int {
	maxSignal := 0

	utils.WithPermutation(settings, func(p []int) {
		chans := []chan int{}
		for range p {
			chans = append(chans, make(chan int))
		}
		done := make(chan struct{})
		for i, v := range p {
			go common.RunIntcode(utils.CopyInts(program), chans[i], chans[(i+1)%len(p)], done)
			chans[i] <- v
		}
		chans[0] <- 0
		<-done

		out := <-chans[0]
		if out > maxSignal {
			maxSignal = out
		}
	})

	return maxSignal
}
