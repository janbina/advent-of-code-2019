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
	program := getInput()

	fmt.Println(findMaxSignal(program, []int{0, 1, 2, 3, 4}))
}

func part2() {
	program := getInput()

	fmt.Println(findMaxSignal(program, []int{5, 6, 7, 8, 9}))
}

func findMaxSignal(program map[int64]int64, settings []int) int64 {
	var maxSignal int64

	utils.WithPermutation(settings, func(p []int) {
		chans := []chan int64{}
		for range p {
			chans = append(chans, make(chan int64))
		}
		done := make(chan struct{})
		for i, v := range p {
			go common.RunIntcode(program, chans[i], chans[(i+1)%len(p)], done)
			chans[i] <- int64(v)
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
