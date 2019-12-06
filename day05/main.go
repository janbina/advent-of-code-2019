package main

import (
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
	ints := getInput()

	output := runProgram(ints, []int{1})
	fmt.Println(output)
}

func part2() {
	ints := getInput()

	output := runProgram(ints, []int{5})
	fmt.Println(output)
}

func getOpCodeAndModes(i int) (int, [3]int) {
	return i % 100, [3]int{(i / 100) % 10, (i / 1000) % 10, (i / 10000) % 10}
}

func getVal(mem []int, modes [3]int, ip int, offset int) int {
	if modes[offset] == 1 {
		return mem[ip+offset+1]
	} else {
		return mem[mem[ip+offset+1]]
	}
}

func runProgram(mem []int, input []int) []int {
	ip := 0
	inCnt := 0
	var output []int

	for {
		opCode, modes := getOpCodeAndModes(mem[ip])

		switch opCode {
		case 1:
			mem[mem[ip+3]] = getVal(mem, modes, ip, 0) + getVal(mem, modes, ip, 1)
			ip += 4
		case 2:
			mem[mem[ip+3]] = getVal(mem, modes, ip, 0) * getVal(mem, modes, ip, 1)
			ip += 4
		case 3:
			mem[mem[ip+1]] = input[inCnt]
			inCnt++
			ip += 2
		case 4:
			output = append(output, getVal(mem, modes, ip, 0))
			ip += 2
		case 5:
			if getVal(mem, modes, ip, 0) != 0 {
				ip = getVal(mem, modes, ip, 1)
			} else {
				ip += 3
			}
		case 6:
			if getVal(mem, modes, ip, 0) == 0 {
				ip = getVal(mem, modes, ip, 1)
			} else {
				ip += 3
			}
		case 7:
			if getVal(mem, modes, ip, 0) < getVal(mem, modes, ip, 1) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
			ip += 4
		case 8:
			if getVal(mem, modes, ip, 0) == getVal(mem, modes, ip, 1) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
			ip += 4
		case 99:
			return output
		default:
			panic("Invalid command")
		}
	}
}
