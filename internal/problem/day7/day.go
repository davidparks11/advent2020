package day7

import (
	"regexp"
	"strconv"
	"strings"

	. "github.com/davidparks11/advent2020/internal/problem"
)

var _ Problem = &handyHaversacks{}

type handyHaversacks struct {
	DailyProblem
}

func New() Problem {
	return &handyHaversacks{
		DailyProblem{Day: 7},
	}
}

func (h *handyHaversacks) Solve() interface{} {
	input := h.GetInputLines()
	var results []int
	results = append(results, bagTypesContainingGold(input))
	results = append(results, bagsInsideOfBagType(input))
	return results
}

func bagTypesContainingGold(bagRules []string) int {
	rulesSet := make(map[string][]string)
	for _, rule := range bagRules {
		parent, children := parseWithoutCount(rule)
		for _, child := range children {
			if parents, found := rulesSet[child]; found {
				rulesSet[child] = append(parents, parent)
			} else {
				rulesSet[child] = []string{parent}
			}
		}
	}

	processed := make(map[string]struct{})
	return recurseCountBagTypes("shiny gold", rulesSet, processed) - 1
}

func bagsInsideOfBagType(bagRules []string) int {
	rulesSet := make(map[string]map[string]int)
	for _, rule := range bagRules {
		parent, children := parseRule(rule)
		rulesSet[parent] = children
	}

	return recurseCountBags("shiny gold", rulesSet) - 1
}

func recurseCountBags(name string, rules map[string]map[string]int) int {
	count := 1

	for child, amount := range rules[name] {
		count += amount * recurseCountBags(child, rules)
	}
	return count
}

func recurseCountBagTypes(name string, rules map[string][]string, processed map[string]struct{}) int {
	parents := rules[name]
	count := 1
	for _, parent := range parents {
		if _, found := processed[parent]; found {
			continue
		} else {
			processed[parent] = struct{}{}
		}
		count += recurseCountBagTypes(parent, rules, processed)
	}

	return count
}

var splitBagsRgx = regexp.MustCompile("( bag, | bags, )")

// inefficient, but I don't care
func parseWithoutCount(rule string) (string, []string) {
	parent, childBagCounts := parseRule(rule)
	children := make([]string, len(childBagCounts))
	for bagName := range childBagCounts {
		children = append(children, bagName)
	}

	return parent, children
}

func parseRule(rule string) (string, map[string]int) {
	sides := strings.Split(rule, " contain ")
	parentBag := sides[0][:len(sides[0])-5] //get left side bag without the word 'bags'

	if sides[1] == "no other bags." {
		return parentBag, nil
	}

	rightSideTrimIndex := 5
	if sides[1][:len(sides[1])-4] == "bag." {
		rightSideTrimIndex = 4 //last bag is singular
	}

	rightSide := sides[1][:len(sides[1])-rightSideTrimIndex]
	rightSideBags := splitBagsRgx.Split(rightSide, -1)

	bagsToCount := make(map[string]int)
	for _, v := range rightSideBags {
		spaceIndex := strings.Index(v, " ")
		num, err := strconv.Atoi(v[:spaceIndex])
		if err != nil {
			panic("input not parsable - rule: " + rule + "\nErr: " + err.Error())
		}

		trimmedString := strings.TrimSpace(v[spaceIndex+1:])
		if strings.Contains(trimmedString, ",") {
			panic("comma at: " + rule)
		}
		bagsToCount[trimmedString] = num
	}

	return parentBag, bagsToCount
}
