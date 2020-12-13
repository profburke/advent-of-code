#!/usr/bin/swift

let target = 325489

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

var cellValues: [Coordinates : Int] = [:]
cellValues[Coordinates(x: 0, y: 0)] = 1

func sumNeighbors(coords: Coordinates) -> Int {
    var result = 0
    for dx in -1...1 {
        for dy in -1...1 {
            result += cellValues[Coordinates(x: coords.x + dx, y: coords.y + dy)] ?? 0
        }
    }
    return result
}

enum Direction {
    case east
    case north
    case west
    case south

    func dx() -> Int {
        switch self {
        case .east: return 1
        case .north: return 0
        case .west: return -1
        case .south: return 0
        }
    }
    
    func dy() -> Int {
        switch self {
        case .east: return 0
        case .north: return 1
        case .west: return 0
        case .south: return -1
        }
    }
    
    func next() -> Direction {
        switch self {
        case .east: return .north
        case .north: return .west
        case .west: return .south
        case .south: return .east
        }
    }
}

func largerValue(than target: Int) -> Int {
    var x = 0
    var y = 0
    var neighborSum = Int.min
    var currentDirection = Direction.east
    var currentStepsInDirection = 0
    var maxStepsInDirection = 1
    var repetition = 0
    while neighborSum < target {
        x += currentDirection.dx()
        y += currentDirection.dy()
        let coords = Coordinates(x: x, y: y)
        neighborSum = sumNeighbors(coords: coords)
        cellValues[coords] = neighborSum
        currentStepsInDirection += 1
        if currentStepsInDirection == maxStepsInDirection {
            currentDirection = currentDirection.next()
            currentStepsInDirection = 0
            repetition += 1
            if repetition == 2 {
                maxStepsInDirection += 1
                repetition = 0
            }
        }
    }
    return neighborSum
}

print(largerValue(than: target))
