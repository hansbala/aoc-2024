package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := MustGetInput()
	// fmt.Printf("%v\n", stones)
	blinks := 75
	for range blinks {
		stones = Blink(stones)
	}
	fmt.Printf("ans: %d\n", len(stones))
}

func Blink(stones []int) []int {
	r := []int{}
	for _, stone := range stones {
		// Rule 1: 0 replaced by 1.
		if stone == 0 {
			r = append(r, 1)
			continue
		}
		// Rule 2: even number of digits.
		strNumber := strconv.Itoa(stone)
		if len(strNumber)%2 == 0 {
			firstHalfStr := strNumber[:len(strNumber)/2]
			secondHalfStr := strNumber[len(strNumber)/2:]
			firstHalf, err := strconv.Atoi(firstHalfStr)
			if err != nil {
				panic(err)
			}
			secondHalf, err := strconv.Atoi(secondHalfStr)
			if err != nil {
				panic(err)
			}
			r = append(r, firstHalf)
			r = append(r, secondHalf)
			continue
		}
		// Rule 3: Multiply by 2024.
		r = append(r, stone*2024)
	}
	return r
}

func MustGetInput() []int {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := []int{}
	strNumbers := strings.Split(string(fileContent), " ")
	for _, strNumber := range strNumbers {
		strNumber = strings.Trim(strNumber, "\n")
		number, err := strconv.Atoi(strNumber)
		if err != nil {
			panic(err)
		}
		r = append(r, number)
	}
	return r
}
