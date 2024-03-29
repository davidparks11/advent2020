package day4

import (
	"regexp"
	"strconv"
	"strings"

	. "github.com/davidparks11/advent2020/internal/problem"
)

const fieldLength = 3

var _ Problem = &passportProcessing{}

type passportProcessing struct {
	DailyProblem
}

func New() Problem {
	return &passportProcessing{
		DailyProblem{Day: 4},
	}
}

func (p *passportProcessing) Solve() interface{} {
	p.Day = 4
	input := p.GetInputLines()
	var results []int
	results = append(results, getValidPassportCount(input))
	results = append(results, getValidPassportCountPart2(input))
	return results
}

func getValidPassportCount(passportData []string) int {
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

func getValidPassportCountPart2(passportData []string) int {
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
				if fieldValue == "" || !isValidField(field, fieldValue) {
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

func isValidField(fieldName string, fieldValue string) bool {
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
