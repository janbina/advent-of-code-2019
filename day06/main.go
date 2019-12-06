package main

import (
	"aoc19/utils"
	"fmt"
	"strings"
)

func main() {
	part12()
}

func getInput() []string {
	return utils.ReadLines("input.txt")
}

func part12() {
	lines := getInput()

	edges := make(map[string][]string)
	for _, l := range lines {
		split := strings.Split(l, ")")
		a, b := split[0], split[1]
		edges[a] = append(edges[a], b)
	}

	acc := 0
	traverse(edges, "COM", &acc)
	fmt.Println(acc)
}

func traverse(edges map[string][]string, node string, acc *int) (int, int, int) {
	cnt := 0
	san, you := -1, -1
	for _, l := range edges[node] {
		num, s, y := traverse(edges, l, acc)
		cnt += num + 1
		if s >= 0 {
			san = s + 1
		}
		if y >= 0 {
			you = y + 1
		}
	}

	if san > 0 && you > 0 {
		fmt.Println(san + you - 2)
		san = -1
		you = -1
	}

	if node == "SAN" {
		san = 0
	} else if node == "YOU" {
		you = 0
	}

	*acc += cnt
	return cnt, san, you
}
