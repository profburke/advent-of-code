package main

import "testing"

func TestRange(t *testing.T) {
	tests := []struct {
		r     Range
		value int
		want  bool
	}{
		{r: Range{Low: 1, High: 3}, value: 4, want: false},
		{r: Range{Low: 5, High: 7}, value: 4, want: false},
		{r: Range{Low: 5, High: 7}, value: 5, want: true},
	}

	for _, tt := range tests {
		got := tt.r.Contains(tt.value)
		if got != tt.want {
			t.Errorf("for %d in %v: got %v want %v",
				tt.value, tt.r, got, tt.want)
		}
	}
}

func TestField(t *testing.T) {
	tests := []struct {
		f     Field
		value int
		want  bool
	}{
		{f: Field{First: Range{Low: 1, High: 3}, Second: Range{Low: 5, High: 7}}, value: 4, want: false},
		{f: Field{First: Range{Low: 1, High: 3}, Second: Range{Low: 5, High: 7}}, value: 3, want: true},
		{f: Field{First: Range{Low: 1, High: 3}, Second: Range{Low: 5, High: 7}}, value: 6, want: true},
	}

	for _, tt := range tests {
		got := tt.f.Valid(tt.value)
		if got != tt.want {
			t.Errorf("for %d in %v: got %v want %v",
				tt.value, tt.f, got, tt.want)
		}
	}
}
