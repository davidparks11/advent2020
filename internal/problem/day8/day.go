package day8

import (
	"strconv"

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
	var results []string
	results = append(results, strconv.Itoa(h.infiniteLoopFinder(input)))
	return results
}

func (h *handheldHalting) infiniteLoopFinder(programLines []string) int {
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
