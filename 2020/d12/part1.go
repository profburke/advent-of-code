package main

func move(position Coordinates, heading Heading, distance int) (newPosition Coordinates) {
	switch heading {
	case N:
		newPosition.X = position.X
		newPosition.Y = position.Y - distance
	case E:
		newPosition.X = position.X + distance
		newPosition.Y = position.Y
	case S:
		newPosition.X = position.X
		newPosition.Y = position.Y + distance
	case W:
		newPosition.X = position.X - distance
		newPosition.Y = position.Y
	}

	return
}

func changeHeading(heading Heading, command Command) (newHeading Heading) {
	modifier := 1
	if command.Direction == "L" {
		modifier = -1
	}

	heading += 4
	switch command.Distance {
	case 90:
		heading += Heading(1 * modifier)
	case 180:
		heading += Heading(2 * modifier)
	case 270:
		heading += Heading(3 * modifier)
	}
	newHeading = heading % 4

	return
}

func steer(position, waypoint Coordinates, heading Heading, command Command) (newPosition, newWaypoint Coordinates, newHeading Heading) {
	newWaypoint = waypoint

	switch command.Direction {
	case "F":
		newPosition = move(position, heading, command.Distance)
		newHeading = heading
	case "R", "L":
		newPosition.X, newPosition.Y = position.X, position.Y
		newHeading = changeHeading(heading, command)
	case "N":
		newPosition.X, newPosition.Y = position.X, position.Y-command.Distance
		newHeading = heading
	case "E":
		newPosition.X, newPosition.Y = position.X+command.Distance, position.Y
		newHeading = heading
	case "S":
		newPosition.X, newPosition.Y = position.X, position.Y+command.Distance
		newHeading = heading
	case "W":
		newPosition.X, newPosition.Y = position.X-command.Distance, position.Y
		newHeading = heading
	}

	return
}

// Local Variables:
// compile-command: "go build"
// End:
