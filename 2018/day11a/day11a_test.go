package main

import "testing"

func TestGridCreation(t *testing.T) {
	cases := []struct { width int
	} {
		{1}, {10}, {101},
	}

	for _, c := range cases {
		grid := newGrid(c.width)
		got := len(grid)
		if got != c.width {
			t.Errorf("grid width is %d, want %d", got, c.width)
		}
		for i := range grid {
			row := grid[i]
			got := len(row)
			if got != c.width {
				t.Errorf("row %d  width is %d, want %d", i, got, c.width)
			}
		}
	}

}

func TestPowerLevel(t *testing.T) {
	cases := []struct {
		x, y, gridId, want int
	} {
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}

	for _, c := range cases {
		rackId := c.x + 10
		got := powerLevel(rackId, c.y, c.gridId)
		if got != c.want {
			t.Errorf("powerLevel(%d, %d, %d) == %d, want %d", rackId, c.y, c.gridId, got, c.want)
		}
	}
}

func TestSetLevels(t *testing.T) {
	cases := []struct {
		width, gridSerialNumber int
		levels [][]int
	} {
		{4, 0, [][]int{{0, 0, 0, 0}, {0, -4, -3, -2}, {0, -4, -3, -1}, {0, -4, -2, 0}}},
		{4, 7, [][]int{{0, 0, 0, 0}, {0, -4, -2, -1}, {0, -3, -2, 0}, {0, -3, -1, 0}}},

		{6, 12, [][]int{{0, 0, 0, 0, 0, 0}, {0, -3, -2, -1, 1, 2}, {0, -3, -1, 0, 2, 3},
			{0, -2, -1, 1, 3, -5}, {0, -2, 0, 2, 4, -4}, {0, -1, 1, 3, -5, -2}}},
	}

	for _, c := range cases {
		grid := newGrid(c.width)
		setLevels(grid, c.width, c.gridSerialNumber)
		for i := 1; i < c.width; i++ {
			row := grid[i]
			for j := 1; j < c.width; j++ {
				level := row[j]
				if level != c.levels[i][j] {
					t.Errorf("level[%d][%d] == %d, want %d", i, j, level, c.levels[i][j])
				}
			}
		}
	}
}
