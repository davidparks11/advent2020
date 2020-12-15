package advent

import "strconv"

var _ Problem = &TobogganTrajectory{}

type TobogganTrajectory struct {
	dailyProblem
}

func (t *TobogganTrajectory) Solve() {
	t.day = 3
	t.name = "Toboggan Trajectory"
	input := t.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(t.treeCountInPath(input)))

	t.WriteResult(results)
}

func (t *TobogganTrajectory) treeCountInPath(trees []string) int {
	treeCollisions := 0
	const (
		tree = uint8('#')
	)
	treePaternLength := len(trees[0])
	horizontalTravel := 3
	horizontalIndex := 0
	for _, treeLine := range trees {
		if treeLine[horizontalIndex] == tree {
			treeCollisions++
		}
		horizontalIndex = (horizontalIndex+horizontalTravel)%treePaternLength
	}
	return treeCollisions
}