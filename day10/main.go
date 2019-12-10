package main

import (
	"aoc19/utils"
	"fmt"
)

func main() {
	part1()
	part2()
}

type point struct {
	x, y int
}

func getInput() map[point]bool {
	lines := utils.ReadLines("input.txt")
	asteroids := make(map[point]bool)
	for x, line := range lines {
		for y, c := range line {
			if c == '#' {
				asteroids[point{x, y}] = true
			}
		}
	}
	return asteroids
}

func part1() {
	asteroids := getInput()
	var asteroidsArr []point
	for k := range asteroids {
		asteroidsArr = append(asteroidsArr, k)
	}

	counts := make(map[point]int)

	for i, a1 := range asteroidsArr {
		for j := i + 1; j < len(asteroidsArr); j++ {
			a2 := asteroidsArr[j]
			if !isBlocked(asteroids, a1, a2) {
				counts[a1]++
				counts[a2]++
			}
		}
	}

	max := 0
	var best point
	for k, v := range counts {
		if v > max {
			max = v
			best = k
		}
	}

	fmt.Println(max, best.x, best.y)
}

// returns true if there is an asteroid on the direct way between a and b
func isBlocked(asteroids map[point]bool, a, b point) bool {
	difX := b.x - a.x
	difY := b.y - a.y

	var gcd int
	if difX == 0 || difY == 0 {
		gcd = utils.Max(utils.Abs(difX), utils.Abs(difY))
	} else {
		gcd = utils.Gcd(utils.Abs(difX), utils.Abs(difY))
	}
	difXX := difX / gcd
	difYY := difY / gcd

	for {
		a.x += difXX
		a.y += difYY
		if a == b {
			return false
		}
		if asteroids[a] {
			return true
		}
	}
}

func part2() {
}
