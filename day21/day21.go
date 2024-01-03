package day21

// Thanks: https://www.youtube.com/watch?v=9UOMZSL0JTg&t=1192s
import (
	"bytes"
	"fmt"
	"os"
)

type Point struct {
	row int
	col int
}

type Puzzle struct {
	grid  [][]byte
	start Point
}

func readInput(file string) Puzzle {
	input, _ := os.ReadFile(file)
	lines := bytes.Split(input, []byte("\n"))
	grid := [][]byte{}
	startX := 0
	startY := 0
	for row := 0; row < len(lines); row++ {
		nRow := []byte{}
		for col := 0; col < len(lines[0]); col++ {
			if lines[row][col] == 'S' {
				startX = row
				startY = col
				nRow = append(nRow, '.')
			} else {
				nRow = append(nRow, lines[row][col])
			}
		}
		grid = append(grid, nRow)
	}
	return Puzzle{grid, Point{startX, startY}}
}

func findNeighbors(grid [][]byte, visisted map[string]int, start Point) []Point {
	result := []Point{}
	nRow := start.row - 1
	if nRow >= 0 && grid[nRow][start.col] == '.' {
		p := Point{nRow, start.col}
		key := fmt.Sprintf("%v,%v", p.row, p.col)
		_, ok := visisted[key]
		if !ok {
			result = append(result, p)
		}
	}
	sRow := start.row + 1
	if sRow < len(grid) && grid[sRow][start.col] == '.' {
		p := Point{sRow, start.col}
		key := fmt.Sprintf("%v,%v", p.row, p.col)
		_, ok := visisted[key]
		if !ok {
			result = append(result, p)
		}
	}
	eCol := start.col + 1
	if eCol < len(grid[0]) && grid[start.row][eCol] == '.' {
		p := Point{start.row, eCol}
		key := fmt.Sprintf("%v,%v", p.row, p.col)
		_, ok := visisted[key]
		if !ok {
			result = append(result, p)
		}
	}

	wCol := start.col - 1
	if wCol >= 0 && grid[start.row][wCol] == '.' {
		p := Point{start.row, wCol}
		key := fmt.Sprintf("%v,%v", p.row, p.col)
		_, ok := visisted[key]
		if !ok {
			result = append(result, p)
		}
	}

	return result
}

func removeDuplicate[T Point](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func fill(grid [][]byte, start Point, limit int) int {
	key := fmt.Sprintf("%v,%v", start.row, start.col)
	visited := map[string]int{key: 0}
	currenDepth := []Point{start}
	counter := limit
	for {
		if counter < 0 || len(currenDepth) == 0 {
			break
		}
		nPaths := []Point{}

		//Proccess
		for _, node := range currenDepth {
			k := fmt.Sprintf("%v,%v", node.row, node.col)
			visited[k] = limit - counter
			newitems := findNeighbors(grid, visited, node)
			nPaths = append(nPaths, newitems...)
		}
		nPaths = removeDuplicate(nPaths)
		currenDepth = nPaths
		counter--
	}

	sum1 := 0
	sum2 := 0
	for _, v := range visited {
		if v%2 == 0 {
			sum1++
		} else {
			sum2++
		}
	}
	if limit%2 == 0 {
		return sum1
	}
	return sum2
}

func part1(puzzle Puzzle) int {
	sum := fill(puzzle.grid, puzzle.start, 64)

	// fmt.Printf("Part 1: %v\n", visited)
	return sum
}

var steps = 26501365

func part2(puzzle Puzzle) int {
	if len(puzzle.grid) != len(puzzle.grid[0]) {
		panic("square grid")
	}

	size := len(puzzle.grid)
	if puzzle.start.row != puzzle.start.col || puzzle.start.row != size/2 {
		panic("start at the middle")
	}
	if steps%size != size/2 {
		panic("grid steps is a multiple of size/2")
	}

	grid_width := steps/size - 1

	odd := (grid_width/2*2 + 1) * (grid_width/2*2 + 1)
	even := ((grid_width + 1) / 2 * 2) * ((grid_width + 1) / 2 * 2)

	odd_points := fill(puzzle.grid, puzzle.start, size*2+1)
	even_points := fill(puzzle.grid, puzzle.start, size*2)

	corner_t := fill(puzzle.grid, Point{size - 1, puzzle.start.col}, size-1)
	corner_b := fill(puzzle.grid, Point{puzzle.start.row, 0}, size-1)
	corner_l := fill(puzzle.grid, Point{0, puzzle.start.col}, size-1)
	corner_r := fill(puzzle.grid, Point{puzzle.start.row, size - 1}, size-1)

	small_tr := fill(puzzle.grid, Point{size - 1, 0}, size/2-1)
	small_tl := fill(puzzle.grid, Point{size - 1, size - 1}, size/2-1)
	small_br := fill(puzzle.grid, Point{0, 0}, size/2-1)
	small_bl := fill(puzzle.grid, Point{0, size - 1}, size/2-1)

	large_tr := fill(puzzle.grid, Point{size - 1, 0}, (size*3)/2-1)
	large_tl := fill(puzzle.grid, Point{size - 1, size - 1}, (size*3)/2-1)
	large_br := fill(puzzle.grid, Point{0, 0}, (size*3)/2-1)
	large_bl := fill(puzzle.grid, Point{0, size - 1}, (size*3)/2-1)

	sum := odd*odd_points +
		even*even_points +
		corner_t + corner_r + corner_b + corner_l +
		(grid_width+1)*(small_tr+small_tl+small_br+small_bl) +
		grid_width*(large_tr+large_tl+large_br+large_bl)

	return sum
}

func Solve(filename string) (interface{}, interface{}) {

	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}
