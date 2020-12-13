#!/usr/bin/env swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName).lowercased().dropLast() else {
    print("Couldn't open \(fileName)")
    exit(1)
}

enum CompassPoint: Int {
    case north = 0
    case east
    case south
    case west

    static let n = 4
}

enum Direction: Int {
    case right = 1
    case left = -1
}

extension Character {

    var direction: Direction {
        if self == "l" {  // hack
            return .left
        } else {
            return .right
        }
    }
}

var heading = CompassPoint.north

struct Coordinates: CustomStringConvertible {
    let x: Int
    let y: Int

    var description: String {
        return "\(x):\(y)"
    }
}

var position = Coordinates(x:0, y:0)
var seen = Set<String>()
seen.insert(position.description)

func turn(_ direction: Direction) {
    let r = (heading.rawValue + direction.rawValue + CompassPoint.n) % CompassPoint.n
    heading = CompassPoint(rawValue: r)!
}

func endgame() {
    print(taxicabDistance())
    exit(0)
}

func step(_ n: Int) {
    for _ in 0..<n {
        switch(heading) {
        case .north:
            position = Coordinates(x: position.x, y: position.y + 1)
        case .east:
            position = Coordinates(x: position.x + 1, y: position.y)
        case .south:
            position = Coordinates(x: position.x, y: position.y - 1)
        case .west:
            position = Coordinates(x: position.x - 1, y: position.y)
        }
        if seen.contains(position.description) { endgame() }
        seen.insert(position.description)
    }
}

func taxicabDistance() -> Int {
    return abs(position.x) + abs(position.y)
}

let orders = data.filter { $0 != " " }.split(separator: ",")
for var order in orders {
    let dir = order.removeFirst().direction
    let n = Int(order)!
    turn(dir)
    step(n)
}
