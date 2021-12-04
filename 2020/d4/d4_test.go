package main

import (
	"testing"
)

type D4Test struct {
	doc  Document
	want bool
}

// QUESTION: not sure if t.Helper() is needed here.

func looper(t *testing.T, tests []D4Test, field string, validator func(string) bool) {
	t.Helper()
	for _, tt := range tests {
		value := tt.doc.Fields[field]
		got := validator(value)
		if got != tt.want {
			t.Errorf("|%s|: got %v want %v", value, got, tt.want)
		}
	}
}

// QUESTION: would it be better to put all these into one function, grouped by
//       t.Run ?

func TestValidByr(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"byr": "1920"}}, true},
		{Document{Fields: map[string]string{"byr": "2000"}}, true},
		{Document{Fields: map[string]string{"byr": "2002"}}, true},
		{Document{Fields: map[string]string{"byr": ""}}, false},
		{Document{Fields: map[string]string{"byr": "1919"}}, false},
		{Document{Fields: map[string]string{"byr": "2003"}}, false},
	}

	looper(t, tests, "byr", validByr)
}

func TestValidIyr(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"iyr": "2010"}}, true},
		{Document{Fields: map[string]string{"iyr": "2020"}}, true},
		{Document{Fields: map[string]string{"iyr": "2012"}}, true},
		{Document{Fields: map[string]string{"iyr": ""}}, false},
		{Document{Fields: map[string]string{"iyr": "1999"}}, false},
		{Document{Fields: map[string]string{"iyr": "2023"}}, false},
	}

	looper(t, tests, "iyr", validIyr)
}

func TestValidEyr(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"eyr": "2020"}}, true},
		{Document{Fields: map[string]string{"eyr": "2020"}}, true},
		{Document{Fields: map[string]string{"eyr": "2022"}}, true},
		{Document{Fields: map[string]string{"eyr": ""}}, false},
		{Document{Fields: map[string]string{"eyr": "1919"}}, false},
		{Document{Fields: map[string]string{"eyr": "2033"}}, false},
	}

	looper(t, tests, "eyr", validEyr)
}

func TestValidHgt(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"hgt": "60in"}}, true},
		{Document{Fields: map[string]string{"hgt": "190cm"}}, true},
		{Document{Fields: map[string]string{"hgt": "190in"}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": ""}}, false},
		{Document{Fields: map[string]string{"hgt": "190"}}, false},
		{Document{Fields: map[string]string{"hgt": "60cz"}}, false},
	}

	looper(t, tests, "hgt", validHgt)
}

func TestValidHcl(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"hcl": "#000000"}}, true},
		{Document{Fields: map[string]string{"hcl": "#123abc"}}, true},
		{Document{Fields: map[string]string{"hcl": "#847912"}}, true},
		{Document{Fields: map[string]string{"hcl": "#abcdef"}}, true},
		{Document{Fields: map[string]string{"hcl": ""}}, false},
		{Document{Fields: map[string]string{"hcl": "#0012"}}, false},
		{Document{Fields: map[string]string{"hcl": "123456"}}, false},
		{Document{Fields: map[string]string{"hcl": "#123GHI"}}, false},
	}

	looper(t, tests, "hcl", validHcl)
}

func TestValidEcl(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"ecl": "amb"}}, true},
		{Document{Fields: map[string]string{"ecl": "blu"}}, true},
		{Document{Fields: map[string]string{"ecl": "brn"}}, true},
		{Document{Fields: map[string]string{"ecl": "gry"}}, true},
		{Document{Fields: map[string]string{"ecl": "grn"}}, true},
		{Document{Fields: map[string]string{"ecl": "hzl"}}, true},
		{Document{Fields: map[string]string{"ecl": "oth"}}, true},
		{Document{Fields: map[string]string{"ecl": ""}}, false},
		{Document{Fields: map[string]string{"ecl": "Amb"}}, false},
		{Document{Fields: map[string]string{"ecl": "xyz"}}, false},
		{Document{Fields: map[string]string{"ecl": "am"}}, false},
	}

	looper(t, tests, "ecl", validEcl)
}

func TestValidPid(t *testing.T) {
	tests := []D4Test{
		{Document{Fields: map[string]string{"pid": "000000001"}}, true},
		{Document{Fields: map[string]string{"pid": "100000001"}}, true},
		{Document{Fields: map[string]string{"pid": "000239801"}}, true},
		{Document{Fields: map[string]string{"pid": "00000001"}}, false},
		{Document{Fields: map[string]string{"pid": "0000000001"}}, false},
		{Document{Fields: map[string]string{"pid": ""}}, false},
		{Document{Fields: map[string]string{"pid": "0K0000001"}}, false},
	}

	looper(t, tests, "pid", validPid)
}
