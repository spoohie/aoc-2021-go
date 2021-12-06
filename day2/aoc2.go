package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	dir string
	val int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []Move
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		val, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Move{dir, val})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
