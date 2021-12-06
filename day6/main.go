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
		items := strings.Split(line, ",")
		for _, timer := range items {
			value := getOrPanicInt(timer)
			fish = append(fish, value)
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

//////

var (
	growthRate    = 6
	newGrowthRate = 8
	fish          = make([]int, 0)
)

func simulate(days int) {
	//fmt.Printf("Simulate Day %d\n", days)
	newFish := 0
	for i, timer := range fish {
		// Tick
		fish[i] = timer - 1
		if timer == 0 {
			fish[i] = growthRate
			newFish += 1
		}
	}
	for i := 0; i < newFish; i++ {
		fish = append(fish, newGrowthRate)
	}

	if days > 0 {
		simulate(days - 1)
	}
}

func main() {
	readInputData()
	simulate(79)
	fmt.Printf("%d", len(fish))
}
