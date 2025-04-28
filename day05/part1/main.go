package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	prevPages, updates := MustGetInput()
	// fmt.Printf("len graph: %d\n", len(backwardGraph))
	// fmt.Printf("len page updates: %d\n", len(pageUpdates))
	validUpdateIndices := []int{}
	for idx, pageUpdate := range updates {
		isValid := true
		for i := 0; i < len(pageUpdate); i++ {
			prevPageNumbers := prevPages[pageUpdate[i]]
			for j := i + 1; j < len(pageUpdate); j++ {
				if slices.Contains(prevPageNumbers, pageUpdate[j]) {
					isValid = false
				}
			}
		}
		if isValid {
			validUpdateIndices = append(validUpdateIndices, idx)
		}
	}
	// fmt.Printf("valid page idx(s): %v\n", validPageUpdateIdx)
	ans := 0
	for _, validUpdateIndex := range validUpdateIndices {
		pageUpdate := updates[validUpdateIndex]
		if len(pageUpdate)%2 == 0 {
			panic("page update has even so no middle element")
		}
		ans += pageUpdate[len(pageUpdate)/2]
	}
	fmt.Printf("ans -> %d\n", ans)
}

func MustGetInput() (map[int][]int /*prevPages*/, [][]int /*pageUpdates*/) {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	processGraph := true
	prevPages := map[int][]int{}
	pageUpdates := [][]int{}
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			processGraph = false
			continue
		}
		if processGraph {
			nodes := strings.Split(line, "|")
			if len(nodes) != 2 {
				panic("expected nodes to be of len 2")
			}
			start, err := strconv.Atoi(nodes[0])
			if err != nil {
				panic(fmt.Errorf("could not convert start to integer: %w", err))
			}
			end, err := strconv.Atoi(nodes[1])
			if err != nil {
				panic(fmt.Errorf("could not convert end to integer: %w", err))
			}
			prevPages[end] = append(prevPages[end], start)
			continue
		}
		// Process page updates.
		pageUpdates = append(pageUpdates, []int{})
		pageNumbers := strings.Split(line, ",")
		for _, strPageNum := range pageNumbers {
			pageNum, err := strconv.Atoi(strPageNum)
			if err != nil {
				panic(fmt.Errorf("failed to convert page num to integer: %w", err))
			}
			pageUpdates[len(pageUpdates)-1] = append(pageUpdates[len(pageUpdates)-1], pageNum)
		}
	}
	return prevPages, pageUpdates
}
