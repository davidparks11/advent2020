package advent

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davidparks11/advent2020/internal/problem"
	"github.com/davidparks11/advent2020/internal/problem/day1"
	"github.com/davidparks11/advent2020/internal/problem/day10"
	"github.com/davidparks11/advent2020/internal/problem/day2"
	"github.com/davidparks11/advent2020/internal/problem/day3"
	"github.com/davidparks11/advent2020/internal/problem/day4"
	"github.com/davidparks11/advent2020/internal/problem/day5"
	"github.com/davidparks11/advent2020/internal/problem/day6"
	"github.com/davidparks11/advent2020/internal/problem/day7"
	"github.com/davidparks11/advent2020/internal/problem/day8"
	"github.com/davidparks11/advent2020/internal/problem/day9"
)

type problemSet map[int]problem.Problem

func NewProblemSet() *problemSet {
	problems := []problem.Problem{
		day1.New(),
		day2.New(),
		day3.New(),
		day4.New(),
		day5.New(),
		day6.New(),
		day7.New(),
		day8.New(),
		day9.New(),
		day10.New(),
	}

	p := make(problemSet)
	for _, problem := range problems {
		p[problem.GetDay()] = problem
	}

	return &p
}

func (p *problemSet) Get(day int) problem.Problem {
	problem, found := (*p)[day]
	if !found {
		log.Fatalf("problem not found in problem set: %d", day)
	}
	return problem
}

const Christmas = 25

func (p *problemSet) Solve(writeToConsole bool, day int) {
	if day != 0 {
		if writeToConsole {
			p.PrintToConsole(day)
		} else {
			p.WriteResultFile(day)
		}
	} else {
		for day := 1; day <= Christmas; day++ {
			if _, found := (*p)[day]; found {
				if writeToConsole {
					p.PrintToConsole(day)
				} else {
					p.WriteResultFile(day)
				}
			}
		}
	}
}

func (p *problemSet) PrintToConsole(day int) {
	results := p.Get(day).Solve()
	if resultStrings, ok := results.([]string); ok {
		log.Printf("Result for Day %d:\n%s\n", day, strings.Join(resultStrings, "\n"))
	} else {
		log.Printf("Result for Day %d: %v\n", day, results)
	}
}

//WriteResult takes result as a string and writes/overwrites the content to a result.txt file
func (p *problemSet) WriteResultFile(day int) {
	problem := p.Get(day)

	fileName := fmt.Sprintf("resources/results/result%d.txt", day)
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	results := problem.Solve()
	if resultStrings, ok := results.([]string); ok {
		_, err = resultFile.WriteString(strings.Join(resultStrings, "\n"))
	} else {
		_, err = resultFile.WriteString(fmt.Sprint(results))
	}
	if err != nil {
		log.Fatal(err)
	}

	if err = resultFile.Close(); err != nil {
		log.Fatal(err)
	}
}
