package main

import (
	"fmt"
	"math"
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
	return calculate_fuel(data, utils.Abs)
}

func part_two(data []int) int {
	return calculate_fuel(data, fuel_sum)
}

func calculate_fuel(data []int, f func(int) int) int {
	current_fuel := math.MaxInt32
	for i := 0; i < utils.MaxIntSlice(data); i++ {
		var new_fuel_elements []int
		for _, d := range data {
			new_fuel_elements = append(new_fuel_elements, f(d-i))
		}
		new_fuel := utils.Sum(new_fuel_elements)
		if new_fuel < current_fuel {
			current_fuel = new_fuel
		}
	}
	return current_fuel
}

func fuel_sum(diff int) int {
	n := utils.Abs(diff)
	return (2 + (n - 1)) * n / 2
}
