package day12

// Based on:
// https://github.com/Anshuman-UCSB/Advent-Of-Code/blob/master/2023/Python/src/day12.py

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	pattern    []rune
	validators []int64
	problems   int64
}

type Puzzle struct {
	problems []Problem
}

func readInput(file string, isPart2 bool) Puzzle {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")

	problems := []Problem{}
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 2)
		re, _ := regexp.Compile(`\.+`)
		pattern := string(re.ReplaceAll([]byte(parts[0]), []byte(".")))
		validatorsS := strings.Split(parts[1], ",")
		validators := []int64{}
		for _, v := range validatorsS {
			value, _ := strconv.Atoi(v)
			validators = append(validators, int64(value))
		}

		p := []rune(pattern)
		v := validators
		if isPart2 {

			p = append(p, '?')
			p = append(p, []rune(pattern)...)
			p = append(p, '?')
			p = append(p, []rune(pattern)...)
			p = append(p, '?')
			p = append(p, []rune(pattern)...)
			p = append(p, '?')
			p = append(p, []rune(pattern)...)
			v = append(v, validators...)
			v = append(v, validators...)
			v = append(v, validators...)
			v = append(v, validators...)
		}
		problems = append(problems, Problem{
			pattern:    p,
			validators: v,
			problems:   int64(strings.Count(string(p), "?")),
		})
	}

	return Puzzle{problems: problems}
}

var dp = map[string]int64{}

func s(line []rune, nums []int64, i, n, b int64) int64 {
	key := fmt.Sprintf("%v,%v,%v", i, n, b)
	v, ok := dp[key]
	if ok {
		return v
	}

	if i == int64(len(line)) {
		if (n == int64(len(nums)) && b == 0) ||
			(n == int64(len(nums)-1) && b == nums[len(nums)-1]) {
			return 1
		} else {
			return 0
		}
	}

	result := int64(0)

	if line[i] == '.' || line[i] == '?' {
		if b == 0 {
			result += s(line, nums, i+1, n, 0)
		} else {
			if n == int64(len(nums)) {
				return 0
			}
			if b == nums[n] {
				result += s(line, nums, i+1, n+1, 0)
			}
		}
	}

	if line[i] == '#' || line[i] == '?' {
		result += s(line, nums, i+1, n, b+1)
	}

	dp[key] = result
	return result
}

func solve(line []rune, nums []int64) int64 {
	dp = map[string]int64{}
	return s(line, nums, 0, 0, 0)
}

func part1(puzzle Puzzle) int64 {
	sum := int64(0)
	for _, problem := range puzzle.problems {
		v := solve(problem.pattern, problem.validators)
		sum += v
	}
	return sum
}

func part2(puzzle Puzzle) int64 {
	sum := int64(0)
	for _, problem := range puzzle.problems {
		v := solve(problem.pattern, problem.validators)
		sum += v
	}
	return sum
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle1 := readInput(filename, false)
	puzzle2 := readInput(filename, true)
	return part1(puzzle1), part2(puzzle2)
}
