#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
     print("Usage: \(CommandLine.arguments[0]) <inputfile>")
     exit(1)
}

enum Cell: Equatable, CustomStringConvertible {
    case vertical
    case horizontal
    case corner
    case letter(Character)
    case blank

    init?(from filerep: Character) {
        switch filerep {
        case "|":
            self = .vertical
        case "-":
            self = .horizontal
        case "+":
            self = .corner
        case " ":
            self = .blank
        case let val where CharacterSet.letters.contains(val.unicodeScalars.first!):
            self = .letter(val)
        default:
            return nil
        }
    }

    var description: String {
        switch self {
        case .vertical: return "|"
        case .horizontal: return "-"
        case .blank: return "•"
        case .corner: return "+"
        case .letter(let a): return(String(a))
        }
    }
}

func ==(lhs: Cell, rhs: Cell) -> Bool {
    switch (lhs, rhs) {
    case (.vertical, .vertical):
        return true
    case (.horizontal, .horizontal):
        return true
    case (.corner, .corner):
        return true
    case (.blank, .blank):
        return true
    case (.letter(let a), .letter(let b)):
        return a == b
    default:
        return false
    }
}

struct Coordinates {
    let r: Int
    let c: Int
}

func +(lhs: Coordinates, rhs: Coordinates) -> Coordinates {
    return Coordinates(r: lhs.r + rhs.r , c: lhs.c + rhs.c)
}

enum Direction {
    case north
    case east
    case south
    case west

    var vertical: Bool {
        return self == .north || self == .south
    }

    var horizontal: Bool {
        return !self.vertical
    }
    
    var offset: Coordinates {
        switch self {
        case .north:
            return Coordinates(r: -1, c: 0)
        case .east:
            return Coordinates(r: 0, c: 1)
        case .south:
            return Coordinates(r: 1, c: 0)
        case .west:
            return Coordinates(r: 0, c: -1)
        }
    }
}


var position = Coordinates(r: 0, c: 0)

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
     print("Couldn't open \(fileName)")
     exit(1)
}

var steps = 0
var direction = Direction.south
var done = false
var letters: [String] = []
var lines = data.split(separator: "\n")
var grid: [[Cell]] = []
for line in lines {
    var row: [Cell] = []
    line.forEach { char in
                   let cell = Cell(from: char)!
                   row.append(cell)
    }
    grid.append(row)
}

func findEntrance() -> Coordinates? {
    for col in 0..<grid[0].count {
        if grid[0][col] == Cell.vertical {
            return Coordinates(r: 0, c: col)
        }
    }
    return nil
}

func onGrid(_ position: Coordinates) -> Bool {
    let r = position.r
    let c = position.c
    return r > -1 && r < grid.count && c > -1 && c < grid[0].count
}

func switchToHorizontal() {
    let r = position.r
    let c = position.c
    let newPos = Coordinates(r: r, c: c + 1)
    if onGrid(newPos) && grid[newPos.r][newPos.c] != .blank {
        direction = .east
    } else {
        direction = .west
    }
}

func switchToVertical() {
    let r = position.r
    let c = position.c
    let newPos = Coordinates(r: r + 1, c: c) 
    if onGrid(newPos) && grid[newPos.r][newPos.c] != .blank {
        direction = .south
    } else {
        direction = .north
    }
}

func changeDirection() {
    if direction.vertical {
        switchToHorizontal()
    } else {
        switchToVertical()
    }
}

position = findEntrance()!

while !done {
    if !onGrid(position) {
        print("Somehow we're off the grid")
        exit(1)
    }
    let cell = grid[position.r][position.c]
    switch cell {
    case .vertical, .horizontal:
        position = position + direction.offset
        steps += 1
    case .corner:
        changeDirection()
        position = position + direction.offset
        steps += 1
    case .blank:
        done = true
    case .letter(let a):
        letters.append(String(a))
        position = position + direction.offset
        steps += 1
    }
}

print(letters.joined())
print(steps)


