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
		for _, value := range items {
			timer := getOrPanicInt(value)
			fish[timer] += 1
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
	// Maps a Fish Timer -> Number of Fish
	fish = make([]int, 9)
)

func simulate(days int) {
	for i := 0; i < days; i++ {
		fmt.Printf("Simulate Day %d\n", i)
		newFish := make([]int, 9)
		for timer, fishCount := range fish {
			if timer == 0 {
				newFish[growthRate] += fishCount
				newFish[newGrowthRate] += fishCount
			} else {
				newFish[timer-1] += fishCount
			}
		}
		fish = newFish
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	readInputData()
	simulate(256)
	fmt.Printf("%d", sum(fish))
}
