package main

import (
	"fmt"
	"github.com/janbina/advent-of-code-2019/utils"
	"math"
	"sort"
)

func main() {
	best, max := part1()
	fmt.Println(max)
	a200th := part2(best)
	fmt.Println(a200th.y*100 + a200th.x)
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

func part1() (point, int) {
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

	return best, max
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

type helper struct {
	angle    float64
	len      int
	asteroid point
	order    int
}

func part2(best point) point {
	asteroids := getInput()
	delete(asteroids, best)
	h := make([]helper, 0, len(asteroids))

	for k := range asteroids {
		vec := point{k.x - best.x, k.y - best.y}
		len := utils.Abs(vec.x) + utils.Abs(vec.y)
		angle := getAngle(float64(vec.x), float64(vec.y))
		h = append(h, helper{angle, len, k, 0})
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].angle < h[j].angle
	})

	cnt := 0
	prev := -1000.0
	for i, x := range h {
		if x.angle == prev {
			cnt++
			h[i].order = cnt
		} else {
			cnt = 0
			h[i].order = 0
			prev = x.angle
		}
	}

	sort.Slice(h, func(i, j int) bool {
		if h[i].order == h[j].order {
			return h[i].angle < h[j].angle
		}
		return h[i].order < h[j].order
	})

	return h[199].asteroid
}

func getAngle(x, y float64) float64 {
	a := math.Atan2(x, y)
	if a >= 0 {
		return a
	}
	a += 2 * math.Pi

	if a >= 3*math.Pi/2 {
		a = a - 2*math.Pi
	}
	return a
}
