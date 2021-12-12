package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	lines, err := readFile("./sample.dat")
	if err != nil {
		panic(err)
	}
	values = make([][]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "|")
		values[i] = strings.Split(parts[1], " ")
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
	values [][]string
	lookup map[string]string = map[string]string{
		"cagedb": "0",
		//"ab":      "1",
		"cdgba": "2",
		//"cedba": "3",
		//"eafb":    "4",
		//"dcbef":  "5",
		"bcgafe": "6",
		//"dab":     "7",
		//"acedgfb": "8",
		"cefbgd": "9",
	}
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortedChars(s string) string {
	charArray := sortRunes([]rune(s))
	sort.Sort(charArray)
	return string(charArray)
}

func calculateAnswer() {
	count := 0
	sortedLookup := make(map[string]string, 10)
	for key, value := range lookup {
		sortedLookup[sortedChars(key)] = value
	}
	for _, line := range values {
		representation := ""
		for _, chunk := range line {
			length := len(chunk)
			if length >= 1 {
				switch length {
				case 2:
					representation = fmt.Sprintf("%s%s", representation, "1")
				case 3:
					representation = fmt.Sprintf("%s%s", representation, "7")
				case 4:
					representation = fmt.Sprintf("%s%s", representation, "4")
				case 7:
					representation = fmt.Sprintf("%s%s", representation, "8")
				default:
					// Do Lookup
					r := sortedLookup[sortedChars(chunk)]
					if r == "" {
						r = "#"
					}
					representation = fmt.Sprintf("%s%s", representation, r)
				}
			}
		}
		fmt.Printf("%s\n", representation)
	}
	fmt.Printf("%d\n", count)
}

func main() {
	readInputData()
	calculateAnswer()
}
