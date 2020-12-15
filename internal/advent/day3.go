package advent

import "strconv"

var _ Problem = &TobogganTrajectory{}

type TobogganTrajectory struct {
	dailyProblem
}

type slope struct {
	x int
	y int
}

func (t *TobogganTrajectory) Solve() {
	t.day = 3
	t.name = "Toboggan Trajectory"
	input := t.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(t.treeCountInPath(1, 3, input)))

	part2Slopes := []slope{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	part2Results := 1
	for _, slope := range part2Slopes {
		part2Results *= t.treeCountInPath(slope.y, slope.x, input)
	}
	results = append(results, strconv.Itoa(part2Results))
	t.WriteResult(results)
}

func (t *TobogganTrajectory) treeCountInPath(slopeY, slopeX int, treeLines []string) int {
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
