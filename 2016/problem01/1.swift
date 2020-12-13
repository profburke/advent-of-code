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
var position = (0, 0)

func turn(_ direction: Direction) {
    let r = (heading.rawValue + direction.rawValue + CompassPoint.n) % CompassPoint.n
    heading = CompassPoint(rawValue: r)!
}

func step(_ n: Int) {
    switch(heading) {
    case .north:
        position.1 += n
    case .east:
        position.0 += n
    case .south:
        position.1 -= n
    case .west:
        position.0 -= n
    }
}

func taxicabDistance() -> Int {
    return abs(position.0) + abs(position.1)
}

let orders = data.filter { $0 != " " }.split(separator: ",")
for var order in orders {
    let dir = order.removeFirst().direction
    let n = Int(order)!
    turn(dir)
    step(n)
}

print(position)
print(taxicabDistance())
