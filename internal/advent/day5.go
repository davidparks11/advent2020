package advent

import (
	"strconv"
)

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
	results = append(results, strconv.Itoa(b.getMissingSeat(input)))
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
				min += (max - min + 1) / 2
			} else {
				max -= (max - min + 1) / 2
			}
		}
		currentID = min * numberOfColumns
		min = 0
		max = 7
		for _, char := range pass[rowSeatDataIndex:] {
			if char == 'R' {
				min += (max - min + 1) / 2
			} else {
				max -= (max - min + 1) / 2
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

//getMissingSeat is part 2 of day 5's problem. Finds the gap in ids where ids[n] - ids[n+1] != 1
func (b *BinaryBoarding) getMissingSeat(boardingPasses []string) int {
	ids := make(map[int]bool, len(boardingPasses))
	currentID := 0
	highestID := 0
	lowestID := 0
	min := 0
	max := 127
	for i, pass := range boardingPasses {
		for _, char := range pass[:rowSeatDataIndex] {
			if char == 'B' {
				min += (max - min + 1) / 2
			} else {
				max -= (max - min + 1) / 2
			}
		}
		currentID = min * numberOfColumns
		min = 0
		max = 7
		for _, char := range pass[rowSeatDataIndex:] {
			if char == 'R' {
				min += (max - min + 1) / 2
			} else {
				max -= (max - min + 1) / 2
			}
		}
		currentID := currentID+min
		if currentID < lowestID  || i == 0 {
			lowestID = currentID
		} else if currentID > highestID {
			highestID = currentID
		}
		ids[currentID] = true
		min = 0
		max = 127
	}
	for i := lowestID; i < highestID; i++ {
		if ids[i] == false {
			return i
		}
	}
	return -1
}
