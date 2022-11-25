package day3

import (
	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &tobogganTrajectory{}

type tobogganTrajectory struct {
	DailyProblem
}

func New() Problem {
	return &tobogganTrajectory{
		DailyProblem{Day: 3},
	}
}

type slope struct {
	x int
	y int
}

func (t *tobogganTrajectory) Solve() interface{} {
	t.Day = 3
	input := t.GetInputLines()
	var results []int
	results = append(results, treeCountInPath(1, 3, input))
	results = append(results, treeCountProductWholeSlope(input))

	return results
}

func treeCountProductWholeSlope(input []string) int {
	slopes := []slope{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	result := 1
	for _, slope := range slopes {
		result *= treeCountInPath(slope.y, slope.x, input)
	}

	return result
}

func treeCountInPath(slopeY, slopeX int, treeLines []string) int {
	treeCollisions := 0
	const (
		tree = uint8('#')
	)
	totalTreeLines := len(treeLines)
	treePatternLength := len(treeLines[0])
	horizontalIndex := 0
	for i := 0; i < totalTreeLines; i += slopeY {
		if treeLines[i][horizontalIndex] == tree {
			treeCollisions++
		}
		horizontalIndex = (horizontalIndex + slopeX) % treePatternLength
	}
	return treeCollisions
}
