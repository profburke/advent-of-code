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

var position = Coordinates(x: 0, y: 2)

let keymap: [Coordinates : String] = [
    Coordinates(x: 0, y: 0) : "-",
    Coordinates(x: 0, y: 1) : "-",
    Coordinates(x: 0, y: 2) : "5",
    Coordinates(x: 0, y: 3) : "-",
    Coordinates(x: 0, y: 4) : "-",

    Coordinates(x: 1, y: 0) : "-",
    Coordinates(x: 1, y: 1) : "A",
    Coordinates(x: 1, y: 2) : "6",
    Coordinates(x: 1, y: 3) : "2",
    Coordinates(x: 1, y: 4) : "-",

    Coordinates(x: 2, y: 0) : "D",
    Coordinates(x: 2, y: 1) : "B",
    Coordinates(x: 2, y: 2) : "7",
    Coordinates(x: 2, y: 3) : "3",
    Coordinates(x: 2, y: 4) : "1",

    Coordinates(x: 3, y: 0) : "-",
    Coordinates(x: 3, y: 1) : "C",
    Coordinates(x: 3, y: 2) : "8",
    Coordinates(x: 3, y: 3) : "4",
    Coordinates(x: 3, y: 4) : "-",

    Coordinates(x: 4, y: 0) : "-",
    Coordinates(x: 4, y: 1) : "-",
    Coordinates(x: 4, y: 2) : "9",
    Coordinates(x: 4, y: 3) : "-",
    Coordinates(x: 4, y: 4) : "-",
]

let lines = data.split(separator: "\n")
for line in lines {
    line.forEach { dir in
                   let (x, y) = (position.x, position.y)
                   if dir == "U" {
                       let newp = Coordinates(x: x, y: min(y + 1, 4))
                       let newkey = keymap[newp]!
                       if newkey != "-" {
                           position = newp
                       }
                   } else if dir == "D" {
                       let newp = Coordinates(x: x, y: max(y - 1, 0))
                       let newkey = keymap[newp]!
                       if newkey != "-" {
                           position = newp
                       }
                   } else if dir == "L" {
                       let newp = Coordinates(x: max(x - 1, 0), y: y)
                       let newkey = keymap[newp]!
                       if newkey != "-" {
                           position = newp
                       }
                   } else {
                       let newp = Coordinates(x: min(x + 1, 4), y: y)
                       let newkey = keymap[newp]!
                       if newkey != "-" {
                           position = newp
                       }
                   }
    }
    print(keymap[position]!, terminator: "")
}
print("")
