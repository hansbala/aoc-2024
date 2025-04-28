package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cMulPrefix  = "mul("
	cDoPrefix   = "do()"
	cDontPrefix = "don't()"
)

func main() {
	fileBytes, err := os.ReadFile("inp.txt")
	if err != nil {
		panic(err)
	}
	s := string(fileBytes)
	ans := 0
	mulEnabled := true
	for {
		newS, mul, ok, newMulEnabled := parseNext(s, mulEnabled)
		if len(newS) == 0 {
			break
		}
		s = newS
		mulEnabled = newMulEnabled
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
func parseNext(s string, mulEnabled bool) (
	string, /*new string*/
	*MulOp, /*parsed op if any*/
	bool, /*parsed*/
	bool, /*new mulEnabled*/
) {
	if mulEnabled && len(s) >= len(cDontPrefix) && s[:len(cDontPrefix)] == cDontPrefix {
		return s[len(cDontPrefix)+1:], nil, false, false
	}
	if !mulEnabled && len(s) >= len(cDoPrefix) && s[:len(cDoPrefix)] == cDoPrefix {
		return s[len(cDoPrefix)+1:], nil, false, true
	}
	if len(s) < len(cMulPrefix)+1 {
		return "", nil, false, mulEnabled
	}
	if s[:len(cMulPrefix)] != cMulPrefix {
		return s[1:], nil, false, mulEnabled
	}
	s = s[len(cMulPrefix):]
	commaIdx := strings.Index(s, ",")
	if commaIdx == -1 {
		return s, nil, false, mulEnabled
	}
	num1, err := strconv.Atoi(s[:commaIdx])
	if err != nil {
		return s, nil, false, mulEnabled
	}
	s = s[commaIdx+1:]
	bracketIndex := strings.Index(s, ")")
	if bracketIndex == -1 {
		return s, nil, false, mulEnabled
	}
	num2, err := strconv.Atoi(s[:bracketIndex])
	if err != nil {
		return s, nil, false, mulEnabled
	}
	return s[bracketIndex+1:], &MulOp{num1: num1, num2: num2}, mulEnabled /*parsed*/, mulEnabled
}
