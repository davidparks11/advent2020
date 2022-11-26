package day10

import (
	"fmt"
	"sort"

	. "github.com/davidparks11/advent2020/internal/problem"
)

type adapterArray struct {
	DailyProblem
}

func New() Problem {
	return &adapterArray{
		DailyProblem{
			Day: 10,
		},
	}
}

func (a *adapterArray) Solve() interface{} {
	input := a.GetInputLines()
	var results []int
	results = append(results, oneAndThreeJouleProduct(IntsFromStrings(input)))
	return results
}

func oneAndThreeJouleProduct(input []int) int {
	sort.Ints(input)
	input = append([]int{0}, input...)
	input = append(input, input[len(input) - 1]+3) //add device adapter

	diffs := make(map[int]int)

	for i := 1; i < len(input); i++ {
		if input[i] - input[i-1] > 3 {
			panic("adapterArray: invalid adapter difference")
		}
		diffs[input[i]-input[i-1]]++ //increment the difference between current and previous rating
	}
	fmt.Println(diffs)
	return diffs[1] * diffs[3]
}
