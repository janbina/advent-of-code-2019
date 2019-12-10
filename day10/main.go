package main

import (
	"aoc19/utils"
	"fmt"
)

func main() {
	part1()
	part2()
}

func getInput() []string {
	lines := utils.ReadLines("input.txt")
	return lines
}

func part1() {
	input := getInput()

	counts := make([]int, len(input)*len(input[0]))

	// for every position [x, y] on the map
	for x, r := range input {
		for y, c := range r {
			// if it is asteroid on that position
			if c == '#' {
				// for every position [x2, y2] on the map
				for x2, r2 := range input {
					for y2, c2 := range r2 {
						// if position [x, y] is not same as [x2, y2] and there is an asteroid
						if (x != x2 || y != y2) && c2 == '#' {
							if !isBlocked(input, x, y, x2, y2) {
								counts[x*len(input)+y]++
							}
						}
					}
				}
			}
		}
	}

	max := 0
	var x, y int
	for i, c := range counts {
		if c > max {
			max = c
			x = i / len(input)
			y = i % len(input)
		}
	}

	fmt.Println(max, x, y)
}

func isBlocked(input []string, x, y, x2, y2 int) bool {
	difX := x2 - x
	difY := y2 - y

	var gcd int
	if difX == 0 || difY == 0 {
		gcd = utils.Max(utils.Abs(difX), utils.Abs(difY))
	} else {
		gcd = utils.Gcd(utils.Abs(difX), utils.Abs(difY))
	}
	difXX := difX / gcd
	difYY := difY / gcd

	for {
		x += difXX
		y += difYY
		if x == x2 && y == y2 {
			return false
		}
		if input[x][y] == '#' {
			return true
		}
	}
}

func part2() {
}
