package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LineParser func(line string) (interface{}, error)

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

func readInputData() {
	lines, err := readFile("./input.dat")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		items := strings.Split(line, ",")
		for _, value := range items {
			position := getOrPanicInt(value)
			positions = append(positions, position)
		}
	}
}

func getOrPanicInt(data string) int {
	value, err := strconv.Atoi(data)
	checkError(err)
	return value
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

////////

var (
	positions []int
)

func calculateAnswer(targetPosition int) int {
	total := 0
	for _, position := range positions {
		steps := int(math.Abs(float64(position) - float64(targetPosition)))
		for i := 0; i < steps; i++ {
			total += i + 1
		}
	}
	return total
}

func main() {
	readInputData()
	sort.Ints(positions)
	min := positions[0]
	max := positions[len(positions)-1]
	results := make([]int, 0)
	for i := min; i <= max; i++ {
		results = append(results, calculateAnswer(i))
	}
	sort.Ints(results)
	fmt.Printf("%d", results[0])
}
