package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	Store map[int]int = make(map[int]int)
	Sums  map[int]int = make(map[int]int)
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

func tryGetStore(i int) int {
	if i <= len(Store) {
		return Store[i]
	}
	return 0
}

func parseData() {
	for i, line := range readInputData() {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		Store[i] = value
	}
}

func calculateData() {
	for i := 0; i < len(Store); i++ {
		if i+2 > len(Store) {
			return
		}
		Sums[i] = tryGetStore(i) + tryGetStore(i+1) + tryGetStore(i+2)
	}
}

func calculateAnswer() {
	count := 0
	store := 0
	for i := 0; i < len(Sums); i++ {
		value := Sums[i]
		if i == 0 {
			store = value
			continue
		}
		if value > store {
			count += 1
		}
		store = value
	}
	fmt.Println(count)
}

func main() {
	parseData()
	calculateData()
	calculateAnswer()
}
