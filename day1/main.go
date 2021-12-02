package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	count := 0
	store := 0
	lines := readInputData()
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
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
