package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type LineParser func(line string) (interface{}, error)

func readInputData(parser LineParser, path string) ([]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item, err := parser(scanner.Text())
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, scanner.Err()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Point struct {
	x int
	y int
}

type LineSegment struct {
	one Point
	two Point
}

func (s *LineSegment) GetPoints() []Point {
	var points []Point = make([]Point, 0)
	displacementX := s.two.x - s.one.x
	displacementY := s.two.y - s.one.y
	directionX := int(math.Copysign(1, float64(displacementX)))
	directionY := int(math.Copysign(1, float64(displacementY)))
	displacementXAbs := int(math.Abs(float64(displacementX)))
	displacementYAbs := int(math.Abs(float64(displacementY)))

	if s.one.x == s.two.x { // Horizontal
		for i := 0; i <= displacementYAbs; i++ {
			points = append(points, Point{
				x: s.one.x,
				y: s.one.y + i*directionY,
			})
		}
	}
	if s.one.y == s.two.y { // Vertical
		for i := 0; i <= displacementXAbs; i++ {
			points = append(points, Point{
				x: s.one.x + i*directionX,
				y: s.one.y,
			})
		}
	}
	if displacementXAbs == displacementYAbs { // 45 Diagonal
		for i := 0; i <= displacementXAbs; i++ {
			points = append(points, Point{
				x: s.one.x + i*directionX,
				y: s.one.y + i*directionY,
			})
		}
	}
	return points
}

var expression = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func getOrPanicInt(data string) int {
	value, err := strconv.Atoi(data)
	checkError(err)
	return value
}

func parser(line string) (interface{}, error) {
	data := expression.FindAllStringSubmatch(line, -1)
	if len(data) == 0 {
		return nil, fmt.Errorf("Failed Regex Parse")
	}
	if len(data[0]) != 5 {
		return nil, fmt.Errorf("Failed Regex Parse")
	}
	return LineSegment{
		one: Point{
			x: getOrPanicInt(data[0][1]),
			y: getOrPanicInt(data[0][2]),
		},
		two: Point{
			x: getOrPanicInt(data[0][3]),
			y: getOrPanicInt(data[0][4]),
		},
	}, nil
}

func calculateAnswer(data []interface{}) {
	store := map[Point]int{}
	for _, item := range data {
		segment := item.(LineSegment)
		for _, point := range segment.GetPoints() {
			if _, ok := store[point]; ok {
				store[point] += 1
			} else {
				store[point] = 1
			}
		}
	}
	count := 0
	for _, value := range store {
		if value >= 2 {
			count += 1
		}
	}
	fmt.Printf("%d\n", count)
}

func main() {
	data, err := readInputData(parser, "./input.dat")
	checkError(err)
	calculateAnswer(data)
}
