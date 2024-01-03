package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readInput(filename string) []string {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sum []string
	for fileScanner.Scan() {
		sum = append(sum, fileScanner.Text())
	}
	readFile.Close()
	return sum
}

type Point struct {
	x int
	y int
}

func isStar(ch rune) bool {
	return ch == '*'
}

func isSymbol(ch rune) bool {
	return ch != '.' && !unicode.IsDigit(ch)
}

func hasStarNeighbor(lines *[]string, lineIndex int, start int, end int) *Point {
	begin := 0
	if start > 1 {
		begin = start - 1
	}

	final := len((*lines)[lineIndex]) - 1
	if end+1 < final {
		final = end + 1
	}

	ch1 := rune((*lines)[lineIndex][begin])
	ch2 := rune((*lines)[lineIndex][final])
	if isStar(ch1) {
		return &Point{
			begin,
			lineIndex,
		}
	}
	if isStar(ch2) {
		return &Point{
			final,
			lineIndex,
		}
	}

	// check line above
	if lineIndex-1 > 0 {
		for i := begin; i <= final; i++ {
			ch := rune((*lines)[lineIndex-1][i])
			if isStar(ch) {
				// symbol
				return &Point{i, lineIndex - 1}
			}
		}
	}

	// check line below
	if lineIndex+1 < len(*lines)-1 {
		for i := begin; i <= final; i++ {
			ch := rune((*lines)[lineIndex+1][i])
			if isStar(ch) {
				// symbol
				return &Point{i, lineIndex + 1}
			}
		}
	}
	return nil
}

func hasSymbolNeighbor(lines *[]string, lineIndex int, start int, end int) bool {
	begin := 0
	if start > 1 {
		begin = start - 1
	}

	final := len((*lines)[lineIndex]) - 1
	if end+1 < final {
		final = end + 1
	}

	ch1 := rune((*lines)[lineIndex][begin])
	ch2 := rune((*lines)[lineIndex][final])
	if isSymbol(ch1) || isSymbol(ch2) {
		return true
	}

	// check line above
	if lineIndex-1 > 0 {
		for i := begin; i <= final; i++ {
			ch := rune((*lines)[lineIndex-1][i])
			if isSymbol(ch) {
				// symbol
				return true
			}
		}
	}

	// check line below
	if lineIndex+1 < len(*lines)-1 {
		for i := begin; i <= final; i++ {
			ch := rune((*lines)[lineIndex+1][i])
			if isSymbol(ch) {
				// symbol
				return true
			}
		}
	}
	return false
}
func part1(lines []string) int {
	sum := 0

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for i := 0; i < len(line); i++ {
			ch := rune(line[i])
			if unicode.IsDigit(ch) {
				start := i
				for {
					i++
					if i == len(line) || !unicode.IsDigit(rune(line[i])) {
						i--
						break
					}
				}
				num, _ := strconv.Atoi(line[start:(i + 1)])
				if hasSymbolNeighbor(&lines, y, start, i) {
					sum += num
				}
			}
		}
	}

	return sum
}
func part2(lines []string) int {
	points := make(map[string][]int)

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for i := 0; i < len(line); i++ {
			ch := rune(line[i])
			if unicode.IsDigit(ch) {
				start := i
				for {
					i++
					if i == len(line) || !unicode.IsDigit(rune(line[i])) {
						i--
						break
					}
				}
				num, _ := strconv.Atoi(line[start:(i + 1)])
				p := hasStarNeighbor(&lines, y, start, i)
				if p != nil {
					key := fmt.Sprintf("%d,%d", p.x, p.y)
					points[key] = append(points[key], num)
				}
			}
		}
	}
	sum := 0
	for _, value := range points {
		if len(value) == 2 {
			sum += value[0] * value[1]
		}
	}

	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	lines := readInput(filename)
	p1 := part1(lines)
	p2 := part2(lines)
	return p1, p2
}
