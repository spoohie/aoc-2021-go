package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part one solution: %d\n", part_one(data))
	fmt.Printf("Part two solution: %d\n", part_two(data))
}

func part_one(data []string) int64 {
	gamma, epsilon := calculate_bit_chains(data)
	return bintodec(gamma) * bintodec(epsilon)
}

func part_two(data []string) int64 {
	return oxygen_rating(data) * co2_rating(data)
}

func calculate_bit_chains(data []string) (string, string) {
	var mcb string
	var lcb string
	for i := 0; i < len(data[0]); i++ {
		var s string
		for _, v := range data {
			s += string([]rune(v)[i])
		}
		mcb += most_common(s)
		lcb += least_common(s)
	}
	return mcb, lcb
}

func most_common(s string) string {
	if strings.Count(s, "1") >= strings.Count(s, "0") {
		return "1"
	}
	return "0"
}

func least_common(s string) string {
	if strings.Count(s, "1") >= strings.Count(s, "0") {
		return "0"
	}
	return "1"
}

func bintodec(bitval string) int64 {
	val, err := strconv.ParseInt(bitval, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func oxygen_rating(data []string) int64 {
	for i := 0; i < len(data[0]); i++ {
		mcb_chain, _ := calculate_bit_chains(data)
		data = match_data_with_chain(data, mcb_chain, i)

		if len(data) == 1 {
			break
		}
	}
	return bintodec(data[0])
}

func co2_rating(data []string) int64 {
	for i := 0; i < len(data[0]); i++ {
		_, lcb_chain := calculate_bit_chains(data)
		data = match_data_with_chain(data, lcb_chain, i)
		if len(data) == 1 {
			break
		}
	}
	return bintodec(data[0])
}

func match_data_with_chain(data []string, chain string, index int) []string {
	var new_data []string
	for _, d := range data {
		if d[index] == chain[index] {
			new_data = append(new_data, d)
		}
	}
	return new_data
}
