package main

import (
	"reflect"
	"testing"
)

func TestDetermineSteps(t *testing.T) {
	tests := []struct {
		line string
		want []Direction
	}{
		{line: "seswneswswsenwwnwse", want: []Direction{
			Southeast,
			Southwest,
			Northeast,
			Southwest,
			Southwest,
			Southeast,
			Northwest,
			West,
			Northwest,
			Southeast,
		}},
		{line: "wseweeenwnesenwwwswnew", want: []Direction{
			West,
			Southeast,
			West,
			East,
			East,
			East,
			Northwest,
			Northeast,
			Southeast,
			Northwest,
			West,
			West,
			Southwest,
			Northeast,
			West,
		}},
	}

	for _, tt := range tests {
		got := determineSteps(tt.line)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func TestWalk(t *testing.T) {
	tests := []struct {
		steps []Direction
		want  Coordinates
	}{
		{
			steps: []Direction{
				Southeast,
				Southwest,
				Northeast,
				Southwest,
				Southwest,
				Southeast,
				Northwest,
				West,
				Northwest,
				Southeast,
			},
			want: Coordinates{X: -2, Y: 3},
		},
		{
			steps: []Direction{
				West,
				Southeast,
				West,
				East,
				East,
				East,
				Northwest,
				Northeast,
				Southeast,
				Northwest,
				West,
				West,
				Southwest,
				Northeast,
				West,
			},
			want: Coordinates{X: -2, Y: -1},
		},
	}

	for _, tt := range tests {
		got := walk(tt.steps)

		if got != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

// Local Variables:
// compile-command: "go build"
// End:
