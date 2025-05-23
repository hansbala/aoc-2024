package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	cCompLimit    = 25
	cProblemLimit = 75
)

func main() {
	stones := MustGetInput()
	// fmt.Printf("%v\n", stones)
	c := NewCache()
	fmt.Printf("ans: %d\n", RecBlink(stones, cProblemLimit, c))
}

func RecBlink(stones []int, blinkNum int, cash *Cache) int {
	if blinkNum == 0 {
		return len(stones)
	}
	if v, ok := cash.Get(stones, blinkNum); ok {
		return v
	}
	acc := 0
	newStones := Blink(stones)
	for _, stone := range newStones {
		acc += RecBlink([]int{stone}, blinkNum-1, cash)
	}
	cash.Set(stones, blinkNum, acc)
	return acc
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

func NewCache() *Cache {
	return &Cache{cache: map[SliceHash]int{}}
}

type SliceHash string
type Cache struct {
	cache map[SliceHash]int
}

func (c *Cache) Set(s []int, blinks int, v int) {
	c.cache[c.getHash(s, blinks)] = v
}

func (c *Cache) Get(s []int, blinks int) (int, bool) {
	v, ok := c.cache[c.getHash(s, blinks)]
	return v, ok
}

func (c *Cache) getHash(s []int, blinks int) SliceHash {
	cloned := slices.Clone(s)
	slices.Sort(cloned)
	r := SliceHash(fmt.Sprintf("%v, %d", cloned, blinks))
	// fmt.Println(r)
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
