#!/usr/bin/env swift

import Darwin

struct Position: Hashable {
    let x: Int
    let y: Int

    static func +(_ lhs: Position, _ rhs: Position) -> Position {
        return Position(x: lhs.x + rhs.x, y: lhs.y + rhs.y)
    }

    static func -(_ lhs: Position, _ rhs: Position) -> Position {
        return Position(x: lhs.x - rhs.x, y: lhs.y - rhs.y)
    }
        
    func distance(to other: Position) -> Int {
        return max(abs(x - other.x), abs(y - other.y))
    }

    func inLine(with other: Position) -> Bool {
        return x == other.x || y == other.y
    }
    
    static let deltaR = Position(x: 1, y: 0)
    static let deltaL = Position(x: -1, y: 0)
    static let deltaU = Position(x: 0, y: 1)
    static let deltaD = Position(x: 0, y: -1)
}

var tailPositions = Set<Position>()
var headP = Position(x: 0, y: 0)
var tailP = Position(x: 0, y: 0)

func moveTail(delta: Position) {
    if headP.distance(to: tailP) <= 1 { return }
    if headP.inLine(with: tailP) { tailP = tailP + delta; return }

    var dx = 0
    var dy = 0

    if headP.x > tailP.x {
        dx = 1
    } else { // we've determined above they can't be equal
        dx = -1
    }

    if headP.y > tailP.y {
        dy = 1
    } else {
        dy = -1
    }
    
    tailP = tailP + Position(x: dx, y: dy)
}

func step(_ direction: String, _ steps: Int) {
    let delta: Position
    
    switch direction {
    case "R":
        delta = Position.deltaR
    case "L":
        delta = Position.deltaL
    case "U":
        delta = Position.deltaU
    case "D":
        delta = Position.deltaD
    default:
        print("'\(direction)' is not a vaild direction")
        exit(1)
    }
    
    for _ in 0 ..< steps {
        headP = headP + delta
        moveTail(delta: delta)
        tailPositions.insert(tailP)
    }
}

func part1(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Can't open \(path)")
        return
    }

    tailPositions.insert(tailP)
    
    while let line = readLine() {
        let parts = line.split(separator: " ")
        let steps = Int(parts[1])!
        let direction = parts[0]

        step(String(direction), steps)
    }

    print("Tail at \(tailPositions.count) unique positions")
}

// index 0 = head; index 9 = tail
var longRope = Array<Position>(repeating: Position(x: 0, y: 0), count: 10)
let taildex = 9

func moveKnot(_ delta: Position, _ lowdex: Int, _ hidex: Int) {
    if longRope[lowdex].distance(to: longRope[hidex]) <= 1 { return }
    // if longRope[lowdex].inLine(with: longRope[hidex]) { longRope[hidex] = longRope[hidex] + delta; return }

    // // note delta is direction head moved in
    // // but we really need to move in direction previous knot
    // // moved in
    
    // var dx = 0
    // var dy = 0

    // if longRope[lowdex].x > longRope[hidex].x {
    //     dx = 1
    // } else { // we've determined above they can't be equal
    //     dx = -1
    // }

    // if longRope[lowdex].y > longRope[hidex].y {
    //     dy = 1
    // } else {
    //     dy = -1
    // }
    
    longRope[hidex] = longRope[hidex] + Position(x: dx, y: dy)
}

func multistep(_ direction: String, _ steps: Int) {
    let delta: Position
    
    switch direction {
    case "R":
        delta = Position.deltaR
    case "L":
        delta = Position.deltaL
    case "U":
        delta = Position.deltaU
    case "D":
        delta = Position.deltaD
    default:
        print("'\(direction)' is not a vaild direction")
        exit(1)
    }
    
    for _ in 0 ..< steps {
        longRope[0] = longRope[0] + delta
        for i in 1 ..< longRope.count {
            moveKnot(delta, i - 1, i)
        }

        // visualize(longRope)
        
        tailPositions.insert(longRope[taildex])
    }
}

func visualize(_ poslist: [Position], indices: Bool = true) {
    var grid: [[String]] = []

    for _ in 0 ..< 21 {
        grid.append(Array<String>(repeating: ".", count: 35))
    }

    grid[15][11] = "s"

    for i in (0 ..< poslist.count).reversed() {
        let label: String
        if indices {
            label = (i == 0) ? "H" : "\(i)"
        } else {
            label = "#"
        }
        let y =  -1 * poslist[i].y + 15
        let x = poslist[i].x + 11
        grid[y][x] = label
    }
    
    grid.forEach { row in
                   print(row.joined())
                   print()
    }

    print()
    print("==========================")
    print()
}

func part2(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Can't open \(path)")
        return
    }

    tailPositions.insert(longRope[taildex])

    // visualize(longRope)
    
    while let line = readLine() {
        let parts = line.split(separator: " ")
        let steps = Int(parts[1])!
        let direction = parts[0]

        multistep(String(direction), steps)
        //visualize(longRope)
    }

    // visualize(longRope)
    print("Tail at \(tailPositions.count) unique positions")

    // visualize(Array(tailPositions), indices: false)
}

part1(path: CommandLine.arguments[1])
tailPositions = Set<Position>()
part2(path: CommandLine.arguments[1])
