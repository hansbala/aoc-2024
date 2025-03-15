package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cXmas = "XMAS"
)

type Direction struct {
	rDiff int
	cDiff int
}

var (
	cDirections = []Direction{
		{
			rDiff: -1,
			cDiff: 0,
		},
		{
			rDiff: 0,
			cDiff: -1,
		},
		{
			rDiff: 1,
			cDiff: 0,
		},
		{
			rDiff: 0,
			cDiff: 1,
		},
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

type Element struct {
	// 0-indexed.
	row int
	col int
	val string // 'X', 'M', 'A', or 'S'
}

type Pair struct {
	r int
	c int
}

func main() {
	grid := getGrid()
	rows := len(grid)
	cols := len(grid[0])
	total := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] != "X" {
				continue
			}
			for _, dir := range cDirections {
				if !isValidXmas(grid, r, c, dir) {
					continue
				}
				total += 1
			}
		}
	}
	fmt.Printf("total: %d\n", total)
}

func isValidXmas(grid [][]string, r int, c int, dir Direction) bool {
	if grid[r][c] != "X" {
		panic("expected starting at X baby")
	}
	lookFor := "M"
	for {
		nextR, nextC := r+dir.rDiff, c+dir.cDiff
		if nextR < 0 || nextR >= len(grid) || nextC < 0 || nextC >= len(grid[0]) {
			return false
		}
		if grid[nextR][nextC] != lookFor {
			return false
		}
		if lookFor == "S" {
			break
		}
		r, c = nextR, nextC
		lookFor = string(cXmas[strings.Index(cXmas, lookFor)+1])
	}
	return true
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

// func bfs(grid [][]string, r int, c int) int {
// 	if grid[r][c] != "X" {
// 		panic("expected grid[r][c] to be 'X'")
// 	}
// 	q := []Element{{row: r, col: c, val: "X"}}
// 	visited := map[Pair]bool{}
// 	total := 0
// 	for {
// 		if len(q) == 0 {
// 			break
// 		}
// 		top := q[0]
// 		q = q[1:]
// 		visited[Pair{r: top.row, c: top.col}] = true
// 		if top.val == "X" || top.val == "M" || top.val == "A" {
// 			// Look for neighbors that have the next character.
// 			neighbors := getValidNeighbors(grid, top.row, top.col, visited, getNextChar(top.val))
// 			for _, neighbor := range neighbors {
// 				q = append(q, neighbor)
// 			}
// 		} else if top.val == "S" {
// 			// Base
// 			total++
// 		} else {
// 			// We don't want these guys.
// 			continue
// 		}
// 	}
// 	return total
// }

// func getValidNeighbors(grid [][]string, row int, col int, visited map[Pair]bool, lookFor string) []Element {
// 	if !(lookFor == "X" || lookFor == "M" || lookFor == "A" || lookFor == "S") {
// 		panic(fmt.Errorf("wrong character for getValidNeighbors. Received: %v", lookFor))
// 	}
// 	r := []Element{}
// 	for _, dir := range cDirections {
// 		newR, newC := row+dir.rDiff, col+dir.cDiff
// 		if newR < 0 || newR >= len(grid) || newC < 0 || newC >= len(grid[0]) {
// 			continue
// 		}
// 		if _, ok := visited[Pair{r: newR, c: newC}]; ok {
// 			continue
// 		}
// 		if grid[newR][newC] != lookFor {
// 			continue
// 		}
// 		r = append(r, Element{row: newR, col: newC, val: lookFor})
// 	}
// 	fmt.Printf("neighbors at (%d, %d) are: %v\n", row, col, r)
// 	return r
// }

// func getNextChar(s string) string {
// 	if !(s == "X" || s == "M" || s == "A") {
// 		panic("wrong character for getNextChar")
// 	}
// 	return string(cXmas[strings.Index(cXmas, s)+1])
// }
