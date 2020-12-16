package advent

import (
	"strconv"
)

const fieldLength = 3

var _ Problem = &PassportProcessing{}

type PassportProcessing struct {
	dailyProblem
}

func (p *PassportProcessing) Solve() {
	p.day = 4
	p.name = "Passport Processing"
	input := p.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(p.getValidPassportCount(input)))

	p.WriteResult(results)
}

func (p *PassportProcessing) getValidPassportCount(passportData []string) int {
	validPassportCount := 0
	requireFields := map[string]bool{
		"ecl": false,
		"pid": false,
		"eyr": false,
		"hcl": false,
		"byr": false,
		"iyr": false,
		"hgt": false,
	}

	valid := true
	for i, dataLine := range passportData {
		if dataLine != "" {
			//loops through lines for :, set previous field to true
			for j, char := range dataLine {
				if char == ':' {
					if dataLine[j-fieldLength:j] != "cid" {
						requireFields[dataLine[j-fieldLength:j]] = true
					}
				}
			}
		}
		if dataLine == "" || i == len(passportData)-1 {
			//validate passports
			for field, fieldPresent := range requireFields {
				if !fieldPresent {
					valid = false
				}
				//reset required field
				requireFields[field] = false
			}
			if valid {
				validPassportCount++
			}
			valid = true
		}
	}
	return validPassportCount
}
