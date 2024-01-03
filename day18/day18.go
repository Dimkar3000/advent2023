package day18

import (
	"bytes"
	"os"
	"strconv"
)

type Point struct {
	x, y int64
}
type Puzzle struct {
	points []Point
}

func readInput1(file string) Puzzle {
	result := Puzzle{}
	input, _ := os.ReadFile(file)
	lines := bytes.Split(input, []byte("\n"))
	x := int64(0)
	y := int64(0)
	result.points = append(result.points, Point{x, y})
	for _, line := range lines {
		parts := bytes.Split(line, []byte(" "))
		direction := string(parts[0])
		count, _ := strconv.ParseInt(string(parts[1]), 10, 64)
		if direction == "R" {
			x += count
		} else if direction == "L" {
			x -= count
		} else if direction == "D" {
			y += count
		} else if direction == "U" {
			y -= count
		}
		result.points = append(result.points, Point{x, y})
	}
	return result
}

func readInput2(file string) Puzzle {
	result := Puzzle{}
	input, _ := os.ReadFile(file)
	lines := bytes.Split(input, []byte("\n"))
	x := int64(0)
	y := int64(0)
	result.points = append(result.points, Point{x, y})
	for _, line := range lines {
		parts := bytes.Split(line, []byte("#"))
		count, _ := strconv.ParseInt(string(parts[1][:5]), 16, 64)
		direction := parts[1][5]
		if direction == '0' {
			x += count
		} else if direction == '2' {
			x -= count
		} else if direction == '1' {
			y += count
		} else if direction == '3' {
			y -= count
		}
		result.points = append(result.points, Point{x, y})
	}
	return result
}

func calculate_outline(points []Point) int64 {
	sum := int64(0)
	for i := 0; i < len(points)-1; i++ {
		xDiff := (points[i].x - points[i+1].x)
		yDiff := (points[i].y - points[i+1].y)
		if xDiff < 0 {
			xDiff *= -1
		}
		if yDiff < 0 {
			yDiff *= -1
		}
		sum += xDiff + yDiff
	}

	return sum
}

func solve(puzzle Puzzle) int64 {
	area := int64(0)
	for i := 0; i < len(puzzle.points)-1; i++ {
		area += puzzle.points[i].x*puzzle.points[i+1].y - puzzle.points[i+1].x*puzzle.points[i].y
	}
	area /= 2
	if area < 0 {
		area *= -1
	}
	outline := calculate_outline(puzzle.points)
	inner := int64(float64(area) - float64(outline)/2 + 1)
	return inner + outline

}

func Solve(path string) (interface{}, interface{}) {
	puzzle1 := readInput1(path)
	puzzle2 := readInput2(path)
	r1 := solve(puzzle1)
	r2 := solve(puzzle2)
	return r1, r2
}
