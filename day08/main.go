package main

import (
	"aoc19/utils"
	"fmt"
)

const width = 25
const height = 6
const layerLen = width * height

func main() {
	part1()
	part2()
}

func getInput() [][]byte {
	line := utils.ReadLines("input.txt")[0]
	numLayers := len(line) / layerLen
	layers := make([][]byte, numLayers)
	for i := 0; i < numLayers; i++ {
		layer := make([]byte, layerLen)
		for j := 0; j < layerLen; j++ {
			layer[j] = line[i*layerLen+j] - '0'
		}
		layers[i] = layer
	}
	return layers
}

func part1() {
	layers := getInput()

	minZeros := layerLen
	checkSum := 0

	for _, layer := range layers {
		var counts [10]int
		for _, pixel := range layer {
			counts[pixel]++
		}
		if counts[0] < minZeros {
			minZeros = counts[0]
			checkSum = counts[1] * counts[2]
		}
	}

	fmt.Println(checkSum)
}

func part2() {
	layers := getInput()

	var image [layerLen]byte
	for i := range image {
		image[i] = 2
	}

	for _, layer := range layers {
		for j, pixel := range layer {
			if image[j] == 2 {
				image[j] = pixel
			}
		}
	}

	for i, p := range image {
		if i%width == 0 {
			fmt.Println()
		}
		if p == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
