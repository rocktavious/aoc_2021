package main

import (
	"bufio"
	"fmt"
	"os"
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
		parts := strings.Split(line, "|")
		chunks := strings.Split(parts[1], " ")
		for _, chunk := range chunks {
			if len(chunk) >= 1 {
				values = append(values, chunk)
			}
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
	values []string
)

func calculateAnswer() {
	count := 0
	for _, chunk := range values {
		switch len(chunk) {
		case 2:
			count += 1
		case 3:
			count += 1
		case 4:
			count += 1
		case 7:
			count += 1
		}
	}
	fmt.Printf("%d\n", count)
}

func main() {
	readInputData()
	calculateAnswer()
}
