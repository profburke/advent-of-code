package main

import "fmt"

type Grid struct {
	Grid    [][][]string
	Current int
	Next    int
}

const (
	Floor        = "."
	FreeSeat     = "L"
	OccupiedSeat = "#"
)

func (g Grid) OnGrid(r, c int) bool {
	return r >= 0 && r < len(g.Grid[g.Current][0]) && c >= 0 && c < len(g.Grid[g.Current])
}

func (g Grid) NeighborCount(r, c int) int {
	neighbors := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			if g.OnGrid(r+dr, c+dc) && g.Grid[g.Current][r+dr][c+dc] == OccupiedSeat {
				neighbors++
			}
		}
	}

	return neighbors
}

func (g Grid) UpdateCell(r, c int) {
	switch g.Grid[g.Current][r][c] {
	case Floor:
		g.Grid[g.Next][r][c] = Floor
	case FreeSeat:
		if g.NeighborCount(r, c) == 0 {
			g.Grid[g.Next][r][c] = OccupiedSeat
		} else {
			g.Grid[g.Next][r][c] = FreeSeat
		}
	case OccupiedSeat:
		if g.NeighborCount(r, c) == 4 {
			g.Grid[g.Next][r][c] = FreeSeat
		} else {
			g.Grid[g.Next][r][c] = OccupiedSeat
		}
	}
}

func (g *Grid) Step() {
	for _, row := range g.Grid[g.Current] {
		fmt.Println(row)
	}

	for r := 0; r < len(g.Grid[g.Current][0]); r++ {
		for c := 0; c < len(g.Grid[g.Current]); c++ {
			g.UpdateCell(r, c)
		}
	}

	fmt.Println("\n\n\n")

	for _, row := range g.Grid[g.Next] {
		fmt.Println(row)
	}

	g.Current, g.Next = g.Next, g.Current
}

func (g *Grid) IsStable() bool {
	fmt.Println("stable indices", g.Current, g.Next)
	for r := 0; r < len(g.Grid[g.Current][0]); r++ {
		for c := 0; c < len(g.Grid[g.Current]); c++ {
			fmt.Println("stable", r, c, g.Grid[g.Current][r][c], g.Grid[g.Next][r][c])
			if g.Grid[g.Current][r][c] != g.Grid[g.Next][r][c] {
				return false
			}
		}
	}

	return true
}

func (g Grid) Count(state string) int {
	count := 0

	for r := 0; r < len(g.Grid[g.Current][0]); r++ {
		for c := 0; c < len(g.Grid[g.Current]); c++ {
			if g.Grid[g.Current][r][c] == state {
				count++
			}
		}
	}

	return count
}

// Local Variables:
// compile-command: "go build"
// End:
