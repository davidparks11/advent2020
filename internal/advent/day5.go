package advent

import (
	"fmt"
	"strconv")


var _ Problem = &BinaryBoarding{}

type BinaryBoarding struct {
	dailyProblem
}

func (b *BinaryBoarding) Solve() {
	b.day = 5
	b.name = "Binary Boarding"
	input := b.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(b.getHighestSeatID(input)))
	b.WriteResult(results)
}

const (
	numberOfColumns  = 8
	rowSeatDataIndex = 7
)

func (b *BinaryBoarding) getHighestSeatID(boardingPasses []string) int {
	highestID := 0
	currentID := 0
	min := 0
	max := 127
	for _, pass := range boardingPasses {
		for _, char := range pass[:rowSeatDataIndex] {
			if char == 'B' {
				min += (max-min + 1)/2
			} else {
				max -= (max-min + 1)/2
			}
		}
		currentID = min * numberOfColumns
		min = 0
		max = 7
		for _, char := range pass[rowSeatDataIndex:] {
			if char == 'R' {
				min += (max-min + 1)/2
			} else {
				max -= (max-min + 1)/2
			}
		}
		currentID += min
		if currentID > highestID {
			highestID = currentID
		}
		min = 0
		max = 127
	}
	return highestID
}
