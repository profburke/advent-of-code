#!/usr/bin/env swift

import Darwin
import class Foundation.NSString

struct Point: Hashable, CustomStringConvertible {
    let x: Int
    let y: Int

    var description: String {
        return "(\(x), \(y))"
    }

    static func +(lhs: Point, rhs: Point) -> Point {
        return Point(x: lhs.x + rhs.x, y: lhs.y + rhs.y)
    }
}

enum Element {
    case air
    case rock
    case sand
}

var grid: [Point : Element] = [:]
var minX = 0
var maxX = 0
var minY = 0
var maxY = 0

func readGrid(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Could not read \(path)")
        exit(1)
    }

    while let line = readLine() {
        let items = line.replacingOccurrences(of: " -> ", with: ":")
        .split(separator: ":").map { String($0) }

        let points = items.map { d in
                                 let coords = d.split(separator: ",")
                                 .map { Int(String($0))! }
                                 return Point(x: coords[0], y: coords[1])
        }

        var current = points[0]
        grid[current] = .rock

        points.dropFirst().forEach { p in
                                     let dx: Int
                                     if p.x == current.x {
                                         dx = 0
                                     } else if p.x > current.x {
                                         dx = 1
                                     } else {
                                         dx = -1
                                     }
                                     
                                     let dy: Int
                                     if p.y == current.y {
                                         dy = 0
                                     } else if p.y > current.y {
                                         dy = 1
                                     } else {
                                         dy = -1
                                     }
                                     
                                     let delta = Point(x: dx, y: dy)

                                     while current != p {
                                         current = current + delta
                                         grid[current] = .rock
                                     }
        }
    }
    minX = grid.keys.map { $0.x }.min()!
    minY = grid.keys.map { $0.y }.min()!
    maxX = grid.keys.map { $0.x }.max()!
    maxY = grid.keys.map { $0.y }.max()!
}

func visualize() {
    for y in (minY - 5) ... (maxY + 5) {
        for x in (minX - 5) ... (maxX + 5) {
            let e = grid[Point(x: x, y: y), default: .air]
            let s: String
            switch e {
            case .rock:
                s = "#"
            case .sand:
                s  = "o"
            case .air:
                s = "."
            }
            print(s, terminator: "")
        }
        print()
    }
}

func move(from s: Point) -> Point {
    let candidates = [Point(x: s.x, y: s.y + 1), Point(x: s.x - 1, y: s.y + 1), Point(x: s.x + 1, y: s.y + 1)]
    for p in candidates {
        if grid[p, default: .air] == .air {
            return p
        }
    }
    return s
}

func removeSand() {
    grid.keys.forEach { k in
                        if grid[k] == .sand {
                            // would deleting mess the iteration up?
                            grid[k] = .air
                        }
    }
}

func buildFloor() {
    maxY += 2
    
    for x in (minX - 200) ... (maxX + 200) {
        grid[Point(x: x, y: maxY)] = .rock
    }
}

func dropSand() {
    let dropPoint = Point(x: 500, y: 0)
    var grains = 0
    
    while true {
        // print("dropping")
        if grid[dropPoint, default: .air] == .sand {
            print("dropPoint is sand")
            break
        }
        var currentPoint = dropPoint
        var movePoint = move(from: currentPoint)
        while movePoint != currentPoint {
            //print("moving to \(movePoint)")
            currentPoint = movePoint
            movePoint = move(from: currentPoint)
            if movePoint.y > maxY { break }
        }
        if movePoint.y > maxY { break }
        grid[movePoint] = .sand
        grains += 1
        //visualize()
    }
    
    print("\(grains) stabilized")
}

if CommandLine.arguments.count > 1 {
    readGrid(path: CommandLine.arguments[1])
    visualize()
    dropSand()
    print(); print()
    visualize()
    removeSand()
    buildFloor()
    dropSand()
    print(); print()
    visualize()
} else {
    print("usage: ./d14.swift <filename>")
}
