package main

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.GetLines("puzzle")

	safeReports := 0
	for _, report := range lines {
		levels := strings.Split(report, " ")

		if ok, _ := IsValid(levels); ok {
			safeReports++
		}
	}

	fmt.Printf("First part solution is %d\n", safeReports)

	safeReports = 0
	for _, report := range lines {
		levels := strings.Split(report, " ")

		ok, errIndex := IsValid(levels)
		if ok {
			safeReports++
			continue
		}

		for i := errIndex; i >= 0; i-- {
			modified := RemoveElement(levels, i)
			if ok, _ = IsValid(modified); ok {
				safeReports++
				break
			}
		}
	}

	fmt.Printf("Second part solution is %d\n", safeReports)
}

func IsValid(levels []string) (bool, int) {
	lastDiff := 0
	isValid := true
	errorIndex := -1
	for i := range levels {
		if i == 0 {
			continue
		}

		p, _ := strconv.Atoi(levels[i-1])
		c, _ := strconv.Atoi(levels[i])

		diff := p - c

		if (diff < 0 && lastDiff > 0) || (diff > 0 && lastDiff < 0) || diff == 0 || diff > 3 || diff < -3 {
			isValid = false
			errorIndex = i
			break
		}

		lastDiff = diff
	}
	return isValid, errorIndex
}

func RemoveElement(levels []string, ind int) []string {
	if ind == 0 {
		return levels[1:]
	} else if ind == len(levels)-1 {
		return levels[0 : len(levels)-1]
	} else {
		left := levels[0:ind]
		right := levels[ind+1:]
		joints := make([]string, 0, len(levels))
		joints = append(joints, left...)
		joints = append(joints, right...)
		return joints
	}
}
