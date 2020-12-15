package advent

import (
	"fmt"
	"strconv"
)

const xmasPreambleLength = 25

var _ Problem = &EncodingError{}

type EncodingError struct {
	dailyProblem
}

func (e *EncodingError) Solve() {
	e.day = 9
	e.name = "Encoding Error"
	input := IntsFromStrings(e.GetInputLines())
	results := make([]string, 2)
	result1 := findXMASDataWeakness(input)
	results[0] = strconv.Itoa(result1)
	results[1] = strconv.Itoa(findXMASDataWeaknessPart2(result1, input))
	e.WriteResult(results)
}

//findXMASDataWeakness checks for invalid 'XMAS' encoding. The first 25 integers are called the preamble and server to
//validate the next (26th) number. Every subsequent number checks the previous 25 for validity. A number is considered
//valid if two of the previous numbers sum to the current number.
func findXMASDataWeakness(input []int) int {
	validators := make(map[int]bool, xmasPreambleLength)
	for i := 0; i < xmasPreambleLength; i++ {
		validators[input[i]] = true
	}
	target := 0
	complement := 0
	for i := xmasPreambleLength; i < len(input); i++ {
		target = input[i]
		isValid := false
		for validator := range validators {
			complement = target - validator
			if complement != validator && validators[complement] {
				isValid = true
				break
			}
		}
		if !isValid {
			return target
		}
		delete(validators, input[i-xmasPreambleLength])
		validators[target] = true
	}
	return -1
}

func findXMASDataWeaknessPart2(target int, input []int) int {
	fmt.Println("finding sum for target: ", target)
	inputLength := len(input)
	if inputLength < 2 {
		return -1
	}
	startSumIndex := 0
	endSumIndex := 1

	//start running sum with first two entries
	runningSum := input[startSumIndex] + input[endSumIndex]

	//break if runningSum < target && endSumIndex = len(input) - 1
	for runningSum > target || endSumIndex != inputLength - 1 {
		if runningSum == target {

			smallest := input[startSumIndex]
			largest := input[startSumIndex]
			for i := startSumIndex + 1; i < endSumIndex; i++ {
				if input[i] < smallest {
					smallest = input[i]
				}
				if input[i] > largest {
					largest = input[i]
				}
			}
			return smallest + largest
		}
		//check if start and end are next to each other, increment both
		if runningSum > target && startSumIndex == endSumIndex - 1 {
			runningSum -= input[startSumIndex]
			startSumIndex++
			endSumIndex++
			runningSum += input[endSumIndex]
		} else if runningSum > target {
			runningSum -= input[startSumIndex]
			startSumIndex++
		} else {
			endSumIndex++
			runningSum += input[endSumIndex]
		}
	}
	return -1
}
