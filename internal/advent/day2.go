package advent

import (
	"strconv"
	"strings"
)

var _ Problem = &PasswordPhilosophy{}

type PasswordPhilosophy struct {
	dailyProblem
}

func (p *PasswordPhilosophy) Solve() {
	p.day = 2
	p.name = "Password Philosophy"
	input := p.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(p.validPassWordCount(input)))
	results = append(results, strconv.Itoa(p.validPassWordCountPart2(input)))
	p.WriteResult(results)

}

func (p *PasswordPhilosophy) validPassWordCount(input []string) int {
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

func (p *PasswordPhilosophy) validPassWordCountPart2(input []string) int {
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
