package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spoohie/AoC-2021-Go/utils"
)

func main() {
	raw_data := strings.Split(utils.ParseFile()[0], ",")
	var data []int
	for _, d := range raw_data {
		val, _ := strconv.Atoi(d)
		data = append(data, val)
	}
	fmt.Printf("Part one solution: %d\n", part_one(data))
	fmt.Printf("Part two solution: %d\n", part_two(data))
}

func part_one(data []int) int {
	return calculate_growth(data, 80)
}

func part_two(data []int) int {
	return calculate_growth(data, 256)
}

func calculate_growth(data []int, days int) int {
	var fish_by_age [9]int
	for _, f := range data {
		fish_by_age[f] += 1
	}
	for day := 0; day < days; day++ {
		zeroes := 0
		for i, f := range fish_by_age {
			if i == 0 {
				if f > 0 {
					zeroes = f
				}
			} else {
				fish_by_age[i-1] += f
			}
			fish_by_age[i] = 0
		}
		fish_by_age[6] += zeroes
		fish_by_age[8] += zeroes
	}
	return utils.Sum(fish_by_age[:])
}
