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

type module struct {
	mem []int
	ip  int
}

func (m *module) run(input []int) []int {
	nIP, out := common.RunIntcode(m.mem, m.ip, input)
	m.ip = nIP
	return out
}

func findMaxSignal(program []int, settings []int) int {
	maxSignal := 0

	utils.WithPermutation(settings, func(p []int) {
		var modules []*module
		signal := 0
		for i := range p {
			mem := utils.CopyInts(program)
			mod := module{mem, 0}
			modules = append(modules, &mod)
			signal = mod.run([]int{p[i], signal})[0]
		}
		for i := 0; true; i++ {
			out := modules[i%len(modules)].run([]int{signal})
			if len(out) == 0 {
				break
			}
			signal = out[0]
		}
		if signal > maxSignal {
			maxSignal = signal
		}
	})

	return maxSignal
}
