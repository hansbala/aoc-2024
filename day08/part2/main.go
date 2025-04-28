package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	EMPTY = "."
)

type Cord struct {
	// 0-indexed.
	r int
	c int
}

type Direction struct {
	rDiff int
	cDiff int
}

func main() {
	grid := MustGetInput()
	freqToLocations := map[string][]Cord{}
	antinodes := map[Cord]bool{}
	for r, row := range grid {
		for c, item := range row {
			if item != EMPTY {
				freqToLocations[item] = append(freqToLocations[item], Cord{r: r, c: c})
			}
		}
	}
	// fmt.Printf("test -> %v\n", getAntinodes(grid, Cord{r: 1, c: 3}, Cord{r: 0, c: 0}))
	for _, locations := range freqToLocations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				// For this paring, find the corresponding antinodes.
				currAntinodes := getAntinodes(grid, locations[i], locations[j])
				// fmt.Printf("loc1 %v, loc2 %v, antinodes %v\n", locations[i], locations[j], currAntinodes)
				for _, currAntinode := range currAntinodes {
					if !isValid(grid, currAntinode) {
						continue
					}
					antinodes[currAntinode] = true
				}
			}
		}
	}
	// fmt.Printf("antinodes -> %v\n", antinodes)
	// fmt.Printf("ans -> %d\n", len(antinodes))
	for _, loc := range freqToLocations {
		if len(loc) == 1 {
			continue
		}
		for _, currLoc := range loc {
			antinodes[currLoc] = true
		}
	}
	fmt.Printf("ans -> %d\n", len(antinodes))
}

func getAntinodes(grid [][]string, loc1 Cord, loc2 Cord) []Cord {
	dirVector := Direction{
		rDiff: loc2.r - loc1.r,
		cDiff: loc2.c - loc1.c,
	}
	r := []Cord{}
	// keep doing loc1 - dir vector.
	currSub := Cord{r: loc1.r - dirVector.rDiff, c: loc1.c - dirVector.cDiff}
	for {
		if isValid(grid, currSub) {
			r = append(r, currSub)
		} else {
			break
		}
		currSub = Cord{r: currSub.r - dirVector.rDiff, c: currSub.c - dirVector.cDiff}
	}
	// keep doing loc2 + dir vector.
	currAdd := Cord{r: loc2.r + dirVector.rDiff, c: loc2.c + dirVector.cDiff}
	for {
		if isValid(grid, currAdd) {
			r = append(r, currAdd)
		} else {
			break
		}
		currAdd = Cord{r: currAdd.r + dirVector.rDiff, c: currAdd.c + dirVector.cDiff}
	}
	return r
}

func isValid(grid [][]string, cord Cord) bool {
	if cord.r < 0 || cord.r >= len(grid) || cord.c < 0 || cord.c >= len(grid[0]) {
		return false
	}
	return true
}

func MustGetInput() [][]string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := [][]string{}
	for i, line := range strings.Split(string(file), "\n") {
		r = append(r, []string{})
		for _, ch := range strings.Split(line, "") {
			r[i] = append(r[i], ch)
		}
	}
	return r
}
