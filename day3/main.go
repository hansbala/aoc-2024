package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// "xxxxmul(1
//      ^

func main() {
	input := getInput()
	totalSum := 0
	for i := 0; i < len(input); i++ {
		if input[i] != 'm' {
			continue
		}
		if input[i:i+4] != "mul(" {
			continue
		}
		i += 4
		firstNumber, ok := parseNextNumber(input[i:])
		if !ok {
			continue
		}
		i += len(strconv.Itoa(firstNumber)) + 1
		secondNumber, ok := parseNextNumber(input[i:])
		if !ok {
			continue
		}
		i += len(strconv.Itoa(secondNumber))
		totalSum += firstNumber * secondNumber
	}
	fmt.Printf("Ans: %d\n", totalSum)
}

func parseNextNumber(input string) (int, bool) {
	commaPosition := strings.Index(input, ",")
	closeBracketPosition := strings.Index(input, ")")
	if commaPosition == -1 && closeBracketPosition == -1 {
		return -1, false
	}
	// parse number till comma
	if commaPosition != -1 {
		number, err := strconv.Atoi(input[:commaPosition])
		if err == nil && number < 1000 {
			return number, true
		}
	}
	// parse number till closing bracket
	if closeBracketPosition != -1 {
		number, err := strconv.Atoi(input[:closeBracketPosition])
		if err == nil && number < 1000 {
			return number, true
		}
	}
	return -1, false
}

func getInput() string {
	content, _ := os.ReadFile("input.txt")
	return string(content)
}
