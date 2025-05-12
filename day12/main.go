package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	cDirs = []struct {
		rDiff int
		cDiff int
	}{
		{rDiff: 1, cDiff: 0},
		{rDiff: -1, cDiff: 0},
		{rDiff: 0, cDiff: 1},
		{rDiff: 0, cDiff: -1},
	}
)

func main() {
	input := MustGetInput()
	garden := GetGarden(input)
	r := 0
	for _, regions := range garden.regionsMap {
		for _, region := range regions {
			perimeter := region.Perimeter(input)
			area := region.Area()
			total := perimeter * area
			fmt.Printf("region %s has a cost of %d (area: %d, perimeter: %d)\n", region.plantType, total, area, perimeter)
			r += region.Perimeter(input) * region.Area()
		}
	}
	fmt.Printf("ans: %d\n", r)
}

type Garden struct {
	// Map from plant type -> Region.
	regionsMap map[string][]*Region
}

type Cord struct {
	row int
	col int
}

func NewRegion(plantType string) *Region {
	return &Region{
		plantType:    plantType,
		cords:        NewSet[Cord](),
		rows:         NewSet[int](),
		cols:         NewSet[int](),
		innerRegions: []*Region{},
	}
}

type Region struct {
	plantType    string    // Letter.
	cords        Set[Cord] // (row, col) pairings.
	rows         Set[int]  // Row 0-indexes.
	cols         Set[int]  // Col 0-indexes.
	innerRegions []*Region // Regions that are enclosed inside this one.
}

func (r *Region) Perimeter(input [][]string) int {
	perimeter := 0
	for _, cord := range r.cords.All() {
		neighbors := GetAllNeighbors(cord, input)
		for _, neighbor := range neighbors.All() {
			if !IsNeighborWithinBounds(neighbor, input) {
				perimeter++
				continue
			}
			if input[neighbor.row][neighbor.col] != r.plantType {
				perimeter++
			}
		}
	}
	return perimeter
}

func (r *Region) Area() int {
	return len(r.cords)
}

func GetGarden(input [][]string) *Garden {
	g := &Garden{regionsMap: map[string][]*Region{}}
	visited := NewSet[Cord]()
	for r := range len(input) {
		for c := range len(input[0]) {
			if visited.Contains(Cord{row: r, col: c}) {
				continue
			}
			newRegion := Bfs(input, Cord{row: r, col: c}, visited)
			// Appending to nil slice is safe.
			g.regionsMap[input[r][c]] = append(g.regionsMap[input[r][c]], newRegion)
		}
	}
	// Handling of inner regions.
	allRegions := []*Region{}
	for _, regions := range g.regionsMap {
		allRegions = append(allRegions, regions...)
	}
	for _, region := range allRegions {
		if parentRegion, ok := MaybeGetParentRegion(region, allRegions, input); ok {
			fmt.Printf("%s has parent region %s\n", region.plantType, parentRegion.plantType)
			parentRegion.innerRegions = append(parentRegion.innerRegions, region)
		}
	}
	return g
}

func MaybeGetParentRegion(region *Region, allRegions []*Region, input [][]string) (*Region /*parent*/, bool /*exists*/) {
	diffNeighbors := NewSet[string]()
	var anyCord *Cord
	for _, cord := range region.cords.All() {
		foundForCord := false
		neighborCords := GetValidNeighbors(cord, input).All()
		for _, neighbor := range neighborCords {
			if input[neighbor.row][neighbor.col] == region.plantType {
				continue
			}
			foundForCord = true
			diffNeighbors.Add(input[neighbor.row][neighbor.col])
			anyCord = &neighbor
		}
		if !foundForCord {
			return nil, false
		}
	}
	if diffNeighbors.Size() != 1 {
		return nil, false
	}
	// Now find the region that contains `anyCord`.
	for _, region := range allRegions {
		if region.cords.Contains(*anyCord) {
			return region, true
		}
	}
	panic("expected to find a matching region")
}

func GetAllNeighbors(c Cord, input [][]string) Set[Cord] {
	r := NewSet[Cord]()
	for _, dir := range cDirs {
		r.Add(Cord{row: c.row + dir.rDiff, col: c.col + dir.cDiff})
	}
	return r
}

func IsNeighborWithinBounds(c Cord, input [][]string) bool {
	return !(c.row < 0 || c.row >= len(input) || c.col < 0 || c.col >= len(input[0]))
}

func GetValidNeighbors(c Cord, input [][]string) Set[Cord] {
	r := NewSet[Cord]()
	for _, dir := range cDirs {
		nr := c.row + dir.rDiff
		nc := c.col + dir.cDiff
		cord := Cord{row: nr, col: nc}
		if IsNeighborWithinBounds(cord, input) {
			r.Add(cord)
		}
	}
	return r
}

func Bfs(input [][]string /*const*/, start Cord, visited Set[Cord] /*in&out*/) *Region {
	region := NewRegion(input[start.row][start.col])
	q := []Cord{start}
	visited.Add(start)
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		region.cords.Add(top)
		region.rows.Add(top.row)
		region.cols.Add(top.col)

		neighbors := GetNeighbors(input, input[start.row][start.col], top, visited).All()
		for _, neighbor := range neighbors {
			q = append(q, neighbor)
			visited.Add(neighbor)
		}
	}
	return region
}

func GetNeighbors(input [][]string, plantType string, curr Cord, visited Set[Cord]) Set[Cord] {
	r := NewSet[Cord]()
	for _, dir := range cDirs {
		nr := curr.row + dir.rDiff
		nc := curr.col + dir.cDiff
		if nr < 0 || nr >= len(input) || nc < 0 || nc >= len(input[0]) {
			continue
		}
		if input[nr][nc] != plantType {
			continue
		}
		cord := Cord{row: nr, col: nc}
		if visited.Contains(cord) {
			continue
		}
		r.Add(cord)
	}
	return r
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) All() []T {
	r := []T{}
	for k := range s {
		r = append(r, k)
	}
	return r
}

func MustGetInput() [][]string {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := [][]string{}
	for _, line := range strings.Split(string(fileBytes), "\n") {
		curr := []string{}
		for _, ch := range strings.Split(line, "") {
			curr = append(curr, ch)
		}
		r = append(r, curr)
	}
	return r[:len(r)-1] // Ignore last newline.
}
