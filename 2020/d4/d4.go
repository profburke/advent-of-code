package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Document struct {
	Fields map[string]string
}

func NewDocument() Document {
	d := Document{}
	d.Fields = make(map[string]string)

	return d
}

func readData() (documents []Document) {
	scanner := bufio.NewScanner(os.Stdin)

	current := NewDocument()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			documents = append(documents, current)
			current = NewDocument()
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {
			parts := strings.Split(field, ":")
			current.Fields[parts[0]] = parts[1]
		}
	}
	documents = append(documents, current)

	return
}

func hasRequiredFields(document Document) bool {
	_, byr := document.Fields["byr"]
	_, iyr := document.Fields["iyr"]
	_, eyr := document.Fields["eyr"]
	_, hgt := document.Fields["hgt"]
	_, hcl := document.Fields["hcl"]
	_, ecl := document.Fields["ecl"]
	_, pid := document.Fields["pid"]

	return byr && iyr && eyr && hgt && hcl && ecl && pid
}

func yearInRange(val string, lower, upper int) bool {
	year, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	return (year >= lower && year <= upper)
}

func validByr(byr string) bool {
	return yearInRange(byr, 1920, 2002)
}

func validIyr(iyr string) bool {
	return yearInRange(iyr, 2010, 2020)
}

func validEyr(eyr string) bool {
	return yearInRange(eyr, 2020, 2030)
}

func validHgt(hgt string) bool {
	re, _ := regexp.Compile("(\\d+)(in|cm)")
	parts := re.FindStringSubmatch(hgt)

	// first item in slice is entire match
	if len(parts) != 3 {
		return false
	}

	val, _ := strconv.Atoi(parts[1])

	if parts[2] == "cm" {
		return (val >= 150 && val <= 193)
	} else if parts[2] == "in" {
		return (val >= 59 && val <= 76)
	} else {
		return false
	}
}

func validHcl(hcl string) bool {
	matched, _ := regexp.Match(`#[0-9a-f]{6}`, []byte(hcl))
	return matched
}

func validEcl(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" ||
		ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

func validPid(pid string) bool {
	matched, _ := regexp.Match(`^\d{9}$`, []byte(pid))
	return matched
}

func fieldsAreValid(document Document) bool {
	return validByr(document.Fields["byr"]) && validIyr(document.Fields["iyr"]) &&
		validEyr(document.Fields["eyr"]) && validHgt(document.Fields["hgt"]) &&
		validHcl(document.Fields["hcl"]) && validEcl(document.Fields["ecl"]) &&
		validPid(document.Fields["pid"])
}

func countValidDocuments(documents []Document, validator func(Document) bool) int {
	validPassports := 0

	for _, document := range documents {
		if validator(document) {
			validPassports++
		}
	}

	return validPassports
}

func part1(documents []Document) {
	fmt.Println("Part 1 = ", countValidDocuments(documents,
		func(document Document) bool { return hasRequiredFields(document) }))
}

func part2(documents []Document) {
	fmt.Println("Part 2 = ", countValidDocuments(documents,
		func(document Document) bool {
			return hasRequiredFields(document) && fieldsAreValid(document)
		}))
}

func main() {
	documents := readData()

	part1(documents)
	part2(documents)
}

// Local Variables:
// compile-command: "go build"
// End:
