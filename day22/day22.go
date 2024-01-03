package day22

import (
	"bytes"
	"os"
	"sort"
	"strconv"
)

type Brick struct {
	startX, startY, startZ int
	endX, endY, endZ       int
}

type Puzzle struct {
	bricks []Brick
}

func readPosition(input []byte) (int, int, int) {
	items := bytes.Split(input, []byte(","))
	a, _ := strconv.ParseInt(string(items[0]), 10, 32)
	b, _ := strconv.ParseInt(string(items[1]), 10, 32)
	c, _ := strconv.ParseInt(string(items[2]), 10, 32)
	return int(a), int(b), int(c)
}

func readInput(filename string) Puzzle {
	input, _ := os.ReadFile(filename)
	lines := bytes.Fields(input)
	bricks := []Brick{}
	for _, line := range lines {
		items := bytes.SplitN(line, []byte("~"), 2)
		startX, startY, startZ := readPosition(items[0])
		endX, endY, endZ := readPosition(items[1])
		bricks = append(bricks, Brick{startX, startY, startZ, endX, endY, endZ})
	}
	sort.Slice(bricks, func(i, j int) bool {
		if bricks[i].startZ > bricks[i].endZ {
			panic("assumption")
		}
		if bricks[j].startZ > bricks[j].endZ {
			panic("assumption")
		}
		return bricks[i].startZ < bricks[j].startZ
	})
	return Puzzle{bricks}
}

func brickInPos(x, y, z int, bricks []Brick) bool {
	for _, brick := range bricks {
		if x >= brick.startX && x <= brick.endX && y >= brick.startY && y <= brick.endY && z >= brick.startZ && z <= brick.endZ {
			return true
		}
	}
	return false
}

func dropDiff(tested Brick, bricks []Brick) int {
	if tested.startZ == 1 {
		return 0
	}

	diff := 1
	for {
		for x := tested.startX; x <= tested.endX; x++ {
			for y := tested.startY; y <= tested.endY; y++ {
				if brickInPos(x, y, tested.startZ-diff, bricks) {
					return diff - 1
				}
			}
		}
		if tested.startZ-diff == 1 {
			return diff
		}
		diff++
	}
}
func optimize(bricks []Brick) int {
	count := 0
	for {
		changed := false
		for i, brick := range bricks {
			dd := dropDiff(brick, bricks)
			if dd > 0 {
				changed = true
			} else {
				continue
			}
			brick.startZ -= dd
			brick.endZ -= dd
			bricks[i] = brick
			count++
		}
		if !changed {
			break
		}
	}
	return count
}

func chainreaction(index int, bricks []Brick) int {
	nBricks := []Brick{}
	for i, v := range bricks {
		if i != index {
			nBricks = append(nBricks, v)
		}
	}
	return optimize(nBricks)
}

func totalResuls(puzzle Puzzle) (int, int) {
	optimize(puzzle.bricks)
	sum := 0
	count := 0
	for i := 0; i < len(puzzle.bricks); i++ {
		dd := chainreaction(i, puzzle.bricks)
		sum += dd
		if dd == 0 {
			count++
		}
	}
	return sum, count
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return totalResuls(puzzle)
}
