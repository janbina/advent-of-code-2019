package main

import (
	"aoc19/common"
	"aoc19/utils"
	"fmt"
	"strings"
)

func main() {
	part12()
}

type point struct {
	x, y int
}

func (p point) plus(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func getInput() map[int64]int64 {
	lines := utils.ReadLines("input.txt")
	strings := strings.Split(lines[0], ",")
	ints := utils.StringsToInts64(strings)
	return utils.IntSliceToMap(ints)
}

const north = 1
const south = 2
const west = 3
const east = 4

var leftTurn = [5]int{-1, west, east, south, north}
var rightTurn = [5]int{-1, east, west, north, south}
var dx = [5]point{point{0, 0}, point{0, 1}, point{0, -1}, point{-1, 0}, point{1, 0}}

const empty = 1
const wall = 2

const hitWall = 0
const moved = 1
const foundOxygen = 2

func turnLeft(prev int) int {
	return leftTurn[prev]
}

func turnRight(prev int) int {
	return rightTurn[prev]
}

func getNextPos(current point, direction int) point {
	return current.plus(dx[direction])
}

func (p point) index() int {
	return p.y*50 + p.x
}

func part12() {
	program := getInput()

	in := make(chan int64)
	out := make(chan int64)
	done := make(chan struct{})

	go common.RunIntcode(program, in, out, done)

	maze := make([]int, 50*50)
	start := point{25, 25}
	current := point{25, 25}
	oxygen := point{-1, -1}
	maze[current.index()] = empty
	nextMove := north

	for {
		in <- int64(nextMove)
		nextPos := getNextPos(current, nextMove)
		switch <-out {
		case hitWall:
			maze[nextPos.index()] = wall
			nextMove = turnLeft(nextMove)
		case moved:
			current = nextPos
			maze[current.index()] = empty
			nextMove = turnRight(nextMove)
		case foundOxygen:
			current = nextPos
			oxygen = current
			maze[current.index()] = empty
			nextMove = turnRight(nextMove)
		}

		if oxygen.x != -1 && current == start {
			break
		}
	}

	fmt.Println(dfs(maze, start, oxygen))
}

type item struct {
	p    point
	dist int
}

func dfs(maze []int, start point, oxygen point) (int, int) {
	var queue []item
	discovered := make(map[point]bool)
	queue = append(queue, item{oxygen, 0})
	max := 0
	startDist := -1

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if discovered[top.p] {
			continue
		}
		discovered[top.p] = true
		if top.p == start && startDist < 0 {
			startDist = top.dist
		}
		if maze[top.p.index()] == empty {
			queue = append(queue, item{point{top.p.x + 1, top.p.y}, top.dist + 1})
			queue = append(queue, item{point{top.p.x - 1, top.p.y}, top.dist + 1})
			queue = append(queue, item{point{top.p.x, top.p.y + 1}, top.dist + 1})
			queue = append(queue, item{point{top.p.x, top.p.y - 1}, top.dist + 1})
			max = utils.Max(max, top.dist)
		}
	}

	return startDist, max
}
