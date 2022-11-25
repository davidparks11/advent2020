package day1

import (
	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &reportRepair{}

type reportRepair struct {
	DailyProblem
}

func New() Problem {
	return &reportRepair{
		DailyProblem{Day: 1},
	}
}

func (r *reportRepair) Solve() interface{} {
	r.Day = 1
	input := IntsFromStrings(r.GetInputLines())
	var results []int
	results = append(results, fixExpenseReport(input, 2020))
	results = append(results, fixExpenseReportPart2(input))
	return results
}

func fixExpenseReport(input []int, target int) int {
	expenses := make(map[int]bool, len(input))
	for _, v := range input {
		expenses[v] = true
	}

	for expense := range expenses {
		complement := target - expense
		if expenses[complement] {
			return complement * expense
		}
	}
	return 0
}

func fixExpenseReportPart2(input []int) int {
	expenses := make(map[int]bool, len(input))
	for _, v := range input {
		expenses[v] = true
	}

	for expense := range expenses {
		complement := 2020 - expense
		if fixedReport := fixExpenseReport(input, complement); fixedReport != 0 {
			return fixedReport * expense
		}
	}
	return 0
}
