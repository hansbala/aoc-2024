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
	for _, locations := range freqToLocations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				// For this paring, find the corresponding antinodes.
				currAntinodes := getAntinodes(locations[i], locations[j])
				// fmt.Printf("loc1 %v, loc2 %v, antinodes %v\n", locations[i], locations[j], currAntinodes)
				for _, currAntinode := range currAntinodes {
					if currAntinode.r < 0 || currAntinode.r >= len(grid) || currAntinode.c < 0 || currAntinode.c >= len(grid[0]) {
						continue
					}
					antinodes[currAntinode] = true
				}
			}
		}
	}
	// fmt.Printf("antinodes -> %v\n", antinodes)
	fmt.Printf("ans -> %d\n", len(antinodes))
}

func getAntinodes(loc1 Cord, loc2 Cord) []Cord {
	dirVector := Direction{
		rDiff: loc2.r - loc1.r,
		cDiff: loc2.c - loc1.c,
	}
	// loc1 - dir vector
	// loc2 + dir vector
	return []Cord{
		{r: loc1.r - dirVector.rDiff, c: loc1.c - dirVector.cDiff},
		{r: loc2.r + dirVector.rDiff, c: loc2.c + dirVector.cDiff},
	}
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
