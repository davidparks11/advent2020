package day6

import (
	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &customCustoms{}

type customCustoms struct {
	DailyProblem
}

func New() Problem {
	return &customCustoms{
		DailyProblem{Day: 6},
	}
}

func (c *customCustoms) Solve() interface{} {
	c.Day = 6
	input := c.GetInputLines()
	var results []int
	results = append(results, getCustomsYesCount(input))
	results = append(results, getCustomsGroupYesCount(input))
	return results
}

func getCustomsYesCount(passengerAnswers []string) int {
	answerCount := len(passengerAnswers)
	yesCount := 0
	yesMap := make(map[rune]bool, 26)
	for i, answerLine := range passengerAnswers {

		if answerLine != "" {
			for _, answer := range answerLine {
				if answer != ' ' {
					yesMap[answer] = true
				}
			}
		}
		if answerLine == "" || i == answerCount-1 {
			yesCount += len(yesMap)
			yesMap = make(map[rune]bool, 26)
		}
	}
	return yesCount
}

func getCustomsGroupYesCount(passengerAnswers []string) int {
	answerCount := len(passengerAnswers)
	yesCount := 0
	groupTotal := 0
	yesMap := make(map[rune]int, 26)
	for i, answerLine := range passengerAnswers {

		if answerLine != "" {
			for _, answer := range answerLine {
				if answer != ' ' {
					yesMap[answer]++
				}
			}
			groupTotal++
		}
		if answerLine == "" || i == answerCount-1 {
			for _, answerCount := range yesMap {
				if answerCount == groupTotal {
					yesCount++
				}
			}
			yesMap = make(map[rune]int, 26)
			groupTotal = 0
		}
	}
	return yesCount
}
