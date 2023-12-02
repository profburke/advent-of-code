#!/usr/bin/env swift

import Darwin

var grid: [[String]] = []
var path: [String] = []
var pptr = 0

enum Command {
    case step(Int)
    case left
    case right
}

enum Orientation {
    case left
    case right
    case top
    case down

    func left() -> Orientation {
        switch self {
        case .left:
            return .down
        case .right:
            return .top
        case .top:
            return .left
        case .down:
            return .right
        }
    }

    func right() -> Orientation {
        switch self {
        case .left:
            return .top
        case .right:
            return .down
        case .top:
            return .right
        case .down:
            return .left
        }
    }
    
    var delta: Point {
        switch self {
        case .left:
            return Point(x: -1, y: 0)
        case .right:
            return Point(x: 1, y: 0)
        case .top:
            return Point(x: 0, y: -1)
        case .down:
            return Point(x: 0, y: 1)
        }
    }
}

struct Point {
    let x: Int
    let y: Int

    static func +(lhs: Point, rhs: Point) -> Point {
        return Point(x: lhs.x + rhs.x, y: lhs.y + rhs.y)
    }
}

func readInput(filepath: String) {
    guard let _ = freopen(filepath, "r", stdin) else {
        exit(1)
    }

    var readingGrid = true
    
    while let line = readLine() {
        if line.isEmpty {
            readingGrid = false
            continue
        }

        let row = line.map { String($0) }
        if readingGrid {
            grid.append(row)
        } else {
            path = row
        }
    }
}

func onboard(_ p: Point) -> Bool {
    if p.y < 0 || p.y >= grid.count || p.x < 0 || p.x >= grid[p.y].count || grid[p.y][p.x] == " " {
        return false
    }
    return true
}

func firstNonSpace(from: Point, facing: Orientation) -> Point {
    var destination = from
    destination = destination + facing.delta

    // TODO: haven't implemented wrapping around ...
    while !onboard(destination) {
        destination = destination + facing.delta
    }
    
    return destination
}

func getCommand() -> Command? {
    if pptr >= path.count { return nil }
    let c = path[pptr]
    pptr += 1
    if c == "R" { return .right }
    else if c == "L" { return .left }
    else {
        var val = Int(c)!

        while pptr < path.count && (path[pptr] != "L" && path[pptr] != "R") {
            val *= 10
            val += Int(path[pptr])!
            pptr += 1
        }
        
        return .step(val)
    }
}

func part1() {
    var current = firstNonSpace(from: Point(x: 0, y: 0), facing: .right)
    var orientation: Orientation = .right

    while let command = getCommand() {
            print(command)
        switch command {
        case .step(let steps):
            for i in 0 ..< steps {
                print("stepping \(i)")
                let ns = firstNonSpace(from: current, facing: orientation)
                if grid[ns.y][ns.x] != "#" { current = ns }
            }
        case .left:
            orientation = orientation.left()
        case .right:
            orientation = orientation.right()
        }
    }

    print(current, orientation)
}

if CommandLine.arguments.count > 1 {
    readInput(filepath: CommandLine.arguments[1])
    part1()
} else {
    print("usage: d22.swift <filename>")
}
