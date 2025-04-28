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

func main() {
	grid := MustGetInput()
	// fmt.Printf("Grid: %v\n", grid)
	visited := map[Coord]bool{}
	currCoord := MustGetStartingCoord(grid)
	grid[currCoord.r][currCoord.c] = EMPTY
	currDir := UP
	for {
		visited[currCoord] = true
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
	fmt.Printf("unique coords -> %d\n", len(visited))
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
