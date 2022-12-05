#!/usr/bin/env swift

import Darwin

// rock - A, X
// paper - B, Y
// scissors - C, Z

enum Choice {
    case rock
    case paper
    case scissors
    
    func score() -> Int {
        switch self {
        case .rock:
            return 1
        case .paper:
            return 2
        case .scissors:
            return 3
        }
    }

    init(code: String) {
        if code == "A" || code == "X" {
            self = .rock
        } else if code == "B" || code == "Y" {
            self = .paper
        } else {
            self = .scissors // ok, _ought_ to be a failable initializer and return nil of not "C" or "Z" ...
        } 
    }
}

enum Outcome {
    case lose
    case draw
    case win
    
    func score() -> Int {
        switch self {
        case .lose:
            return 0
        case .draw:
            return 3
        case .win:
            return 6
        }
    }

    init(code: String) {
        if code == "X" {
            self = .lose
        } else if code == "Y" {
            self = .draw
        } else {
            self = .win // see previous comment re failable initializers
        }
    }
}

func roundResultForYou(l: Choice, r: Choice) -> Outcome {
    switch (l,r) {
    case (.scissors, .rock):
        return .win
    case (.paper, .rock):
        return .lose
    case (.paper, .scissors):
        return .win
    case (.rock, .scissors):
        return .lose
    case (.rock, .paper):
        return .win
    case (.scissors, .paper):
        return .lose
    default:
        return .draw
    }
}

func part1(path: String) {
    var score = 0
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        let parts = line.split(separator: " ")
        let elfMove = Choice(code: String(parts[0]))
        let myMove = Choice(code: String(parts[1]))
        score += myMove.score() + roundResultForYou(l: elfMove, r: myMove).score()
    }

    print("Score is \(score)")
}

func necessaryChoice(elfMove: Choice, outcome: Outcome) -> Choice {
    switch elfMove {
    case .rock:
        switch outcome {
        case .lose:
            return .scissors
        case .draw:
            return .rock
        case .win:
            return .paper
        }
    case .paper:
        switch outcome {
        case .lose:
            return .rock
        case .draw:
            return .paper
        case .win:
            return .scissors
        }
    case .scissors:
        switch outcome {
        case .lose:
            return .paper
        case .draw:
            return .scissors
        case .win:
            return .rock
        }
    }
}

func part2(path: String) {
    var score = 0
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        let parts = line.split(separator: " ")
        let elfMove = Choice(code: String(parts[0]))
        let requiredOutcome = Outcome(code: String(parts[1]))
        let myMove = necessaryChoice(elfMove: elfMove, outcome: requiredOutcome)
        score += myMove.score() + roundResultForYou(l: elfMove, r: myMove).score()
    }

    print("Score is \(score)")
}

part1(path: CommandLine.arguments[1])
part2(path: CommandLine.arguments[1])

