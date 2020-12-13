package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Ephemeris struct {
	X int
	Y int
	Dx int
	Dy int
}

func max(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 { a = a * -1 }
	return a
}

func readEphemerides(r io.Reader) (ephemerides []*Ephemeris, maxSteps int, err error) {
	scanner := bufio.NewScanner(r)
	maxSteps = 0
	
	for scanner.Scan() {
		line := scanner.Text()
		var (
			x, y int
			dx, dy int
		)

		matches, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy)
		if err != nil || matches != 4 {
			return ephemerides, 0, err
		}

		ephemeris := new(Ephemeris)
		ephemeris.X = x
		ephemeris.Y = y
		ephemeris.Dx = dx
		ephemeris.Dy = dy
		ephemerides = append(ephemerides, ephemeris)

		var c, d int

		// Try to get a handle on how long this needs to run
		// assumption: time it takes farthest/slowest particle to get to 0, 0
		// I'm being a little sloppy about directions, etc. and am assuming that
		// no particle is initially heading away from 0, 0
		
		if dx != 0 {
			c = abs(x)/abs(dx)
		}

		if dy != 0 {
			d = abs(y)/abs(dy)
		}

		if c > maxSteps {
			maxSteps = c
		}
		if d > maxSteps {
			maxSteps = d
		}
		
	}

	return ephemerides, maxSteps, nil
}

func display(time int, ephemerides []*Ephemeris) bool {
	const size = 19
	var screen [2 * size + 1][2 * size + 1]string
	
	for x := -size; x < (size + 1); x++ {
		for y := -size; y < (size + 1); y++ {
			screen[x + size][y + size] = " "
		}
	}

	visibleCharacters := false
	for _, e := range ephemerides {
		if e.X > -(size + 1) && e.X < (size + 1) && e.Y > -(size + 1) && e.Y < (size + 1) {
			screen[e.Y + size][e.X + size] = "#"
			visibleCharacters = true
		}
	}

	// only print if there's anything to see...
	if visibleCharacters {
		fmt.Println("\033[2J") // clear the screen
		fmt.Println("Time :", time)
		for i := 0; i < len(screen); i++  {
			line := screen[i]
			for j := 0; j < len(line); j++ {
				fmt.Print(screen[i][j])
			}
			fmt.Println()
		}
		fmt.Println("\n\n")
	}

	return visibleCharacters
}

func main() {
	ephemerides, maxSteps, err := readEphemerides(os.Stdin)
	if err != nil {
		fmt.Println("trouble reading ephemerides: ", err)
		os.Exit(1)
	}
	fmt.Println("Max steps: ", maxSteps)
	
	step := 0
	for step < 2 * maxSteps {
		displayed := display(step, ephemerides)
		
		// pause
		if displayed {
			time.Sleep(200 * time.Millisecond)
		}
		
		// update positions
		for _, ephemeris := range ephemerides {
			ephemeris.X += ephemeris.Dx
			ephemeris.Y += ephemeris.Dy
		}
		step++
	}

}
