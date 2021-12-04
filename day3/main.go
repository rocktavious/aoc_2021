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

func runeCommonality(lines []string, i int) (uint8, uint8) {
	count0, count1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if count0 > count1 {
		return '0', '1'
	}
	return '1', '0'
}

func filter(lines []string, i int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	most, least := runeCommonality(lines, i)
	comparator := least
	if mostCommon {
		comparator = most
	}
	filtered := make([]string, 0)
	for _, l := range lines {
		if l[i] == comparator {
			filtered = append(filtered, l)
		}
	}
	return filter(filtered, i+1, mostCommon)
}

func fromBinary(s string) (r int) {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func main() {
	lines := readInputData()
	oxygen := filter(lines, 0, true)
	scrubber := filter(lines, 0, false)
	fmt.Println(fromBinary(oxygen) * fromBinary(scrubber))
}
