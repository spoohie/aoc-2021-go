package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		data = append(data, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
