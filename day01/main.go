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
	modules := utils.ReadNumbers("input.txt")

	totalFuel := 0
	for _, module := range modules {
		totalFuel += fuelForModuleP1(module)
	}

	fmt.Println("Total fuel:", totalFuel)
}

func part2() {
	modules := utils.ReadNumbers("input.txt")

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