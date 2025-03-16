package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	WALL  = "#"
	EMPTY = "."
	GUARD = "^"
)

var (
	UP          = Direction{rDiff: -1, cDiff: 0}
	DOWN        = Direction{rDiff: 1, cDiff: 0}
	LEFT        = Direction{rDiff: 0, cDiff: -1}
	RIGHT       = Direction{rDiff: 0, cDiff: 1}
	cDirections = []Direction{UP, DOWN, LEFT, RIGHT}
)

type Coord struct {
	r int
	c int
}

type Direction struct {
	rDiff int
	cDiff int
}

type CoordWithDir struct {
	coord Coord
	dir   Direction
}

func main() {
	grid := MustGetInput()
	ans := 0
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] != EMPTY {
				continue
			}
			grid[r][c] = WALL
			if IsCyclicGrid(grid) {
				ans++
			}
			grid[r][c] = EMPTY
		}
	}
	fmt.Printf("ans -> %d\n", ans)
}

func IsCyclicGrid(grid [][]string) bool {
	visited := map[CoordWithDir]bool{}
	guardPos := MustGetStartingCoord(grid)
	currCoord := Coord{r: guardPos.r, c: guardPos.c}
	grid[currCoord.r][currCoord.c] = EMPTY
	currDir := UP
	for {
		if _, ok := visited[CoordWithDir{coord: currCoord, dir: currDir}]; ok {
			grid[guardPos.r][guardPos.c] = GUARD
			return true
		}
		visited[CoordWithDir{coord: currCoord, dir: currDir}] = true
		nextCoord := GetNext(currCoord, currDir)
		if nextCoord.r < 0 || nextCoord.r >= len(grid) || nextCoord.c < 0 || nextCoord.c >= len(grid[0]) {
			break
		}
		if grid[nextCoord.r][nextCoord.c] == WALL {
			currDir = MustTurnRight(currDir)
			continue
		}
		currCoord = nextCoord
	}
	grid[guardPos.r][guardPos.c] = GUARD
	return false
}

func GetNext(coord Coord, dir Direction) Coord {
	return Coord{r: coord.r + dir.rDiff, c: coord.c + dir.cDiff}
}

func MustTurnRight(dir Direction) Direction {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		panic("unknown direction!")
	}
}

func MustGetStartingCoord(grid [][]string) Coord {
	for r, row := range grid {
		for c, ch := range row {
			if ch == GUARD {
				return Coord{r: r, c: c}
			}
		}
	}
	panic("could not find guard!")
}

func MustGetInput() [][]string {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	for i, line := range strings.Split(string(contents), "\n") {
		grid = append(grid, []string{})
		for _, ch := range strings.Split(line, "") {
			grid[i] = append(grid[i], ch)
		}
	}
	return grid
}
