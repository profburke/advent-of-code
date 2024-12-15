import Foundation

var commands: [Command] = []
while true {
    guard let line = readLine(strippingNewline: true) else { break }
    let parts = line.split(separator: " ")
    let command = Command(d: Direction(rawValue: String(parts[0]))!, n: Int(parts[1])!)
    commands.append(command)
}

var coords: [Coords] = []
coords.append(Coords(c1: 0, c2: 0))

var current = coords[0]
var minRow = Int.max
var maxRow = Int.min
var minCol = Int.max
var maxCol = Int.min
commands.forEach { command in
                   for _ in 0..<command.n {
                       current += command.d.delta
                       coords.append(current)
                       if current.c1 < minRow { minRow = current.c1 }
                       if current.c1 >= maxRow { maxRow = current.c1 }
                       if current.c2 < minCol { minCol = current.c2 }
                       if current.c2 >= maxCol { maxCol = current.c2 }
                   }
}

print(minRow, maxRow, minCol, maxCol)

struct Coords: CustomStringConvertible, Equatable {
    let c1: Int
    let c2: Int

    var description: String {
        "(\(c1), \(c2))"
    }

    static func +(lhs: Coords, rhs: Coords) -> Coords {
        Coords(c1: lhs.c1 + rhs.c1, c2: lhs.c2 + rhs.c2)
    }

    static func +=(lhs: inout Coords, rhs: Coords) {
        lhs = Coords(c1: lhs.c1 + rhs.c1, c2: lhs.c2 + rhs.c2)
    }
}

struct Command: CustomStringConvertible {
    let d: Direction
    let n: Int
    // add the color later

    var description: String {
        "\(d)-\(n)"
    }
}

enum Direction: String, CustomStringConvertible {
    case down = "D"
    case left = "L"
    case right = "R"
    case up = "U"

    var delta: Coords {
        switch self {
        case .up:
            return Coords(c1: -1, c2: 0)
        case .right:
            return Coords(c1: 0, c2: 1)
        case .down:
            return Coords(c1: 1, c2: 0)
        case .left:
            return Coords(c1: 0, c2: -1)
        }
    }
    
    var description: String {
        self.rawValue
    }
}
