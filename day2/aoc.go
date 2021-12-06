package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spoohie/AoC-2021-Go/utils"
)

type Move struct {
	dir string
	val int
}

func main() {
	raw_data := utils.ParseFile()
	var data []Move
	for _, v := range raw_data {
		line := strings.Split(v, " ")
		dir := line[0]
		val, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Move{dir, val})
	}

	fmt.Printf("Part one solution: %d\n", part_one(data))
	fmt.Printf("Part two solution: %d\n", part_two(data))
}

func part_one(moves []Move) int {
	hor := 0
	depth := 0
	for _, m := range moves {
		if m.dir == "forward" {
			hor += m.val
		} else if m.dir == "down" {
			depth += m.val
		} else if m.dir == "up" {
			depth -= m.val
		}
	}
	return hor * depth
}

func part_two(moves []Move) int {
	hor := 0
	depth := 0
	aim := 0

	for _, m := range moves {
		if m.dir == "forward" {
			hor += m.val
			depth = depth + aim*m.val
		} else if m.dir == "down" {
			aim += m.val
		} else if m.dir == "up" {
			aim -= m.val
		}
	}
	return hor * depth
}
