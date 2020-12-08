package advent

import (
	"strconv"
)

var _ Problem = &ReportRepair{}

type ReportRepair struct {
	dailyProblem
}

func (r *ReportRepair) Solve() {
	r.day = 1
	r.name = "Report Repair"
	input := parseInput(r.GetInputFile())
	if len(input) == 0 {
		return
	}
	result := fixExpenseReport(input)
	r.WriteResult(strconv.Itoa(result))
}

func parseInput(inputLines []string) []int {
	input := make([]int, len(inputLines))
	for _, line := range inputLines {
		intValue, err := strconv.Atoi(line)
		if err != nil {
			return []int{}
		}
		input = append(input, intValue)
	}
	return input
}

func fixExpenseReport(input []int) int {
	expenses := make(map[int]bool, len(input))
	for _, v := range input {
		expenses[v] = true
	}

	for expense, _ := range expenses {
		complement := 2020 - expense
		if expenses[complement] {
			return complement * expense
		}
	}

	return 0
}


