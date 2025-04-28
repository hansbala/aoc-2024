package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := MustGetInput()
	// fmt.Printf("%v\n", input)
	score := 0
	for r := range input {
		for c := range input[0] {
			if input[r][c] != 0 {
				continue
			}
			score += GetScore(input, r, c)
		}
	}
	fmt.Printf("score: %d\n", score)
}

type Cord struct {
	r int
	c int
}

func GetScore(heights [][]int, row int, col int) int {
	score := 0
	visited := map[Cord]bool{{r: row, c: col}: true}
	q := []Cord{{r: row, c: col}}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		if heights[top.r][top.c] == 9 {
			score++
		}

		validNeighbors := GetNeighbors(heights, top, visited)
		for _, neighbor := range validNeighbors {
			q = append(q, neighbor)
			visited[neighbor] = true
		}
	}
	return score
}

func GetNeighbors(heights [][]int, curr Cord, visited map[Cord]bool) []Cord {
	r := []Cord{}
	diffs := []struct {
		rDiff int
		cDiff int
	}{
		{rDiff: 0, cDiff: 1},
		{rDiff: 0, cDiff: -1},
		{rDiff: 1, cDiff: 0},
		{rDiff: -1, cDiff: 0},
	}
	for _, diff := range diffs {
		newR := curr.r + diff.rDiff
		newC := curr.c + diff.cDiff
		if newR < 0 || newR >= len(heights) || newC < 0 || newC >= len(heights[0]) {
			continue
		}
		cord := Cord{r: newR, c: newC}
		if _, ok := visited[cord]; ok {
			continue
		}
		if heights[newR][newC]-heights[curr.r][curr.c] != 1 {
			continue
		}
		r = append(r, cord)
	}
	return r
}

func MustGetInput() [][]int {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(fileContent), "")
	r := [][]int{}
	r = append(r, []int{})
	lineNumber := 0
	for _, ch := range content {
		if ch == "\n" {
			lineNumber++
			r = append(r, []int{})
			continue
		}
		height, err := strconv.Atoi(ch)
		if err != nil {
			panic(err)
		}
		r[lineNumber] = append(r[lineNumber], height)
	}
	return r[:len(r)-1]
}
