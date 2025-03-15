package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cMas = "MAS"
)

type Pair struct {
	r int
	c int
}

type Direction struct {
	rDiff int
	cDiff int
}

var (
	cDirections = []Direction{
		{
			rDiff: 1,
			cDiff: 1,
		},
		{
			rDiff: -1,
			cDiff: -1,
		},
		{
			rDiff: 1,
			cDiff: -1,
		},
		{
			rDiff: -1,
			cDiff: 1,
		},
	}
)

type Match struct {
	start Pair
	end   Pair
}

func main() {
	grid := getGrid()
	rows := len(grid)
	cols := len(grid[0])
	matches := []Match{}
	for r := range rows {
		for c := range cols {
			if grid[r][c] != "M" {
				continue
			}
			for _, dir := range cDirections {
				start, end, ok := isValidMas(grid, r, c, dir)
				if !ok {
					continue
				}
				matches = append(matches, Match{start: *start, end: *end})
			}
		}
	}
	numberOfXes := getXMatches(matches)
	fmt.Printf("number of x matches: %d\n", numberOfXes)
}

func getXMatches(matches []Match) int {
	numberOfXes := 0
	visited := map[Match]bool{}
	for _, match := range matches {
		if _, ok := visited[match]; ok {
			continue
		}
		corrMatches := getCorrMatches(matches, match)
		for _, corrMatch := range corrMatches {
			if _, ok := visited[corrMatch]; ok {
				continue
			}
			// This is a valid 'X' styles MAS so add both to visited and break
			// (we don't want to consider the second corr match).
			numberOfXes++
			visited[corrMatch] = true
			visited[match] = true
			break
		}
	}
	return numberOfXes
}

func getCorrMatches(matches []Match, match Match) []Match {
	possibleCorrs := []Match{
		{
			start: Pair{r: match.start.r, c: match.start.c + 2},
			end:   Pair{r: match.end.r, c: match.end.c - 2},
		},
		{
			start: Pair{r: match.start.r, c: match.start.c - 2},
			end:   Pair{r: match.end.r, c: match.end.c + 2},
		},
		{
			start: Pair{r: match.start.r + 2, c: match.start.c},
			end:   Pair{r: match.end.r - 2, c: match.end.c},
		},
		{
			start: Pair{r: match.start.r - 2, c: match.start.c},
			end:   Pair{r: match.end.r + 2, c: match.end.c},
		},
	}
	r := []Match{}
	for _, m := range matches {
		for _, possibleCorr := range possibleCorrs {
			if possibleCorr == m {
				r = append(r, possibleCorr)
			}
		}
	}
	return r
}

func isValidMas(grid [][]string, r int, c int, dir Direction) (*Pair /*start*/, *Pair /*end*/, bool /*isValid*/) {
	if grid[r][c] != "M" {
		panic("expected starting at X baby")
	}
	start := Pair{r: r, c: c}
	end := Pair{}
	lookFor := "A"
	for {
		nextR, nextC := r+dir.rDiff, c+dir.cDiff
		if nextR < 0 || nextR >= len(grid) || nextC < 0 || nextC >= len(grid[0]) {
			return nil, nil, false
		}
		if grid[nextR][nextC] != lookFor {
			return nil, nil, false
		}
		if lookFor == "S" {
			end = Pair{r: nextR, c: nextC}
			break
		}
		r, c = nextR, nextC
		lookFor = string(cMas[strings.Index(cMas, lookFor)+1])
	}
	return &start, &end, true
}

func getGrid() [][]string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	for lineIdx, line := range strings.Split(string(fileContent), "\n") {
		grid = append(grid, []string{})
		characters := strings.Split(line, "")
		for _, ch := range characters {
			grid[lineIdx] = append(grid[lineIdx], ch)
		}
	}
	return grid
}
