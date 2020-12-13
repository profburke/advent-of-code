#!/usr/bin/swift
import Foundation

/*

The coordinate system I use is from http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/

Basically consider a 1/2 pyramid of cubes (think Qbert) and then take the 2d projection....


Anyway, what I tried:

First attempt:  cancel out pairs of "opposites" (e.g. N and S, NE and SW), then replace pairs that are 120 
degrees apart with the "middle" direction (e.g. Replace a NW and NE with a N).  Then keep repeating the process.
But I couldn't quite get that to work.

Second attempt: Starting from origin, "paint" each hex until you reach the destination...


 */

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
    print("Couldn't open \(fileName)")
    exit(1)
}

struct Coordinates {
    let x: Int
    let y: Int
    let z: Int
}

func +(lhs: Coordinates, rhs: Coordinates) -> Coordinates {
    return Coordinates(x: lhs.x + rhs.x, y: lhs.y + rhs.y, z: lhs.z + rhs.z)
}

enum Step: String {
    case nw
    case n
    case ne
    case se
    case s
    case sw

    var offset: Coordinates {
        switch self {
        case .nw:
            return Coordinates(x: -1, y: 1, z: 0)
        case .n:
            return Coordinates(x: 0, y: 1, z: -1)
        case .ne:
            return Coordinates(x: 1, y: 0, z: -1)
        case .se:
            return Coordinates(x: 1, y: -1, z: 0)
        case .s:
            return Coordinates(x: 0, y: -1, z: 1)
        case .sw:
            return Coordinates(x: -1, y: 0, z: 1)
        }
    }
}

func distanceToOrigin(_ position: Coordinates) -> Int {
    return max(abs(position.x), abs(position.y), abs(position.z))
}


var position = Coordinates(x: 0, y: 0, z:0)

var steps = data.split(separator: ",").map { Step(rawValue: String($0).filter { $0 != "\n" })! }

var max = Int.min

_ = steps.map {
    position = position + $0.offset
    let d = distanceToOrigin(position)
    if d > max { max = d }
}

print("\(position) is \(distanceToOrigin(position)) steps from the start.")
print("Max distance is \(max).")

