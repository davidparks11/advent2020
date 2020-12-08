package advent

import (
	"bufio"
	"fmt"
	"os"
)

type Problem interface {
	Solve()
}

type dailyProblem struct {
	day int
	name string
}

//WriteResult takes result as a string and writes/overwrites the content to a result.txt file
func (d *dailyProblem) WriteResult(result string) {
	fmt.Printf("Result for Day %d, the %s Problem: %v", d.day, d.name, result)
	fileName := fmt.Sprintf("resources/results/result%d.txt", d.day)
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Sprintf(err.Error())
		return
	}
	resultFile.WriteString(result)
	if err := resultFile.Close(); err != nil {
		fmt.Sprintf(err.Error())
		return
	}
}

//GetInputFile reads an input.txt file and returns its contents separated by lines as a string array
func (d *dailyProblem) GetInputFile() []string {
	fileName := fmt.Sprintf("resources/inputs/input%d.txt", d.day)
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
