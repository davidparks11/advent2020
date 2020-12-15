package advent

import (
	"strconv"
	"strings"
)

const minUsageIndex = 0
const maxUsageIndex = 2
const characterIndex = 4
const startPassWordIndex = 7

var _ Problem = &PasswordPhilosophy{}

type PasswordPhilosophy struct {
	dailyProblem
}

func (p *PasswordPhilosophy) Solve() {
	p.day = 2
	p.name = "PasswordPhilosophy"
	input := p.GetInputLines()
	results := make([]string, 2)
	results[0] = strconv.Itoa(p.validPassWordCount(input))

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
		maxUsage, _ = strconv.ParseInt(inputLine[hyphenIndex+1:colonIndex - 2], 10, 0)
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

