import Foundation

var grid: [[Item]] = []

func readInput() {

    while true {
        guard let line = readLine(strippingNewline: true) else {
            break
        }
        
        var row: [Item] = []
        line.forEach { c in
                       row.append(Item(rawValue: String(c))!)
        }
        
        grid.append(row)
    }
}

func stepNorth() {
    let n = grid.count
    for i in 0..<n {
        for j in 0..<grid[i].count {
            if i < n - 1 {
                if grid[i][j] == .space && grid[i+1][j] == .roundRock {
                    grid[i][j] = .roundRock
                    grid[i+1][j] = .space
                }
            // } else {
            //     if grid[i][j] == .roundRock { grid[i][j] = .space }
            }
        }
    }
}

func printGrid() {
    grid.forEach { line in
                   line.forEach { c in print(c, terminator: "") }
                   print()
    }
}

func tiltNorth() {
    for _ in 0..<grid.count {
        stepNorth()
    }

    printGrid()
}

func calculateLoad() -> Int {
    var load = 0

    for (index, line) in grid.enumerated() {
        line.forEach { item in
                       if item == .roundRock { load += grid.count - index }
        }
    }
    
    return load
}

func part1() {
    readInput()
    tiltNorth()
    let load = calculateLoad()

    print("Part 1 - \(load)")
}

part1()

enum Item: String, CustomStringConvertible {
    case cubeRock = "#"
    case roundRock = "O"
    case space = "."

    var description: String {
        switch self {
        case .cubeRock: return "#"
        case .roundRock: return "O"
        case .space: return "."
        }
    }
}
