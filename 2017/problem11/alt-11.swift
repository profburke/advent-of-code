#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
    print("Couldn't open \(fileName)")
    exit(1)
}

enum Step: String {
    case nw
    case n
    case ne
    case se
    case s
    case sw

    var opposite: Step {
        switch self {
        case .nw:
            return .se
        case .n:
            return .s
        case .ne:
            return .sw
        case .se:
            return .nw
        case .s:
            return .n
        case .sw:
            return .ne
        }
    }
    
    func inbetween(_ other: Step) -> Step? {
        switch (self, other) {
        case (.ne, .nw), (.nw, .ne):
            return .n
        case (.ne, .s), (.s, .ne):
            return .se
        case (.se, .n), (.n, .se):
            return .ne
        case (.se, .sw), (.sw, .se):
            return .s
        case (.s, .nw), (.nw, .s):
            return .sw
        case (.sw, .n), (.n, .sw):
            return .nw
        default:
            return nil
        }
    }
}


func stripOpposites(from data: [Int]) -> ([Int], Bool) {
    var result = data
    var changes = false

    return (result, changes)
}

var steps = data.split(separator: ",").dropLast().map { Step(rawValue:
                                                             String($0))! }


var done = false
var cycles = 0

while !done {
    var base = 0
    var madeAdjustment = false

    while base < steps.count {
        let first = steps[base]
        for current in base..<steps.count {
            let second = steps[current]
            if first.opposite == second {
                madeAdjustment = true
                steps.remove(at: current)
                break
            }
        }
        base += 1
    }

    base = 0
    
    while base < (steps.count - 1) {
        let first = steps[base]
        for current in (base + 1) ..< (steps.count - 1) {
            if let result = first.inbetween(steps[current]) {
                madeAdjustment = true
                steps[base] = result
                steps.remove(at: current)
                break
            }
        }
        base += 1
    }

    if !madeAdjustment { done = false }
    print(".", terminator: "")
    if cycles%40 == 0 { print("") }
    cycles += 1
}

print(steps.count)

