package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFile() []string {
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
	return data
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func MaxIntSlice(s []int) (m int) {
	for i, e := range s {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
