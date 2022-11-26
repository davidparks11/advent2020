package day8

import (
	"strconv"

	"github.com/davidparks11/advent2020/internal/maps"
	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &handheldHalting{}

type handheldHalting struct {
	DailyProblem
}

func New() Problem {
	return &handheldHalting{
		DailyProblem{Day: 8},
	}
}

func (h *handheldHalting) Solve() interface{} {
	h.Day = 8
	input := h.GetInputLines()
	var results []int
	results = append(results, infiniteLoopFinder(input))
	results = append(results, fixBootCode(input))
	return results
}

func infiniteLoopFinder(programLines []string) int {
	accVal := 0
	runOrder := make(map[int]bool, len(programLines))
	currentInstruction := ""
	instructionValue := 0
	index := 0
	for {
		if runOrder[index] {
			break
		}
		runOrder[index] = true
		currentInstruction = programLines[index][:3]
		if currentInstruction == "acc" {
			instructionValue, _ = strconv.Atoi(programLines[index][4:])
			accVal += instructionValue
			index++
		} else if currentInstruction == "jmp" {
			instructionValue, _ = strconv.Atoi(programLines[index][4:])
			index += instructionValue
		} else {
			index++
		}
	}
	return accVal
}

func fixBootCode(programLines []string) int {
	return recurse(programLines, 0, 0, make(map[int]struct{}), false)
}

func recurse(instructions []string, index int, acc int, visited map[int]struct{}, switchedCommand bool) int {
	if _, found := visited[index]; found {
		return 0
	}
	visited[index] = struct{}{}
	
	if index == len(instructions) {
		return acc
	} else if index > len(instructions) {
		panic("something went wrong")
	}

	op, val := parseInstruction(instructions[index])

	if op == "acc" {
		return recurse(instructions, index+1, acc+val, maps.Copy(visited), switchedCommand)
	} else if op == "nop" {
		if switchedCommand {
			return recurse(instructions, index+1, acc, maps.Copy(visited), true)	
		}
		return recurse(instructions, index+1, acc, maps.Copy(visited), false) + recurse(instructions, index+val, acc, maps.Copy(visited), true)
	}

	if switchedCommand {
		return recurse(instructions, index+val, acc, maps.Copy(visited), true) 
	}
	return recurse(instructions, index+val, acc, maps.Copy(visited), false) + recurse(instructions, index+1, acc, maps.Copy(visited), true)
}	

func parseInstruction(instruction string) (string, int) {
	op := instruction[:3]
	instructionValue, err := strconv.Atoi(instruction[4:])
	if err != nil {
		panic(err.Error())
	}
	return op, instructionValue
}
