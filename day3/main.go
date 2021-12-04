package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	gamma   int
	epsilon int
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readInputData() []string {
	lines, err := readFile("./input.dat")
	if err != nil {
		panic(err)
	}
	return lines
}

func parseData() {
	data := make(map[int]int, 0)
	for _, line := range readInputData() {
		for i, bit := range line {
			index := len(line) - 1 - i
			switch bit {
			case rune('0'):
				data[index] -= 1
			case rune('1'):
				data[index] += 1
			}
		}
	}
	var result uint = 0
	var inverse uint = 0
	for i, count := range data {
		if count == 0 {
			panic("is this possible?")
		}
		if count > 0 {
			result |= 1 << (i)
		} else {
			result |= 0 << (i)
		}
		inverse |= 1 << (i)
	}
	gamma = int(result)
	epsilon = int(result ^ inverse)
}

func calculateAnswer() {
	fmt.Println(gamma * epsilon)
}

func main() {
	parseData()
	calculateAnswer()
}
