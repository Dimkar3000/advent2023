package day17

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	NORTH = 1
	SOUTH = 2
	EAST  = 3
	WEST  = 4
)

type Path struct {
	x, y, direction, directionLegth, cost int
}

type Puzzle struct {
	grid [][]int
}

func exists(x, y int, grid [][]int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func getNewDirectionLength(currentDirection, currentDirectionLegth, newDirection int) int {

	if currentDirection == newDirection {
		return currentDirectionLegth + 1
	}

	return 1
}

func generateKey(path Path) string {
	key := fmt.Sprintf("%v,%v,%v, %v", path.x, path.y, path.direction, path.directionLegth)
	return key
}

func generatePath(visited map[string]int, x, y, newDirection, currentDirection, currentDirectionLegth, currentCost int, grid [][]int) (Path, bool) {
	if !exists(x, y, grid) {
		return Path{}, false
	}
	if currentDirection == newDirection && currentDirectionLegth > 9 {
		return Path{}, false
	}
	if currentDirection != newDirection && currentDirectionLegth < 4 {
		return Path{}, false
	}
	if currentDirection == SOUTH && newDirection == NORTH {
		return Path{}, false
	}
	if currentDirection == NORTH && newDirection == SOUTH {
		return Path{}, false
	}
	if currentDirection == WEST && newDirection == EAST {
		return Path{}, false
	}
	if currentDirection == EAST && newDirection == WEST {
		return Path{}, false
	}
	result := Path{
		x:              x,
		y:              y,
		direction:      newDirection,
		directionLegth: getNewDirectionLength(currentDirection, currentDirectionLegth, newDirection),
		cost:           currentCost + grid[x][y],
	}
	key := generateKey(result)
	if _, ok := visited[key]; ok {
		return Path{}, false
	}
	return result, true
}

func generatePaths(visited map[string]int, x, y, currentDirectionLength, currentDirection, cost int, grid [][]int) []Path {
	result := []Path{}

	north, ok := generatePath(visited, x-1, y, NORTH, currentDirection, currentDirectionLength, cost, grid)
	if ok {
		result = append(result, north)
	}

	south, ok := generatePath(visited, x+1, y, SOUTH, currentDirection, currentDirectionLength, cost, grid)
	if ok {
		result = append(result, south)
	}

	west, ok := generatePath(visited, x, y-1, WEST, currentDirection, currentDirectionLength, cost, grid)
	if ok {
		result = append(result, west)
	}

	east, ok := generatePath(visited, x, y+1, EAST, currentDirection, currentDirectionLength, cost, grid)
	if ok {
		result = append(result, east)
	}

	return result
}

func readFile(input string) Puzzle {
	grid := [][]int{}

	file, _ := os.ReadFile(input)
	for _, line := range bytes.Split(file, []byte("\n")) {
		row := []int{}
		for _, v := range line {
			n, _ := strconv.Atoi(string(v))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	return Puzzle{grid: grid}
}

func splitPaths(paths []Path) ([]Path, []Path) {

	sort.Slice(paths, func(i, j int) bool {
		return paths[i].cost < paths[j].cost
	})
	limit := paths[0].cost
	limitIndex := 0
	for i := 0; i < len(paths); i++ {
		if paths[i].cost == limit {
			limitIndex = i
		} else {
			break
		}
	}
	limitIndex++
	return paths[:limitIndex], paths[limitIndex:]
}

func part1(puzzle Puzzle) int {

	currentClosest := []Path{
		{
			x:              0,
			y:              0,
			direction:      EAST,
			directionLegth: 0,
			cost:           0,
		},
	}
	puzzle.grid[0][0] = 0
	restPaths := []Path{}
	visited := map[string]int{}
	for len(currentClosest) > 0 {
		for _, path := range currentClosest {
			key := generateKey(path)
			_, ok := visited[key]
			if ok {
				continue
			}
			visited[key] = path.cost
			// fmt.Printf("(%v,%v): %v\r", path.x, path.y, path.cost)
			restPaths = append(restPaths, generatePaths(visited, path.x, path.y, path.directionLegth, path.direction, path.cost, puzzle.grid)...)
		}

		ncurrentClosest, nrestrestPaths := splitPaths(restPaths)
		currentClosest = ncurrentClosest
		restPaths = nrestrestPaths

		for _, path := range currentClosest {
			if path.x == len(puzzle.grid)-1 && path.y == len(puzzle.grid[0])-1 && path.directionLegth >= 4 {
				return path.cost
			}
		}
	}
	return -1
}

func Solve(filename string) (interface{}, interface{}) {
	grid := readFile(filename)
	return part1(grid), "(won't do)"
}
