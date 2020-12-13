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

enum Giver: Int {
    case santa = 0
    case robo = 1
}
var who = Giver.santa
var currentLocation: [Coordinates] = [Coordinates(x: 0, y: 0), Coordinates(x: 0, y: 0)]
var visited: [Coordinates : Int ] = [:]

data.forEach { char in
               let w = who.rawValue
               visited[currentLocation[w]] = (visited[currentLocation[w]] ?? 0) + 1
               switch char {
               case ">":
                   currentLocation[w] = Coordinates(x: currentLocation[w].x + 1, y: currentLocation[w].y)
               case "<":
                   currentLocation[w] = Coordinates(x: currentLocation[w].x - 1, y: currentLocation[w].y)
               case "^":
                   currentLocation[w] = Coordinates(x: currentLocation[w].x, y: currentLocation[w].y + 1)
               case "v":
                   currentLocation[w] = Coordinates(x: currentLocation[w].x, y: currentLocation[w].y - 1)
               default: // do nothing
                   ()
               }
               if who == .santa {
                   who = .robo
               } else {
                   who = .santa
               }
}

print(visited.keys.count)
