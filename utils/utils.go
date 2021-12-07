package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFile() (data []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func Sum(array []int) (s int) {
	for _, v := range array {
		s += v
	}
	return s
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
