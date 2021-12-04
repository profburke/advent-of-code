package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		point   Coordinates
		command Command
		want    Coordinates
	}{
		{point: Coordinates{X: 10, Y: -4}, command: Command{Direction: "R", Distance: 90}, want: Coordinates{X: 4, Y: 10}},
		{point: Coordinates{X: 10, Y: -4}, command: Command{Direction: "L", Distance: 90}, want: Coordinates{X: 4, Y: 10}},
	}

	for _, tt := range tests {
		got := rotateWaypoint(tt.point, tt.command)
		if got != tt.want {
			t.Errorf("%v - %v :: got %v wanted %v", tt.point, tt.command, got, tt.want)
		}
	}
}
