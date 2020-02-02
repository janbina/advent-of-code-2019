package main

import (
	"fmt"
	"github.com/janbina/advent-of-code-2019/utils"
)

func main() {
	part1()
	part2()
}

func getInput() []int {
	lines := utils.ReadLines("input.txt")
	return utils.StringsToInts(lines)
}

func part1() {
	modules := getInput()

	totalFuel := 0
	for _, module := range modules {
		totalFuel += fuelForModuleP1(module)
	}

	fmt.Println("Total fuel:", totalFuel)
}

func part2() {
	modules := getInput()

	totalFuel := 0
	for _, module := range modules {
		totalFuel += fuelForModuleP2(module)
	}

	fmt.Println("Total fuel:", totalFuel)
}

func fuelForModuleP1(module int) int {
	return fuelForMass(module)
}

func fuelForModuleP2(module int) int {
	totalFuel := 0
	fuel := fuelForMass(module)
	for fuel > 0 {
		totalFuel += fuel
		fuel = fuelForMass(fuel)
	}
	return totalFuel
}

func fuelForMass(mass int) int {
	return mass/3 - 2
}
