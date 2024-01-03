package main

import (
	"advent2023/day1"
	"advent2023/day10"
	"advent2023/day11"
	"advent2023/day12"
	"advent2023/day13"
	"advent2023/day14"
	"advent2023/day15"
	"advent2023/day16"
	"advent2023/day17"
	"advent2023/day18"
	"advent2023/day19"
	"advent2023/day2"
	"advent2023/day20"
	"advent2023/day21"
	"advent2023/day22"
	"advent2023/day23"
	"advent2023/day24"
	"advent2023/day25"
	"advent2023/day3"
	"advent2023/day4"
	"advent2023/day5"
	"advent2023/day6"
	"advent2023/day7"
	"advent2023/day8"
	"advent2023/day9"
	"fmt"
)

func main() {
	funcs := []func(string) (interface{}, interface{}){
		day1.Solve,
		day2.Solve,
		day3.Solve,
		day4.Solve,
		day5.Solve,
		day6.Solve,
		day7.Solve,
		day8.Solve,
		day9.Solve,
		day10.Solve,
		day11.Solve,
		day12.Solve,
		day13.Solve,
		day14.Solve,
		day15.Solve,
		day16.Solve,
		day17.Solve,
		day18.Solve,
		day19.Solve,
		day20.Solve,
		day21.Solve,
		day22.Solve,
		day23.Solve,
		day24.Solve,
		day25.Solve,
	}
	for i := 1; i <= len(funcs); i++ {

		// Slow
		if i == 5 || i == 10 || i == 12 || i == 16 || i == 22 {
			continue
		}

		arg := fmt.Sprintf("day%d/input.txt", i)
		p1, p2 := funcs[i-1](arg)
		fmt.Printf("Problem %d:\n\tPart 1: %v\n\tPart 2: %v\n\n", i, p1, p2)
	}
	// day17.Solve17("day17/input.txt")
	// day18.Solve("day18/input.txt")
	// day19.Solve("day19/input.txt")
	// day20.Solve("day20/input.txt")
	// day21.Solve("day21/input.txt")
	// day22.Solve("day22/input.txt")
	// day23.Solve("day23/input.txt")
	// day24.Solve("day24/input.txt")
	// day25.Solve("day25/input.txt")
}
