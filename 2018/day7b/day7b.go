// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Worker struct {
	Task string
	whenFinished int
}

type Constraint struct {
	Head string
	Tail string
}

func readConstraints(r io.Reader) (constraints []Constraint, err error) {
	scanner := bufio.NewScanner(r)
	
	for scanner.Scan() {
		var (
			head string
			tail string
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &head, &tail)
		if err != nil || matches != 2 {
			fmt.Fprintln(os.Stderr, "error reading constraint: ", err)
			os.Exit(1)
		}

		constraint := Constraint{head, tail}
		constraints = append(constraints, constraint)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return constraints, nil
}

func nodes(constraints []Constraint) (nodes []string) {
	existing := make(map[string]bool)

	for _, c := range constraints {
		existing[c.Head] = true
		existing[c.Tail] = true
	}

	for node := range existing {
		nodes = append(nodes, node)
	}
	
	sort.Strings(nodes)
	return nodes
}

func prerequisites(constraints []Constraint) (prerequisites map[string][]string) {
	prerequisites = make(map[string][]string)

	for _, constraint := range constraints {
		if _, ok := prerequisites[constraint.Tail]; ok {
			parents := prerequisites[constraint.Tail]
			parents = append(parents, constraint.Head)
			prerequisites[constraint.Tail] = parents
		} else {
			prerequisites[constraint.Tail] = []string{constraint.Head}
		}
	}

	for k,v := range prerequisites {
		sorted := v
		sort.Strings(sorted)
		prerequisites[k] = sorted
	}
	
	return prerequisites
}

func dependencies(constraints []Constraint) (dependencies map[string][]string) {
	dependencies = make(map[string][]string)

	for _, constraint := range constraints {
		if _, ok := dependencies[constraint.Head]; ok {
			children := dependencies[constraint.Head]
			children = append(children, constraint.Tail)
			dependencies[constraint.Head] = children
		} else {
			dependencies[constraint.Head] = []string{constraint.Tail}
		}
	}

	for k,v := range dependencies {
		sorted := v
		sort.Strings(sorted)
		dependencies[k] = sorted
	}
	
	return dependencies
}

// figure out which rules have no prerequisites
func sources(nodes []string, prerequisites map[string][]string) (sources []string) {
	for _, node := range nodes {
		_, ok := prerequisites[node]
		if !ok {
			sources = append(sources, node)
		}
	}

	sort.Strings(sources)
	return sources
}

func contains(a []string, e string) bool {
	for _, v := range a {
		if v == e {
			return true
		}
	}
	return false
}

func allMet(node string, completed []string, prerequisites map[string][]string) bool {
	for _, rule := range prerequisites[node] {
		if !contains(completed, rule) {
			return false
		}
	}
	
	return true
}

func duration(rule string) int {
	r := rune(strings.ToUpper(rule)[0])
	return int(r - 'A') + 61
}

func anyBusy(workers []*Worker) bool {
	for _, w := range workers {
		if w.whenFinished >= 0 {
			return true
		}
	}
	return false
}

func main() {
	constraints, err := readConstraints(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading constraints: ", err)
		os.Exit(1)
	}

	nodes := nodes(constraints)
	dependencies := dependencies(constraints)
	prerequisites := prerequisites(constraints)

	// enabled is set of all rules which have all their prerequisites met
	// initially that's all the sources
	enabled := sources(nodes, prerequisites)

	var completed []string
	result := ""

	const nWorkers = 6
	var workers [nWorkers]*Worker
	for i := 0; i < nWorkers; i++ {
		w := new(Worker)
		w.Task = ""
		w.whenFinished = -1
		workers[i] = w
	}

	step := 0
	for len(enabled) > 0 || anyBusy(workers[:]) {
		for _, w := range workers {
			if w.Task == "" && len(enabled) > 0 { // assign an enabled task to each free worker
				w.Task = enabled[0]
				w.whenFinished = step + duration(w.Task) - 1
				enabled = enabled[1:]
			}
		}
		for _, w := range workers {
			if w.whenFinished == step { // each worker releases completed tasks 
				task := w.Task
				w.Task = ""
				w.whenFinished = -1
				completed = append(completed, task)
				result += task
				for _, rule := range dependencies[task] {
					if !contains(completed, rule) && allMet(rule, completed, prerequisites) && !contains(enabled, rule) {
						enabled = append(enabled, rule)
					}
				}
			}
		}
		step++
		sort.Strings(enabled)
	}

	fmt.Println(fmt.Sprintf("Result (%d): '%s'", len(result), result))
	fmt.Println("Steps: ", step)
}
