package day14

import (
	"bytes"
	"os"
	"strings"
)

type Puzzle struct {
	grid [][]byte
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return Puzzle{grid: grid}
}

func calculateScore(grid [][]byte) int {
	sum := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 'O' {
				sum += len(grid) - row
			}
		}
	}

	return sum
}

func RollNorth(grid [][]byte) {
	for col := 0; col < len(grid[0]); col++ {
		lowest := 0
		for row := 0; row < len(grid); row++ {
			// fmt.Printf("Coords: (%v,%v)\n", row, col)
			// snapshot(puzzle.grid)
			char := grid[row][col]
			if char == '#' {
				lowest = row + 1
			} else if char == 'O' {
				if lowest < row {
					grid[row][col], grid[lowest][col] = grid[lowest][col], grid[row][col]
				}
				lowest++
			}
		}
	}
}

func RollWest(grid [][]byte) {
	for row := 0; row < len(grid); row++ {
		leastEast := 0
		for col := 0; col < len(grid[0]); col++ {
			// fmt.Printf("Coords: (%v,%v)\n", row, col)
			// snapshot(puzzle.grid)
			char := grid[row][col]
			if char == '#' {
				leastEast = col + 1
			} else if char == 'O' {
				if leastEast < col {
					grid[row][col], grid[row][leastEast] = grid[row][leastEast], grid[row][col]
				}
				leastEast++
			}
		}
	}
}

func RollSouth(grid [][]byte) {
	for col := 0; col < len(grid[0]); col++ {
		highest := len(grid) - 1
		for row := len(grid) - 1; row >= 0; row-- {
			// fmt.Printf("Coords: (%v,%v)\n", row, col)
			// snapshot(grid)
			char := grid[row][col]
			if char == '#' {
				highest = row - 1
			} else if char == 'O' {
				if highest > row {
					grid[row][col], grid[highest][col] = grid[highest][col], grid[row][col]
				}
				highest--
			}
		}
	}
}

func RollEast(grid [][]byte) {
	for row := len(grid) - 1; row >= 0; row-- {
		moreEast := len(grid) - 1
		for col := len(grid[0]) - 1; col >= 0; col-- {
			// fmt.Printf("Coords: (%v,%v)\n", row, col)
			// snapshot(grid)
			char := grid[row][col]
			if char == '#' {
				moreEast = col - 1
			} else if char == 'O' {
				if moreEast > col {
					grid[row][col], grid[row][moreEast] = grid[row][moreEast], grid[row][col]
				}
				moreEast--
			}
		}
	}
}

func cycle(grid [][]byte) {

	RollNorth(grid)
	RollWest(grid)
	RollSouth(grid)
	RollEast(grid)
}

func part1(puzzle Puzzle) int {
	RollNorth(puzzle.grid)
	return calculateScore(puzzle.grid)
}

func part2(puzzle Puzzle) int {
	state := map[string]int{}
	len := 1000_000_000
	for i := 0; i < len; i++ {
		cycle(puzzle.grid)
		h := string(bytes.Join(puzzle.grid, []byte{}))
		if _, ok := state[h]; ok {
			cycle := i - state[h]
			fit := (len - i) / cycle
			i += fit * cycle
			// for i+cycle < 1000_000_000 {
			// 	i += cycle
			// }
		}
		state[h] = i
	}

	return calculateScore(puzzle.grid)
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	puzzle2 := readInput(filename)
	return part1(puzzle), part2(puzzle2)
}
