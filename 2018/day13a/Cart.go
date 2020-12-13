// -*- compile-command: "go build"; -*-
package main

import (
	"errors"
	"fmt"
)

type Cart struct {
	Position Position
	Direction Direction
	Behavior Behavior
}

func (c Cart)Move(track [][]TrackType) (newCart Cart, err error) {
	pos := c.Position
	t := track[pos.X][pos.Y]
	if t == Intersection {
		if c.Direction == Up {
			if c.Behavior == LeftTurn {
				c.Position = Position{pos.X, pos.Y - 1}
				c.Direction = Left
			} else if c.Behavior == Straight {
				c.Position = Position{pos.X - 1, pos.Y}
			} else if c.Behavior == RightTurn {
				c.Position = Position{pos.X, pos.Y + 1}
				c.Direction = Right
			}
		} else if c.Direction == Down {
			if c.Behavior == LeftTurn {
				c.Position = Position{pos.X, pos.Y + 1}
				c.Direction = Right
			} else if c.Behavior == Straight {
				c.Position = Position{pos.X + 1, pos.Y}
			} else if c.Behavior == RightTurn {
				c.Position = Position{pos.X, pos.Y - 1}
				c.Direction = Left
			}
		} else if c.Direction == Left {
			if c.Behavior == LeftTurn {
				c.Position = Position{pos.X + 1, pos.Y}
				c.Direction = Down
			} else if c.Behavior == Straight {
				c.Position = Position{pos.X, pos.Y - 1}
			} else if c.Behavior == RightTurn {
				c.Position = Position{pos.X - 1, pos.Y}
				c.Direction = Up
			}
		} else { // Right
			if c.Behavior == LeftTurn {
				c.Position = Position{pos.X - 1, pos.Y}
				c.Direction = Up
			} else if c.Behavior == Straight {
				c.Position = Position{pos.X, pos.Y + 1}
			} else if c.Behavior == RightTurn {
				c.Position = Position{pos.X + 1, pos.Y}
				c.Direction = Down
			}
		}
		c.Behavior = c.Behavior.Next()
	} else if t == Vertical {
		if c.Direction == Up {
			c.Position = Position{pos.X - 1, pos.Y}
		} else if c.Direction == Down {
			c.Position = Position{pos.X + 1, pos.Y}
		} else {
			return Cart{}, errors.New(fmt.Sprintf("direction/track mis-match at (%d, %d) - track: %s - direction: %s", pos.X, pos.Y, t, c.Direction))
		}
	} else if t == Horizontal {
		if c.Direction == Left {
			c.Position = Position{pos.X, pos.Y - 1}
		} else if c.Direction == Right {
			c.Position = Position{pos.X, pos.Y + 1}
		} else {
			return Cart{}, errors.New(fmt.Sprintf("direction/track mis-match at (%d, %d) - track: %s - direction: %s", pos.X, pos.Y, t, c.Direction))
		}
	} else if t == BackslashCurve {
		if c.Direction == Down {
			c.Position = Position{pos.X, pos.Y + 1}
			c.Direction = Right
		} else if c.Direction == Up {
			c.Position = Position{pos.X, pos.Y - 1}
			c.Direction = Left
		} else if c.Direction == Left {
			c.Position = Position{pos.X - 1, pos.Y}
			c.Direction = Up
		} else if c.Direction == Right {
			c.Position = Position{pos.X + 1, pos.Y}
			c.Direction = Down
		}
	} else if t == SlashCurve {
		if c.Direction == Down {
			c.Position = Position{pos.X, pos.Y - 1}
			c.Direction = Left
		} else if c.Direction == Up {
			c.Position = Position{pos.X, pos.Y + 1}
			c.Direction = Right
		} else if c.Direction == Left {
			c.Position = Position{pos.X + 1, pos.Y}
			c.Direction = Down
		} else if c.Direction == Right {
			c.Position = Position{pos.X - 1, pos.Y}
			c.Direction = Up
		}
	} else { // Space
		return Cart{}, errors.New("cart landed on a space")
	}

	return Cart{c.Position, c.Direction, c.Behavior}, nil
}

func cartAt(carts []Cart, row, col int) (cart string) {
	cart = "-"
	for _, c := range carts {
		if c.Position.X == row && c.Position.Y == col {
			switch c.Direction {
			case Up:
				cart = "^"
			case Down:
				cart = "v"
			case Left:
				cart = "<"
			case Right:
				cart = ">"
			}
		}
	}
	
	return cart
}

func isCart(char string) bool {
	return (char == "v" || char == "^" || char == "<" || char == ">")
}

func newCart(row, col int, char string) (cart Cart) {
	var direction Direction
	switch char {
	case "^":
		direction = Up
	case "v":
		direction = Down
	case ">":
		direction = Right
	case "<":
		direction = Left
	default:
		// should raise an error, but I'm getting lazy
		direction = UnknownDirection
	}
	return Cart{Position{row, col}, direction, LeftTurn}
}

func collisionOccurred(carts []Cart) (result bool, position Position) {
	locations := make(map[Position]bool)

	for _, c := range carts {
		if occupied := locations[c.Position]; occupied {
			return true, c.Position
		} else {
			locations[c.Position] = true
		}
	}
	
	return false, Position{-1, -1}
}
