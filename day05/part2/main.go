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
	invalidUpdateIndices := []int{}
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
		if !isValid {
			invalidUpdateIndices = append(invalidUpdateIndices, idx)
		}
	}
	// fmt.Printf("valid page idx(s): %v\n", validPageUpdateIdx)
	ans := 0
	for _, invalidUpdateIndex := range invalidUpdateIndices {
		invalidUpdate := updates[invalidUpdateIndex]
		sortUpdate(invalidUpdate, prevPages)
		if len(invalidUpdate)%2 == 0 {
			panic("page update has even so no middle element")
		}
		ans += invalidUpdate[len(invalidUpdate)/2]
	}
	fmt.Printf("ans -> %d\n", ans)
}

func sortUpdate(update []int, prevPages map[int][]int) {
	slices.SortFunc(update, func(pageNum1 int, pageNum2 int) int {
		if slices.Contains(prevPages[pageNum1], pageNum2) {
			return -1
		}
		if slices.Contains(prevPages[pageNum2], pageNum1) {
			return 1
		}
		return 0
	})
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
