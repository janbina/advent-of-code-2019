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
	x, y, z int
}

func (p point) absSum() int {
	return utils.Abs(p.x) + utils.Abs(p.y) + utils.Abs(p.z)
}

type moon struct {
	position *point
	velocity *point
}

func getInput() []moon {
	return []moon{
		moon{&point{-7, -8, 9}, &point{0, 0, 0}},
		moon{&point{-12, -3, -4}, &point{0, 0, 0}},
		moon{&point{6, -17, -9}, &point{0, 0, 0}},
		moon{&point{4, -10, -6}, &point{0, 0, 0}}}
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func part1() {
	moons := getInput()

	for step := 0; step < 1000; step++ {
		for i, m1 := range moons {
			for j, m2 := range moons {
				if i == j {
					continue
				}
				m1.velocity.x += sign(m2.position.x - m1.position.x)
				m1.velocity.y += sign(m2.position.y - m1.position.y)
				m1.velocity.z += sign(m2.position.z - m1.position.z)
			}
		}
		for _, m := range moons {
			m.position.x += m.velocity.x
			m.position.y += m.velocity.y
			m.position.z += m.velocity.z
		}
	}

	acc := 0
	for _, m := range moons {
		acc += m.position.absSum() * m.velocity.absSum()
	}

	fmt.Println(acc)
}

func part2() {
	moons := getInput()
	init := getInput()

	var xCycleLen, yCycleLen, zCycleLen int

	for step := 0; ; step++ {
		for i, m1 := range moons {
			for j, m2 := range moons {
				if i == j {
					continue
				}
				m1.velocity.x += sign(m2.position.x - m1.position.x)
				m1.velocity.y += sign(m2.position.y - m1.position.y)
				m1.velocity.z += sign(m2.position.z - m1.position.z)
			}
		}
		for _, m := range moons {
			m.position.x += m.velocity.x
			m.position.y += m.velocity.y
			m.position.z += m.velocity.z
		}
		xF, yF, zF := true, true, true
		for i, m := range moons {
			xF = xF && m.position.x == init[i].position.x && m.velocity.x == init[i].velocity.x
			yF = yF && m.position.y == init[i].position.y && m.velocity.y == init[i].velocity.y
			zF = zF && m.position.z == init[i].position.z && m.velocity.z == init[i].velocity.z
		}
		if xCycleLen == 0 && xF {
			xCycleLen = step + 1
		}
		if yCycleLen == 0 && yF {
			yCycleLen = step + 1
		}
		if zCycleLen == 0 && zF {
			zCycleLen = step + 1
		}
		if xCycleLen > 0 && yCycleLen > 0 && zCycleLen > 0 {
			break
		}
	}

	res := utils.Lcm(utils.Lcm(xCycleLen, yCycleLen), zCycleLen)
	fmt.Println(res)
}
