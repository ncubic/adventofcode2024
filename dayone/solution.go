package main

import (
	"aoc2024/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := common.GetLines("puzzle")

	leftValues := make([]int, len(lines))
	rightValues := make([]int, len(lines))

	for i, value := range lines {
		parts := strings.Split(value, "   ")
		left, errL := strconv.Atoi(parts[0])
		right, errR := strconv.Atoi(parts[1])

		if errL != nil {
			log.Fatal(errL)
		}

		if errR != nil {
			log.Fatal(errR)
		}

		leftValues[i] = left
		rightValues[i] = right
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	total := 0
	for i := range lines {
		diff := leftValues[i] - rightValues[i]
		if diff < 0 {
			diff = diff * -1
		}
		total += diff
	}

	fmt.Printf("Part 1 solution is %d\n", total)

	rightValuesOccurence := make(map[int]int)

	for _, val := range rightValues {
		_, ok := rightValuesOccurence[val]
		if !ok {
			rightValuesOccurence[val] = 1
		} else {
			rightValuesOccurence[val] += 1
		}
	}

	similarity := 0
	for _, val := range leftValues {
		occurences, ok := rightValuesOccurence[val]
		if !ok {
			continue
		}
		similarity += (val * occurences)
	}

	fmt.Printf("Part 2 solution is %d\n", similarity)
}
