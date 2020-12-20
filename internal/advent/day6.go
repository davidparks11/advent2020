package advent

import (
	"strconv"
)

var _ Problem = &CustomCustoms{}

type CustomCustoms struct {
	dailyProblem
}

func (c *CustomCustoms) Solve() {
	c.day = 6
	c.name = "Custom Customs"
	input := c.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(c.getCustomsYesCount(input)))
	results = append(results, strconv.Itoa(c.getCustomsGroupYesCount(input)))
	c.WriteResult(results)
}

func (c *CustomCustoms) getCustomsYesCount(passengerAnswers []string) int {
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

func (c *CustomCustoms) getCustomsGroupYesCount(passengerAnswers []string) int {
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
