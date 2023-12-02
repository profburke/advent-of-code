#!/usr/bin/env swift

import Darwin

struct Point: Hashable {
    let x: Int
    let y: Int

    static func +(lhs: Point, rhs: Point) -> Point {
        return Point(x: lhs.x + rhs.x, y: lhs.y + rhs.y)
    }
}

struct StepData {
    let nSteps: Int
    let p: Point
}

var grid: [[Int]] = []
var s = Point(x: 0, y: 0)
var e = Point(x: 0, y: 0)
let aAscii = Character("a").asciiValue!

var aPoints: [Point] = []

func readGrid(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Can't open \(path)")
        return
    }

    var row = 0
    while let line = readLine() {
        var r: [Int] = []

        var col = 0
        line.forEach { c in
                       if c == Character("S") {
                           s = Point(x: col, y: row)
                           r.append(0)
                           aPoints.append(s)
                       } else if c == Character("E") {
                           e = Point(x: col, y: row)
                           r.append(25)
                       } else {
                           let val = Int(c.asciiValue! - aAscii)
                           r.append(val)
                           if val == 0 {
                               aPoints.append(Point(x: col, y: row))
                           }
                       }

                       col += 1
        }

        grid.append(r)
        row += 1
    }

    //grid.forEach { print($0) }
}

func analyze(startingFrom from: Point) -> Int {
    var seen = Set<Point>()
    var queue: [StepData] = []
    // queue.append(StepData(nSteps: 0, p: s))
    // seen.insert(s)
    queue.append(StepData(nSteps: 0, p: from))
    seen.insert(from)

    while queue.count > 0 {
        //print("Queue length is \(queue.count)")
        let sd = queue.removeFirst()
        //seen.insert(sd.p)
        
        if sd.p == e {
            // print("Min steps is \(sd.nSteps)")
            // break
            return sd.nSteps
        }
        
        let currentHeight = grid[sd.p.y][sd.p.x]

        // check each nbr and if ok, add to queue
        for delta in [Point(x: 0, y: -1), Point(x: 0, y: 1), Point(x: -1, y: 0), Point(x: 1, y: 0)] {
            let np = sd.p + delta
            if !seen.contains(np)
            && np.x >= 0 && np.y >= 0
            && np.x < grid[0].count && np.y < grid.count
            && grid[np.y][np.x] <= currentHeight + 1 {
                seen.insert(np)
                let nsd = StepData(nSteps: sd.nSteps + 1, p: np)
                queue.append(nsd)
            }
        }
    }

    return -100
}

func part1() {
    let steps = analyze(startingFrom: s)
    print("Min steps is \(steps)")
}

func part2() {
    var best = Int.max
    aPoints.forEach { p in
                      let steps = analyze(startingFrom: p)
                      if steps > 0 && steps < best { best = steps }
    }
    print("Scenic best is \(best)")
}

if CommandLine.arguments.count > 1 {
    readGrid(path: CommandLine.arguments[1])
    part1()
    part2()
} else {
    print("usage: ./d12.swift <input>")
}
