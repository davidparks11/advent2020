package problem

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Problem interface {
	Solve() interface{}
	GetDay() int
}

type DailyProblem struct {
	Day int
}

func (d *DailyProblem) GetDay() int {
	return d.Day
}

//GetInputLines reads an input.txt file and returns its contents separated by lines as a string array
func (d *DailyProblem) GetInputLines() []string {
	if d.Day == 0 {
		log.Fatal("error getting input lines with no set day")
	}
	fileName := fmt.Sprintf("resources/inputs/input%d.txt", d.Day)
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	inputFile.Close()

	return lines
}

//IntsFromStrings takes a string array and returns array of those strings converted to ints
func IntsFromStrings(inputLines []string) []int {
	input := make([]int, len(inputLines))
	for i, line := range inputLines {
		intValue, err := strconv.Atoi(line)
		if err != nil {
			return []int{}
		}
		input[i] = intValue
	}
	return input
}
