package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	hand string
	bid  int
}

type Puzzle struct {
	games []Game
}

func readInput(file string) Puzzle {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	games := []Game{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		g := Game{
			hand: parts[0],
			bid:  bid,
		}
		games = append(games, g)
	}

	readFile.Close()

	return Puzzle{
		games: games,
	}
}

func countCards(cards []rune) map[rune]int {
	results := map[rune]int{}
	for _, v := range cards {
		if _, ok := results[v]; !ok {
			results[v] = 1
		} else {
			results[v]++
		}
	}
	return results
}

func isSingePair(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	for _, count := range counts {
		if count == 2 {
			return true
		}
	}
	if !isPart2 {
		return false

	}
	jokers := counts['J']
	return jokers == 1
}

func isTwoPair(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	sum := 0
	for _, count := range counts {
		if count == 2 {
			sum++
		}
	}
	return sum == 2
}

func isThreeOfAKind(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	hastriplet := false
	for _, count := range counts {
		if count == 3 {
			hastriplet = true
		}
	}
	if !isPart2 {
		return hastriplet

	}
	jokers := counts['J']
	if jokers == 0 {
		return hastriplet
	}
	for card, count := range counts {
		if card == 'J' {
			continue
		}
		if count == 2 && jokers == 1 {
			return true
		}
	}
	return jokers == 2
}

func isFullHouse(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	r := hasThree && hasTwo
	if r {
		return true
	}
	if !isPart2 {
		return false
	}
	jokers := counts['J']
	if jokers == 0 {
		return false
	}
	pairs := 0
	for card, count := range counts {
		if card == 'J' {
			continue
		}
		if count == 2 {
			pairs++
		}
	}

	return pairs == 2 && jokers == 1

}

func isFourOfAKind(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	for _, count := range counts {
		if count == 4 {
			return true
		}
	}
	if !isPart2 {
		return false
	}
	jokers := counts['J']
	if jokers == 0 {
		return false
	}
	for card, count := range counts {
		if card == 'J' {
			continue
		}
		if count+jokers == 4 {
			return true
		}
	}

	return false
}

func isFiveOfAKind(hand string, isPart2 bool) bool {
	cards := []rune(hand)
	counts := countCards(cards)
	if counts[cards[0]] == 5 {
		return true
	}
	if !isPart2 {
		return false
	}
	jokers := counts['J']
	if jokers == 0 {
		return false
	}
	for card, count := range counts {
		if card == 'J' {
			continue
		}
		if count+jokers == 5 {
			return true
		}
	}

	return jokers == 5
}

var data = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var data2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func compareCards(handA, handB string, mapped *map[rune]int) bool {
	a := []rune(handA)
	b := []rune(handB)
	for i := 0; i < len(handA); i++ {
		if (*mapped)[a[i]] == (*mapped)[b[i]] {
			continue
		}
		return (*mapped)[a[i]] > (*mapped)[b[i]]
	}
	return false
}

func compareGames(a, b *Game, mapped *map[rune]int, isPart2 bool) bool {
	handlers := []func(hand string, isPart2 bool) bool{isFiveOfAKind, isFourOfAKind, isFullHouse, isThreeOfAKind, isTwoPair, isSingePair}

	for _, handler := range handlers {
		aH := handler(a.hand, isPart2)
		bH := handler(b.hand, isPart2)
		if aH == bH && aH {
			return compareCards(a.hand, b.hand, mapped)
		}
		if aH {
			return true
		}
		if bH {
			return false
		}
	}
	return compareCards(a.hand, b.hand, mapped)
}

func part1(puzzle Puzzle) int {
	sort.Slice(puzzle.games, func(a int, b int) bool {
		return !compareGames(&puzzle.games[a], &puzzle.games[b], &data, false)
	})
	score := 0
	for i := 0; i < len(puzzle.games); i++ {
		score += puzzle.games[i].bid * (i + 1)
	}

	return score
}

func part2(puzzle Puzzle) int {
	sort.Slice(puzzle.games, func(a int, b int) bool {
		return !compareGames(&puzzle.games[a], &puzzle.games[b], &data2, true)
	})
	score := 0
	for i := 0; i < len(puzzle.games); i++ {
		score += puzzle.games[i].bid * (i + 1)
	}
	return score
}

func Solve(filename string) (interface{}, interface{}) {
	puzzle := readInput(filename)
	return part1(puzzle), part2(puzzle)
}
