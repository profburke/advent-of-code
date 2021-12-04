package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spec struct {
	Count       int
	Description string
}

type Rule struct {
	Head string
	Tail []Spec
}

func readData() (rules []Rule) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "contain", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "bags.", "")
		line = strings.ReplaceAll(line, "bag.", "")
		line = strings.ReplaceAll(line, "bags", "bag")
		line = strings.ReplaceAll(line, "no other", "")

		parts := strings.Split(line, "bag")

		rule := Rule{}
		rule.Head = strings.TrimSpace(parts[0])

		for _, part := range parts[1:] {
			subparts := strings.Split(strings.TrimSpace(part), " ")
			if len(subparts) == 1 {
				continue
			}
			spec := Spec{}
			count, _ := strconv.Atoi(subparts[0])
			spec.Count = count
			spec.Description = strings.TrimSpace(strings.Join(subparts[1:], " "))
			rule.Tail = append(rule.Tail, spec)
		}

		rules = append(rules, rule)
	}

	return
}

func findRule(d string, rules []Rule) Rule {
	for _, rule := range rules {
		if rule.Head == d {
			return rule
		}
	}

	fmt.Println("BOGOSITY")
	return rules[0] // this is bogus
}

func containsGold(rule Rule, rules []Rule) bool {
	for _, spec := range rule.Tail {
		if spec.Description == "shiny gold" {
			return true
		}

		newRule := findRule(spec.Description, rules)
		if containsGold(newRule, rules) {
			return true
		}
	}

	return false
}

func part1(rules []Rule) {
	count := 0

	for _, rule := range rules {
		if containsGold(rule, rules) {
			count++
		}
	}

	fmt.Println("Part 1 =", count)
}

func sumUp(d string, rules []Rule) int {
	count := 0
	rule := findRule(d, rules)

	for _, spec := range rule.Tail {
		count += spec.Count // for the bags themselves
		count += spec.Count * sumUp(spec.Description, rules)
	}

	return count
}

func part2(rules []Rule) {
	count := 0

	start := findRule("shiny gold", rules)

	for _, spec := range start.Tail {
		count += spec.Count // for the bags themselves
		count += spec.Count * sumUp(spec.Description, rules)
	}

	fmt.Println("Part 2 =", count)
}

func main() {
	rules := readData()

	part1(rules)
	part2(rules)
}

// Local Variables:
// compile-command: "go build"
// End:
