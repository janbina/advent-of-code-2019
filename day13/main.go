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

type point struct {
	x, y int
}

func getInput() map[int64]int64 {
	lines := utils.ReadLines("input.txt")
	strings := strings.Split(lines[0], ",")
	ints := utils.StringsToInts64(strings)
	return utils.IntSliceToMap(ints)
}

func part1() {
	program := getInput()

	out := common.RunIntcodeSimple(program, []int64{})

	screen := make(map[point]int)

	for i := 0; i < len(out); i += 3 {
		x := int(out[i])
		y := int(out[i+1])
		z := int(out[i+2])
		screen[point{x, y}] = z
	}

	cnt := 0
	for _, v := range screen {
		if v == 2 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func part2() {
	program := getInput()
	program[0] = 2

	in := make(chan int64)
	inRequest := make(chan struct{})
	out := make(chan int64)
	done := make(chan struct{})

	go common.RunIntcodeInRequest(program, in, inRequest, out, done)

	score := 0
	paddleX := 0
	ballX := 0

loop:
	for {
		select {
		case <-inRequest:
			if paddleX < ballX {
				in <- 1
			} else if paddleX > ballX {
				in <- -1
			} else {
				in <- 0
			}
		case x := <-out:
			y := <-out
			z := <-out

			if x == -1 && y == 0 {
				score = int(z)
			} else if z == 3 {
				paddleX = int(x)
			} else if z == 4 {
				ballX = int(x)
			}
		case <-done:
			break loop
		}
	}

	fmt.Println(score)
}
