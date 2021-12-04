package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Document struct {
	Fields map[string]string
}

func validHgt(hgt string) bool {
	re, _ := regexp.Compile("(\\d+)(in|cm)")
	parts := re.FindStringSubmatch(hgt)

	if len(parts) != 2 {
		return false
	}

	val, _ := strconv.Atoi(parts[0])

	fmt.Println("p0", parts[0], "p1", parts[1])

	if parts[1] == "cm" {
		return (val >= 150 && val <= 193)
	} else if parts[1] == "in" {
		return (val >= 59 && val <= 76)
	} else {
		return false
	}
}

func main() {
	tests := []struct {
		doc  Document
		want bool
	}{
		{Document{Fields: map[string]string{"hgt": "60in"}}, true},
		{Document{Fields: map[string]string{"hgt": "190cm"}}, true},
		{Document{Fields: map[string]string{"hgt": "190in"}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": "190"}}, false},
		{Document{Fields: map[string]string{"hgt": "60cz"}}, false},
	}

	for _, tt := range tests {
		fmt.Println(validHgt(tt.doc.Fields["hgt"]))
	}

}

// Local Variables:
// compile-command: "go build"
// End:
