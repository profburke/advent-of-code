package main

type Command struct {
	Direction string
	Distance  int
}

type Coordinates struct {
	X int
	Y int
}

type Heading int

const (
	N Heading = iota
	E
	S
	W
)

// Local Variables:
// compile-command: "go build"
// End:
