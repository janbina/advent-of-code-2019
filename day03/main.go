package main

import (
	"fmt"
	"github.com/janbina/advent-of-code-2019/utils"
	"math"
	"strings"
)

type instruction struct {
	dir byte
	len int
}

type point struct {
	x, y int
}

func main() {
	part12()
}

func getInput() ([]instruction, []instruction) {
	lines := utils.ReadLines("input.txt")
	var inst1 []instruction
	var inst2 []instruction

	for _, inst := range strings.Split(lines[0], ",") {
		inst1 = append(inst1, instruction{inst[0], utils.ToInt(inst[1:])})
	}

	for _, inst := range strings.Split(lines[1], ",") {
		inst2 = append(inst2, instruction{inst[0], utils.ToInt(inst[1:])})
	}

	return inst1, inst2
}

func getPoints(inst []instruction) map[point]int {
	m := make(map[point]int)

	var x, y, steps int

	d := map[byte]point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}

	for _, i := range inst {
		for j := 0; j < i.len; j++ {
			x += d[i.dir].x
			y += d[i.dir].y
			steps++
			m[point{x, y}] = steps
		}
	}

	return m
}

func part12() {
	i1, i2 := getInput()
	p1, p2 := getPoints(i1), getPoints(i2)

	var minDist, minLen int = math.MaxInt32, math.MaxInt32

	for k, v := range p1 {
		if v2, ok := p2[k]; ok {
			dist := utils.ManhattanDist(k.x, 0, k.y, 0)
			minDist = utils.Min(minDist, dist)

			len := v + v2
			minLen = utils.Min(minLen, len)
		}
	}

	fmt.Println("Part 1:", minDist)
	fmt.Println("Part 2:", minLen)
}
