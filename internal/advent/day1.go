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
	results := make([]string, 2)
	results[0] = strconv.Itoa(fixExpenseReport(input, 2020))
	results[1] = strconv.Itoa(fixExpenseReportPart2(input))
	r.WriteResult(results)
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
	for _, v  := range input {
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

