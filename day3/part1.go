package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cMulPrefix = "mul("
)

func main() {
	fileBytes, err := os.ReadFile("inp2.txt")
	if err != nil {
		panic(err)
	}
	s := string(fileBytes)
	ans := 0
	for {
		newS, mul, ok := parseNext(s)
		if len(newS) == 0 {
			break
		}
		s = newS
		if ok {
			ans += (mul.num1 * mul.num2)
		}
	}
	fmt.Printf("ans -> %d\n", ans)
}

type MulOp struct {
	num1 int
	num2 int
}

// xxxmul(123,456)xxx
//
//	^
func parseNext(s string) (string /*new string*/, *MulOp /*parsed op if any*/, bool /*parsed*/) {
	if len(s) < len(cMulPrefix)+1 {
		return "", nil, false
	}
	if s[:len(cMulPrefix)] != cMulPrefix {
		return s[1:], nil, false
	}
	s = s[len(cMulPrefix):]
	commaIdx := strings.Index(s, ",")
	if commaIdx == -1 {
		return s, nil, false
	}
	num1, err := strconv.Atoi(s[:commaIdx])
	if err != nil {
		return s, nil, false
	}
	s = s[commaIdx+1:]
	bracketIndex := strings.Index(s, ")")
	if bracketIndex == -1 {
		return s, nil, false
	}
	num2, err := strconv.Atoi(s[:bracketIndex])
	if err != nil {
		return s, nil, false
	}
	s = s[bracketIndex+1:]
	return s, &MulOp{num1: num1, num2: num2}, true
}
