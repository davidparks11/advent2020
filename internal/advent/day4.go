package advent

import (
	"regexp"
	"strconv"
	"strings"
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
	results = append(results, strconv.Itoa(p.getValidPassportCountPart2(input)))
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

func (p *PassportProcessing) getValidPassportCountPart2(passportData []string) int {
	validPassportCount := 0
	valid := true
	nextSpace := 0

	requireFields := map[string]string{
		"ecl": "",
		"pid": "",
		"eyr": "",
		"hcl": "",
		"byr": "",
		"iyr": "",
		"hgt": "",
	}

	for i, dataLine := range passportData {
		if dataLine != "" {
			//loops through lines for :, set previous field to true
			for j, char := range dataLine {
				if char == ':' {
					if dataLine[j-fieldLength:j] != "cid" {
						nextSpace = strings.Index(dataLine[j:], " ")
						if nextSpace == -1 {
							nextSpace = len(dataLine)
						} else {
							nextSpace += j
						}
						requireFields[dataLine[j-fieldLength:j]] = dataLine[j+1 : nextSpace]
					}
				}
			}
		}
		if dataLine == "" || i == len(passportData)-1 {
			//validate passports
			for field, fieldValue := range requireFields {
				if fieldValue == "" || !p.isValidField(field, fieldValue) {
					valid = false
				}
				//reset required field
				requireFields[field] = ""
			}
			if valid {
				validPassportCount++
			}
			valid = true
		}
	}
	return validPassportCount
}

var (
	validBYR = func(stringVal string) bool {
		val, _ := strconv.Atoi(stringVal)
		return 1920 <= val && val <= 2002
	}
	validIYR = func(stringVal string) bool {
		val, _ := strconv.Atoi(stringVal)
		return 2010 <= val && val <= 2020
	}
	validEYR = func(stringVal string) bool {
		val, _ := strconv.Atoi(stringVal)
		return 2020 <= val && val <= 2030
	}
	validHGT = func(stringVal string) bool {
		valLength := len(stringVal)
		unit := stringVal[valLength-2 : valLength]
		val, _ := strconv.Atoi(stringVal[:valLength-2])
		if unit == "cm" {
			return 150 <= val && val <= 193
		} else if unit == "in" {
			return 59 <= val && val <= 76
		}
		return false
	}
	validHCL = regexp.MustCompile(`(^#[a-f0-9]{6}$)`)
	validECL = regexp.MustCompile(`^((amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth))$`)
	validPID = regexp.MustCompile(`^([0-9]{9})$`)
)

func (p *PassportProcessing) isValidField(fieldName string, fieldValue string) bool {
	switch fieldName {
	case "byr":
		return validBYR(fieldValue)
	case "eyr":
		return validEYR(fieldValue)
	case "iyr":
		return validIYR(fieldValue)
	case "hgt":
		return validHGT(fieldValue)
	case "hcl":
		return validHCL.MatchString(fieldValue)
	case "ecl":
		return validECL.MatchString(fieldValue)
	case "pid":
		return validPID.MatchString(fieldValue)
	default:
		return false
	}

}
