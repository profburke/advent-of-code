package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readData() (data []Command) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		command := Command{Direction: line[0:1], Distance: distance}
		data = append(data, command)
	}

	return
}

// clean this up by using interfaces? (i.e. smooth over the waypoint vs heading)
type SteeringFunction func(Coordinates, Coordinates, Heading, Command) (Coordinates, Coordinates, Heading)

func doit(commands []Command, partN int, steer SteeringFunction) {
	position := Coordinates{X: 0, Y: 0}
	waypoint := Coordinates{X: 10, Y: -1}
	heading := E

	for _, command := range commands {
		position, waypoint, heading = steer(position, waypoint, heading, command)
	}

	manhattan := int(math.Abs(float64(position.X)) + math.Abs(float64(position.Y)))
	fmt.Printf("Part %d = %d\n", partN, manhattan)
}

func main() {
	commands := readData()

	doit(commands, 1, steer)
	doit(commands, 2, steerByWaypoint)
}

// Local Variables:
// compile-command: "go build"
// End:
