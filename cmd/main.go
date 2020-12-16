package main

import (
	"github.com/davidparks11/advent2020/internal/advent"
)

func main() {
	problems := []advent.Problem {
		&advent.ReportRepair{},
		&advent.PasswordPhilosophy{},
		&advent.TobogganTrajectory{},
		&advent.PassportProcessing{},
		&advent.EncodingError{},
	}
	for _, problem := range problems {
		problem.Solve()
	}
}