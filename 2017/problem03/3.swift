#!/usr/bin/swift

let target = 325489

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

func coordinates(for target: Int) -> (Int, Int) {
    var x = 0
    var y = 0
    var currentVal = 1
    var currentDirection = Direction.east
    var currentStepsInDirection = 0
    var maxStepsInDirection = 1
    var repetition = 0
    while currentVal < target {
        currentVal += 1
        x += currentDirection.dx()
        y += currentDirection.dy()
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
    return (x, y)
}

let (x, y) = coordinates(for: target)
print(abs(x) + abs(y))
