#!/usr/bin/swift
import Foundation

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
}

extension Coordinates: Hashable {
    var hashValue: Int {
        return x.hashValue ^ y.hashValue &* 16777619
    }

    static func == (lhs: Coordinates, rhs: Coordinates) -> Bool {
        return lhs.x == rhs.x && lhs.y == rhs.y
    }
}

var currentLocation = Coordinates(x: 0, y: 0)

var visited: [Coordinates : Int ] = [:]

data.forEach { char in
               visited[currentLocation] = (visited[currentLocation] ?? 0) + 1
               switch char {
               case ">":
                   currentLocation = Coordinates(x: currentLocation.x + 1, y: currentLocation.y)
               case "<":
                   currentLocation = Coordinates(x: currentLocation.x - 1, y: currentLocation.y)
               case "^":
                   currentLocation = Coordinates(x: currentLocation.x, y: currentLocation.y + 1)
               case "v":
                   currentLocation = Coordinates(x: currentLocation.x, y: currentLocation.y - 1)
               default: // do nothing
                   ()
               }
}

print(visited.keys.count)
