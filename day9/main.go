package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	heightmap = make([][]int, len(lines))
	for i, line := range lines {
		parts := []rune(line)
		heightmap[i] = make([]int, len(parts))
		for j, part := range parts {
			heightmap[i][j] = getOrPanicInt(string(part))
		}
	}
	xMax = len(heightmap[0])
	yMax = len(heightmap)
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
	heightmap [][]int
	xMax      int
	yMax      int
	north     Point = Point{
		x: 0,
		y: -1,
	}
	south Point = Point{
		x: 0,
		y: 1,
	}
	east Point = Point{
		x: 1,
		y: 0,
	}
	west Point = Point{
		x: -1,
		y: 0,
	}
)

type Point struct {
	x int
	y int
}

func addPoints(p1 Point, p2 Point) Point {
	return Point{
		x: p1.x + p2.x,
		y: p1.y + p2.y,
	}
}

func isWithinMap(point Point) bool {
	return point.x >= 0 && point.x < xMax && point.y >= 0 && point.y < yMax
}

func getHeight(point Point) int {
	if isWithinMap(point) {
		return heightmap[point.y][point.x]
	} else {
		return 9
	}
}

func isAllCardinalWithinMapHigher(point Point, height int) bool {
	// North -y
	if point.y-1 >= 0 {
		if height >= getHeight(addPoints(point, north)) {
			return false
		}
	}
	// South +y
	if point.y+1 < yMax {
		if height >= getHeight(addPoints(point, south)) {
			return false
		}
	}
	// East +x
	if point.x+1 < xMax {
		if height >= getHeight(addPoints(point, east)) {
			return false
		}
	}
	// West -x
	if point.x-1 >= 0 {
		if height >= getHeight(addPoints(point, west)) {
			return false
		}
	}
	return true
}

func calculateAnswer() {
	lowPoints := make([]int, 0)

	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			point := Point{
				x: x,
				y: y,
			}
			height := heightmap[point.y][point.x]
			if isAllCardinalWithinMapHigher(point, height) == false {
				continue
			}
			fmt.Printf("%d,%d\n", x, y)
			lowPoints = append(lowPoints, height)
		}
	}
	//fmt.Printf("%v\n", lowPoints)
	sum := 0
	for _, point := range lowPoints {
		sum += point + 1
	}
	fmt.Printf("%d\n", sum)
}

func main() {
	readInputData()
	calculateAnswer()
}
