package day5

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Puzzle struct {
	seeds      []int
	categories []Category
}

type Category struct {
	ranges []Range
}

type Range struct {
	dest   int
	source int
	size   int
}

func readInput1(filename string) Puzzle {
	readFile, err := os.ReadFile(filename)
	input := strings.Split(string(readFile), "\n\n")
	if err != nil {
		fmt.Println(err)
	}

	results := []Category{}
	for _, v := range input[1:] {
		category := Category{}
		lines := strings.Split(v, "\n")
		for _, line := range lines[1:] {
			rr := Range{}
			numbers := strings.Split(line, " ")
			rr.dest, _ = strconv.Atoi(numbers[0])
			rr.source, _ = strconv.Atoi(numbers[1])
			rr.size, _ = strconv.Atoi(numbers[2])

			category.ranges = append(category.ranges, rr)
		}

		results = append(results, category)
	}

	seeds := []int{}
	for _, seed := range strings.Split(input[0], " ") {
		s, e := strconv.Atoi(seed)
		if e != nil {
			continue
		}
		seeds = append(seeds, s)
	}

	return Puzzle{
		seeds:      seeds,
		categories: results,
	}

}

func mapSeeds(seeds *[]int, category *Category) []int {
	result := []int{}
	for _, seed := range *seeds {
		result = append(result, mapSeed(seed, category))
	}
	return result
}

func mapSeed(seed int, category *Category) int {
	for _, rr := range category.ranges {
		if rr.source <= seed && seed < (rr.source+rr.size) {
			diff := seed - rr.source
			return rr.dest + diff
		}
	}
	return seed
}

type BySource []Range

func (a BySource) Len() int           { return len(a) }
func (a BySource) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySource) Less(i, j int) bool { return a[i].source < a[j].source }

func part1(inputs Puzzle) int {

	results := inputs.seeds

	for _, category := range inputs.categories {
		results = mapSeeds(&results, &category)
	}
	min := results[0]
	for _, v := range results {
		if v < min {
			min = v
		}
	}
	return min
}

func generateNumbers(ch chan<- int, wg *sync.WaitGroup, inputs *Puzzle, i int) {
	defer wg.Done()
	min := math.MaxInt
	for j := (*inputs).seeds[i]; j < (*inputs).seeds[i]+(*inputs).seeds[i+1]; j++ {
		result := j
		for _, category := range (*inputs).categories {
			result = mapSeed(result, &category)
		}
		if result < min {
			min = result
		}
	}

	ch <- min
}

func part2(inputs Puzzle) int {
	var wg sync.WaitGroup
	numberChan := make(chan int, len(inputs.seeds)/2)
	min := math.MaxInt
	for i := 0; i < len(inputs.seeds)-1; i += 2 {
		wg.Add(1)
		go generateNumbers(numberChan, &wg, &inputs, i)
	}
	wg.Wait()
	close(numberChan)
	for mins := range numberChan {
		if mins < min {
			min = mins
		}
	}
	return min
}

func Solve(filename string) (interface{}, interface{}) {
	input1 := readInput1(filename)
	return part1(input1), part2(input1)
}
