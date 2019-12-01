package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	modules := readNumbers("input.txt")

	totalFuel := 0
	for _, module := range modules {
		totalFuel += fuelForModuleP1(module)
	}

	fmt.Println("Total fuel:", totalFuel)
}

func part2() {
	modules := readNumbers("input.txt")

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
	fuel := fuelForMass(module)
	added := fuelForMass(fuel)
	for added > 0 {
		fuel += added
		added = fuelForMass(added)
	}
	return fuel
}

func fuelForMass(mass int) int {
	return mass/3 - 2
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
