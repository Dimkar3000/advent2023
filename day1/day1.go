package day1

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parseLine(line string) int {
	var digits []int
	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, (int(c) - '0'))
		}
	}
	if len(digits) == 0 {
		return 0
	}
	if len(digits) == 1 {
		return digits[0]*10 + digits[0]
	}
	d1 := digits[0]
	d2 := digits[len(digits)-1]
	return d1*10 + d2
}

var wordDigits = [10]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseLine2(line string) int {
	var digits []int
	for i := 0; i < len(line); i++ {
		// fmt.Printf("i: %d ", i)
		if unicode.IsDigit(rune(line[i])) {
			digits = append(digits, int(line[i]-'0'))
		}
		for d := 0; d < 10; d++ {
			// fmt.Printf("d: %d ", d)
			if line[i:i+min(len(wordDigits[d]), len(line)-i)] == wordDigits[d] {
				digits = append(digits, d)
				break
			}
		}
	}
	if len(digits) == 0 {
		return 0
	}
	if len(digits) == 1 {
		return digits[0]*10 + digits[0]
	}
	d1 := digits[0]
	d2 := digits[len(digits)-1]
	return d1*10 + d2
}

func part1(filename string) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += parseLine(line)
	}
	readFile.Close()
	return sum
}

func part2(filename string) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += parseLine2(line)
	}
	readFile.Close()
	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	p1 := part1(filename)
	p2 := part2(filename)
	return p1, p2
}
