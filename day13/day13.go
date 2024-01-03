package day13

import (
	"os"
	"strings"
)

type Grid [][]rune

type Puzzle struct {
	grids []Grid
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)
	parts := strings.Split(string(input), "\n\n")
	grids := []Grid{}

	for _, grid := range parts {
		lines := strings.Split(grid, "\n")
		items := [][]rune{}
		for _, line := range lines {
			l := []rune(line)
			items = append(items, l)
		}
		grids = append(grids, items)
	}

	return Puzzle{grids: grids}
}
func isValidRow(grid Grid, mirror int, smudges int) bool {
	count := smudges
	for row := 1; mirror-row >= 0 && mirror+row+1 < len(grid); row++ {
		left := mirror - row
		right := mirror + row + 1
		for col := 0; col < len(grid[0]); col++ {
			if grid[left][col] != grid[right][col] {
				if count == 0 {
					return false
				} else {
					count--
				}
			}
		}
	}
	return true
}

func isValidCol(grid Grid, mirror int, smudges int) bool {
	count := smudges
	for col := 1; mirror-col >= 0 && mirror+1+col < len(grid[0]); col++ {
		left := mirror - col
		right := mirror + col + 1
		for row := 0; row < len(grid); row++ {
			if grid[row][left] != grid[row][right] {
				if count == 0 {
					return false
				} else {
					count--
				}
			}
		}
	}
	return smudges == 0
}

func part1(puzzle Puzzle) int {

	rows := 0
	cols := 0

	for _, grid := range puzzle.grids {
		for row := 0; row < len(grid)-1; row++ {
			if grid[row][0] == grid[row+1][0] {
				reflection := true
				for col := 1; col < len(grid[0]); col++ {
					if grid[row][col] != grid[row+1][col] {
						reflection = false
						break
					}
				}
				if reflection && isValidRow(grid, row, 0) {
					rows += row + 1
					break
				}
			}
		}

		for col := 0; col < len(grid[0])-1; col++ {
			if grid[0][col] == grid[0][col+1] {
				reflection := true
				for row := 1; row < len(grid); row++ {
					if grid[row][col] != grid[row][col+1] {
						reflection = false
						break
					}
				}
				if reflection && isValidCol(grid, col, 0) {
					cols += col + 1
					break
				}
			}
		}

	}

	return cols + 100*rows
}

func part2(puzzle Puzzle) int {

	rows := 0
	cols := 0

	for _, grid := range puzzle.grids {
		for row := 0; row < len(grid)-1; row++ {
			reflection := 0
			if grid[row][0] == grid[row+1][0] || reflection == 0 {
				for col := 0; col < len(grid[0]); col++ {
					if grid[row][col] != grid[row+1][col] {
						reflection++
					}
				}
				if reflection <= 1 && isValidRow(grid, row, 1-reflection) {
					rows += row + 1
					break
				}
			}
		}

		for col := 0; col < len(grid[0])-1; col++ {
			reflection := 0
			if grid[0][col] == grid[0][col+1] || reflection == 0 {
				for row := 0; row < len(grid); row++ {
					if grid[row][col] != grid[row][col+1] {
						reflection++
					}
				}
				if reflection <= 1 && isValidCol(grid, col, 1-reflection) {
					cols += col + 1
					break
				}
			}
		}

	}

	return cols + 100*rows
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}
