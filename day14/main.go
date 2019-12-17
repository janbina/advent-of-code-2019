package main

import (
	"aoc19/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	part1()
	part2()
}

type ingredient struct {
	amount int
	name   string
}

type reaction struct {
	input  []ingredient
	output ingredient
}

func getInput() map[string]reaction {
	lines := utils.ReadLines("input.txt")
	reactions := make(map[string]reaction)

	r := strings.NewReplacer(" => ", " ", ", ", " ")
	for i := range lines {
		tokens := strings.Split(r.Replace(lines[i]), " ")
		ingredients := make([]ingredient, len(tokens)/2)

		for j := range ingredients {
			ingredients[j] = ingredient{utils.ToInt(tokens[j*2]), tokens[j*2+1]}
		}

		lastIndex := len(ingredients) - 1
		reactions[ingredients[lastIndex].name] = reaction{ingredients[:lastIndex], ingredients[lastIndex]}
	}
	return reactions
}

func part1() {
	reactions := getInput()

	fmt.Println(getOre(map[string]int{"FUEL": 1}, reactions))
}

func part2() {
	reactions := getInput()

	trillion := 1000000000000

	firstOver := sort.Search(trillion, func(n int) bool {
		return getOre(map[string]int{"FUEL": n}, reactions) > trillion
	})

	fmt.Println(firstOver - 1)
}

func getOre(needed map[string]int, reactions map[string]reaction) int {
	for {
		flag := true
		for k, v := range needed {
			if k != "ORE" && v > 0 {
				flag = false

				r := reactions[k]
				number := int(math.Ceil(float64(v) / float64(r.output.amount)))
				needed[k] -= number * r.output.amount
				for _, i := range r.input {
					needed[i.name] += number * i.amount
				}
			}
		}
		if flag {
			return needed["ORE"]
		}
	}
}
