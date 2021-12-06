package main

import (
	"fmt"
	"strconv"

	"github.com/spoohie/AoC-2021-Go/utils"
)

func main() {
	raw_data := utils.ParseFile()
	var data []int
	for _, d := range raw_data {
		val, _ := strconv.Atoi(d)
		data = append(data, val)
	}
	fmt.Printf("Part one solution: %d\n", part_one(data))
	fmt.Printf("Part two solution: %d\n", part_two(data))
}

func part_one(data []int) int {
	val := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			val += 1
		}
	}
	return val
}

func part_two(data []int) int {
	val := 0
	for i := 3; i < len(data); i++ {
		a := data[i-3] + data[i-2] + data[i-1]
		b := data[i-2] + data[i-1] + data[i]
		if a < b {
			val += 1
		}
	}
	return val
}
