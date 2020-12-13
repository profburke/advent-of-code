#!/usr/bin/env swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName).dropLast() else {
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

var position = Coordinates(x: 1, y: 1)

let keymap: [Coordinates : Int] = [
    Coordinates(x: 0, y: 0) : 7,
    Coordinates(x: 0, y: 1) : 4,
    Coordinates(x: 0, y: 2) : 1,
    Coordinates(x: 1, y: 0) : 8,
    Coordinates(x: 1, y: 1) : 5,
    Coordinates(x: 1, y: 2) : 2,
    Coordinates(x: 2, y: 0) : 9,
    Coordinates(x: 2, y: 1) : 6,
    Coordinates(x: 2, y: 2) : 3,
    ]

let lines = data.split(separator: "\n")
for line in lines {
    line.forEach { dir in
                   let (x, y) = (position.x, position.y)
                   if dir == "U" {
                       position = Coordinates(x: x, y: min(y + 1, 2))
                   } else if dir == "D" {
                       position = Coordinates(x: x, y: max(y - 1, 0))
                   } else if dir == "L" {
                       position = Coordinates(x: max(x - 1, 0), y: y)
                   } else {
                       position = Coordinates(x: min(x + 1, 2), y: y)
                   }
    }
    print(keymap[position]!, terminator: "")
}
print("")
