package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	list1, list2, err := getInput()
	if err != nil {
		panic(err)
	}

	simScore := 0
	counts := getFrequencyMap(list2)
	for _, val := range list1 {
		freq, ok := counts[val]
		if ok {
			simScore += val * freq
		}
	}
	fmt.Println(simScore)
}

func getFrequencyMap(list []int) map[int]int {
	sim := map[int]int{}
	for _, val := range list {
		if _, ok := sim[val]; !ok {
			sim[val] = 1
		} else {
			sim[val] += 1
		}
	}
	return sim
}

func getInput() ([]int, []int, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, nil, err
	}

	r1, r2 := []int{}, []int{}
	re := regexp.MustCompile(`\s+`)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		words := re.Split(line, -1)
		num1, _ := strconv.Atoi(words[0])
		num2, _ := strconv.Atoi(words[1])
		r1 = append(r1, num1)
		r2 = append(r2, num2)
	}
	return r1, r2, nil
}
