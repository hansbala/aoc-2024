package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type KeyWrapper struct {
	goal  int64
	index int
}

func main() {
	cases := MustGetInput()
	ans := int64(0)
	for kw, perm := range cases {
		goal := kw.goal
		if IsValid(perm, goal) {
			ans += goal
		}
	}
	fmt.Printf("ans -> %d\n", ans)
}

func IsValid(perm []int64, goal int64) bool {
	if len(perm) == 0 {
		panic("received 0 in perm which should not happen")
	}
	if len(perm) == 1 {
		if perm[0] == goal {
			return true
		}
		return false
	}
	first, second := perm[0], perm[1]
	// Case 1: Apply addition.
	if IsValid(append([]int64{first + second}, perm[2:]...), goal) {
		return true
	}
	// Case 2: Apply multiplication.
	if IsValid(append([]int64{first * second}, perm[2:]...), goal) {
		return true
	}
	// Case 3: Apply concatenation.
	if IsValid(append([]int64{concatenate(first, second)}, perm[2:]...), goal) {
		return true
	}
	// Doing either is invalid so return false.
	return false
}

func concatenate(first int64, second int64) int64 {
	r, err := strconv.Atoi(strconv.Itoa(int(first)) + strconv.Itoa(int(second)))
	if err != nil {
		panic(err)
	}
	return int64(r)
}

func MustGetInput() map[KeyWrapper][]int64 {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := map[KeyWrapper][]int64{}
	for lineIdx, line := range strings.Split(string(contents), "\n") {
		testAndNums := strings.Split(line, ": ")
		if len(testAndNums) != 2 {
			panic("expected two sections")
		}
		goal, err := strconv.Atoi(testAndNums[0])
		if err != nil {
			panic(fmt.Errorf("failed to parse goal: %w", err))
		}
		for _, numStr := range strings.Split(testAndNums[1], " ") {
			currNum, err := strconv.Atoi(numStr)
			if err != nil {
				panic(fmt.Errorf("failed to parse number: %w", err))
			}
			keyWrapper := KeyWrapper{goal: int64(goal), index: lineIdx}
			r[keyWrapper] = append(r[keyWrapper], int64(currNum))
		}
	}
	return r
}
