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
