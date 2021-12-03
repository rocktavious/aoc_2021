package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	horizontal int
	vertical   int
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
	exp := regexp.MustCompile(`^(forward|down|up) (\d+)$`)
	for _, line := range readInputData() {
		data := exp.FindAllStringSubmatch(line, -1)
		if len(data) == 0 {
			continue
		}
		if len(data[0]) != 3 {
			continue
		}
		amount, err := strconv.Atoi(data[0][2])
		if err != nil {
			panic(err)
		}
		switch data[0][1] {
		case "forward":
			horizontal += amount
		case "down":
			vertical += amount
		case "up":
			vertical -= amount
		}
	}
}

func calculateAnswer() {
	fmt.Printf("%d\n", horizontal*vertical)
}

func main() {
	parseData()
	calculateAnswer()
}
