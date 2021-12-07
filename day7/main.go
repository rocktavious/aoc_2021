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
	positions      []int
	targetPosition int
)

func calculateMedian(input ...int) int {
	sort.Ints(input)

	count := len(input)
	target := count / 2

	if count%2 != 0 {
		return input[target]
	}

	return (input[target-1] + input[target]) / 2
}

func calculateAnswer() {
	total := 0
	for _, position := range positions {
		total += int(math.Abs(float64(position) - float64(targetPosition)))
	}

	fmt.Printf("%d\n", total)
}

func main() {
	readInputData()

	targetPosition = calculateMedian(positions...)
	calculateAnswer()
}
