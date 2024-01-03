package day24

import (
	"bytes"
	"os"
	"strconv"
	"strings"
)

type Hail struct {
	x, y, z, dx, dy, dz float64
	a, b, c             float64
}

func newHail(x, y, z, dx, dy, dz float64) Hail {
	a := dy
	b := -dx
	c := dy*x - dx*y
	return Hail{x, y, z, dx, dy, dz, a, b, c}
}

type Puzzle struct {
	hails []Hail
}

func readPosition(input []byte) (float64, float64, float64) {
	items := bytes.Split(input, []byte(","))
	a, _ := strconv.ParseFloat(strings.TrimSpace(string(items[0])), 64)
	b, _ := strconv.ParseFloat(strings.TrimSpace(string(items[1])), 64)
	c, _ := strconv.ParseFloat(strings.TrimSpace(string(items[2])), 64)
	return float64(a), float64(b), float64(c)
}
func readInput(filename string) Puzzle {
	result := []Hail{}

	input, _ := os.ReadFile(filename)
	lines := bytes.Split(input, []byte("\n"))
	for _, line := range lines {
		pp := bytes.Split(line, []byte("@"))
		x, y, z := readPosition(pp[0])
		dx, dy, dz := readPosition(pp[1])
		result = append(result, newHail(x, y, z, dx, dy, dz))
	}

	return Puzzle{result}
}

func part1(puzzle Puzzle) int {
	total := 0
	for i := 0; i < len(puzzle.hails); i++ {
		for j := 0; j < i; j++ {
			hs1 := puzzle.hails[i]
			hs2 := puzzle.hails[j]
			a1, b1, c1 := hs1.a, hs1.b, hs1.c
			a2, b2, c2 := hs2.a, hs2.b, hs2.c
			if a1*b2 == b1*a2 {
				continue
			}
			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)
			if (x-hs1.x)*hs1.dx >= 0 && (y-hs1.y)*hs1.dy >= 0 && (x-hs2.x)*hs2.dx >= 0 && (y-hs2.y)*hs2.dy >= 0 {
				if 200000000000000 <= x && x <= 400000000000000 && 200000000000000 <= y && y <= 400000000000000 {
					total++
				}
			}
		}
	}
	return total
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), "won't solve, this is stupid"
}
