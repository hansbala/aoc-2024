package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getInput()
	numSafeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			numSafeReports += 1
			continue
		}
		if reportSafeByRemovingLevel(report) {
			numSafeReports += 1
		}
	}
	fmt.Println(numSafeReports)
}

func reportSafeByRemovingLevel(report []int) bool {
	// Try to remove each index one-by-one and check if it is safe.
	for idx, _ := range report {
		newReport := constructNewReport(report, idx)
		if isReportSafe(newReport) {
			return true
		}
	}
	return false
}

func constructNewReport(report []int, idxToRemove int) []int {
	newReport := []int{}
	for i, val := range report {
		if i == idxToRemove {
			continue
		}
		newReport = append(newReport, val)
	}
	return newReport
}

func isReportSafe(report []int) bool {
	if !isStrictlyIncreasing(report) && !isStrictlyDecreasing(report) {
		return false
	}
	for idx := 0; idx < len(report)-1; idx++ {
		if !isDiffSafe(report[idx], report[idx+1]) {
			return false
		}
	}
	return true
}

func isStrictlyIncreasing(report []int) bool {
	for idx := len(report) - 1; idx >= 1; idx-- {
		if report[idx] < report[idx-1] {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(report []int) bool {
	for idx := len(report) - 1; idx >= 1; idx-- {
		if report[idx] > report[idx-1] {
			return false
		}
	}
	return true
}

func isDiffSafe(num1, num2 int) bool {
	diff := int(math.Abs(float64(num1) - float64(num2)))
	return diff >= 1 && diff <= 3
}

func getInput() [][]int {
	r := [][]int{}
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		reportStr := strings.Split(line, " ")
		report := []int{}
		for _, levelStr := range reportStr {
			level, _ := strconv.Atoi(levelStr)
			report = append(report, level)
		}
		r = append(r, report)
	}
	return r
}
