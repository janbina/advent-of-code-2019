package main

import (
	"aoc19/common"
	"aoc19/utils"
	"fmt"
	"math"
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

type point struct {
	x, y int
}

type robot struct {
	position point
	facing   int
}

func (r *robot) turn(direction int) {
	if direction == 0 {
		r.facing = (r.facing + 3) % 4
	} else {
		r.facing = (r.facing + 5) % 4
	}
	switch r.facing {
	case 0:
		r.position.y++
	case 1:
		r.position.x++
	case 2:
		r.position.y--
	case 3:
		r.position.x--
	}
}

func runRobot(program map[int64]int64, initPaint byte) map[point]byte {
	// capacity = 1 is important, because we will send one more value
	// before we can realise no more output will come
	in := make(chan int64, 1)
	out := make(chan int64)
	done := make(chan struct{})

	go common.RunIntcode(program, in, out, done)

	colors := make(map[point]byte)
	robot := &robot{point{0, 0}, 0}
	colors[robot.position] = initPaint

	for {
		in <- int64(colors[robot.position])
		color, ok := <-out
		if !ok {
			break
		}
		colors[robot.position] = byte(color)
		robot.turn(int(<-out))
	}
	<-done

	return colors
}

func part1() {
	mem := getInput()

	colors := runRobot(mem, 0)

	fmt.Println(len(colors))
}

func part2() {
	mem := getInput()

	colors := runRobot(mem, 1)

	drawIdentifier(colors)
}

func drawIdentifier(colors map[point]byte) {
	topLeft, bottomRight := point{math.MaxInt32, math.MinInt32}, point{math.MinInt32, math.MaxInt32}
	for k := range colors {
		topLeft.x = utils.Min(topLeft.x, k.x)
		topLeft.y = utils.Max(topLeft.y, k.y)
		bottomRight.x = utils.Max(bottomRight.x, k.x)
		bottomRight.y = utils.Min(bottomRight.y, k.y)
	}
	for y := topLeft.y; y >= bottomRight.y; y-- {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			if colors[point{x, y}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
