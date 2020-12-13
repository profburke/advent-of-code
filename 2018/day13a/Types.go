// -*- compile-command: "go build"; -*-
package main

import "fmt"

type Direction int
const (
	UnknownDirection Direction = iota
	Up
	Down
	Left
	Right
)

func (d Direction) String() string {
	switch d {
	case UnknownDirection:
		return "direction ???"
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	default:
		return "Right"
	}
}

type TrackType int
const (
	UnknownTrackType TrackType = iota
	Vertical
	Horizontal
	Intersection
	SlashCurve
	BackslashCurve
	Space
)

func trackType(char string) (t TrackType) {
	switch char {
	case "+":
		t = Intersection
	case "|", "v", "^":
		t = Vertical
	case "-", "<", ">":
		t = Horizontal
	case "/":
		t = SlashCurve
	case "\\":
		t = BackslashCurve
	default:
		t = Space
	}
	return t
}

// which way cart will turn at intersection
type Behavior int
const (
	UnknownBehavaior Behavior = iota
	LeftTurn
	Straight
	RightTurn
)

func (b Behavior) Next() Behavior {
	if b == LeftTurn {
		return Straight
	} else if b == Straight {
		return RightTurn
	} else {
		return LeftTurn
	}
}

func (b Behavior) String() string {
	switch b {
	case LeftTurn:
		return "LT"
	case Straight:
		return "S"
	default:
		return "RT"
	}
}

type Position struct {
	X, Y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
