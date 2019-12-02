package main

import (
	"aoc19/utils"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	ints := utils.ReadNumbers2("input.txt")

	out := runProgram(ints, 12, 2)

	fmt.Println("Output =", out)
}

func part2() {
	ints := utils.ReadNumbers2("input.txt")

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			out := runProgram(utils.CopyInts(ints), noun, verb)
			if out == 19690720 {
				fmt.Println("Noun =", noun, "verb =", verb, "answer =", noun*100+verb)
				return
			}
		}
	}
}

func runProgram(mem []int, noun int, verb int) int {
	mem[1] = noun
	mem[2] = verb
	position := 0

	for {
		command := mem[position]
		if command == 99 {
			break
		} else if command == 1 {
			mem[mem[position+3]] = mem[mem[position+1]] + mem[mem[position+2]]
		} else if command == 2 {
			mem[mem[position+3]] = mem[mem[position+1]] * mem[mem[position+2]]
		}
		position += 4
	}

	return mem[0]
}
