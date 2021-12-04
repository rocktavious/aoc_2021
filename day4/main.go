package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	moves  []int
	boards []*BingoBoard
)

type BingoBoardSlot struct {
	Value int
	Data  bool
}

func NewBingoBoardSlot() (slot BingoBoardSlot) {
	return
}

func NewBingoBoardRow() [5]BingoBoardSlot {
	return [5]BingoBoardSlot{
		NewBingoBoardSlot(),
		NewBingoBoardSlot(),
		NewBingoBoardSlot(),
		NewBingoBoardSlot(),
		NewBingoBoardSlot(),
	}
}

func NewBingoBoardData() [5][5]BingoBoardSlot {
	return [5][5]BingoBoardSlot{
		NewBingoBoardRow(),
		NewBingoBoardRow(),
		NewBingoBoardRow(),
		NewBingoBoardRow(),
		NewBingoBoardRow(),
	}
}

type BingoBoard struct {
	Data   [5][5]BingoBoardSlot
	Solved bool
}

func (b *BingoBoard) IsSolved(x int, y int) {
	// Check Horizontal
	for i := 0; i < 5; i++ {
		if b.Data[y][i].Data == false {
			break
		}
		if i == 4 {
			b.Solved = true
			return
		}
	}
	// Check Vertical
	for i := 0; i < 5; i++ {
		if b.Data[i][x].Data == false {
			break
		}
		if i == 4 {
			b.Solved = true
			return
		}
	}
}

func (b *BingoBoard) Set(value int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.Data[y][x].Value == value {
				b.Data[y][x].Data = true
				b.IsSolved(x, y)
			}
		}
	}
}

func NewBingoBoard(lines []string) *BingoBoard {
	board := &BingoBoard{
		Data: NewBingoBoardData(),
	}
	for j := 0; j < 5; j++ {
		for i, value := range strings.Fields(lines[j]) {
			board.Data[j][i].Value = stringToInt(value)
		}
	}
	return board
}

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

func stringToInt(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return int(value)
}

func loadData(lines []string) {
	for _, value := range strings.Split(lines[0], ",") {
		moves = append(moves, stringToInt(value))
	}

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, NewBingoBoard(lines[i:i+5]))
	}
}

func areAllBoardsSolved() bool {
	for _, board := range boards {
		if board.Solved == false {
			return false
		}
	}
	return true
}

func play() {
	for _, value := range moves {
		for _, board := range boards {
			board.Set(value)
			if areAllBoardsSolved() {
				calculateAnswer(board, value)
				return
			}
		}
	}
	panic(fmt.Errorf("Ran Out of Moves! No Boards Won"))
}

func calculateAnswer(board *BingoBoard, finalMove int) {
	var sum int = 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if board.Data[y][x].Data == false {
				sum += board.Data[y][x].Value
			}
		}
	}
	fmt.Printf("%v\n", sum*finalMove)
}

func main() {
	loadData(readInputData())
	play()
}
