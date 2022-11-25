package day2

import (
	"strconv"
	"strings"

	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &passwordPhilosophy{}

type passwordPhilosophy struct {
	DailyProblem
}

func New() Problem {
	return &passwordPhilosophy{
		DailyProblem{Day: 2},
	}
}

func (p *passwordPhilosophy) Solve() interface{} {
	p.Day = 2
	input := p.GetInputLines()
	var results []int
	results = append(results, validPassWordCount(input))
	results = append(results, validPassWordCountPart2(input))

	return results
}

func validPassWordCount(input []string) int {
	validPasswordCount := 0
	hyphenIndex := 0
	colonIndex := 1
	var minUsage int64
	var maxUsage int64
	var timesUsed int64
	var policyChar uint8

	//Errors are ignored assuming all input is correctly formatted
	for _, inputLine := range input {
		hyphenIndex = strings.Index(inputLine, "-")
		colonIndex = strings.Index(inputLine, ":")
		minUsage, _ = strconv.ParseInt(inputLine[:hyphenIndex], 10, 8)
		maxUsage, _ = strconv.ParseInt(inputLine[hyphenIndex+1:colonIndex-2], 10, 0)
		policyChar = inputLine[colonIndex-1]
		password := inputLine[colonIndex+2:]
		for i := 0; i < len(password); i++ {
			if password[i] == policyChar {
				timesUsed++
			}
		}
		if timesUsed >= minUsage && timesUsed <= maxUsage {
			validPasswordCount++
		}
		timesUsed = 0
	}
	return validPasswordCount
}

func validPassWordCountPart2(input []string) int {
	validPasswordCount := 0
	hyphenIndex := 0
	colonIndex := 1
	var firstPos int64
	var secondPos int64
	var policyChar uint8
	usedFirst := false
	usedSecond := false
	//Errors are ignored assuming all input is correctly formatted
	for _, inputLine := range input {
		hyphenIndex = strings.Index(inputLine, "-")
		colonIndex = strings.Index(inputLine, ":")
		firstPos, _ = strconv.ParseInt(inputLine[:hyphenIndex], 10, 8)
		secondPos, _ = strconv.ParseInt(inputLine[hyphenIndex+1:colonIndex-2], 10, 0)
		policyChar = inputLine[colonIndex-1]
		password := inputLine[colonIndex+2:]
		usedFirst = password[firstPos-1] == policyChar
		usedSecond = password[secondPos-1] == policyChar
		//password is not 0 indexed ;)
		if usedFirst != usedSecond {
			validPasswordCount++
		}
	}
	return validPasswordCount
}
