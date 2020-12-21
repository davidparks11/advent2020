package advent

import (
	"strconv"
)

var _ Problem = &HandheldHalting{}

type HandheldHalting struct {
	dailyProblem
}

func (h *HandheldHalting) Solve() {
	h.day = 8
	h.name = "Handheld Halting"
	input := h.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(h.infiniteLoopFinder(input)))
	h.WriteResult(results)
}

func (h *HandheldHalting) infiniteLoopFinder(programLines []string) int {
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

