package main

import "testing"

func TestPatternToInt(t *testing.T) {
	cases := []struct {
		pattern string
		want int
	} {
		{".....", 0},
		{"#..#.", 18},
		{".##.#", 13},
		{"#####", 31},
	}

	for _, c := range cases {
		got, err := patternToInt(c.pattern)
		if err != nil {
			t.Errorf("patternToInt(%s) want %d, got error: %s", c.pattern, c.want, err)
		} else if got != c.want {
			t.Errorf("patternToInt(%s) == %d, want %d", c.pattern, got, c.want)
		}
	}
}

func TestParseRule(t *testing.T) {
	cases := []struct {
		line string
		index int
		want string
	} {
		{"...#. => #", 2, "#"},
		{"#..#. => .", 18, "."},
	}

	rules := make(map[int]string)

	for _, c := range cases {
		err := parseRule(rules, c.line)
		got, ok := rules[c.index]

		if err != nil {
			t.Errorf("parseRule(rule, %s) raised error: %s", c.line, err)
		} else if ok != true {
			t.Errorf("parseRule(rule, %s) did not set map entry", c.line)
		} else if got != c.want {
			t.Errorf("parseRule(rule, %s) got %s, wanted %s", c.line, got, c.want)
		}
	}
}
