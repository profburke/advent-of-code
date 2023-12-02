#!/usr/bin/env swift

import Darwin

enum Operation {
    case add
    case multiply
}

enum Operand: Equatable {
    case literal(Int)
    case old
}

enum ReadState {
    case monkeyNumber
    case items
    case operation
    case test
    case trueBranch
    case falseBranch
    case blank
}

class Monkey: CustomStringConvertible {
    var monkeyNumber = 0
    var items: [Int] = []
    var test = 0
    var operation: Operation = .add
    var operandA: Operand = .old
    var operandB: Operand = .old
    var trueDestination = 0
    var falseDestination = 0
    var inspections = 0

    // static let numberRegex = /Monkey (\d+)/
    // Regex requires macOS 13+  (ugh!)

    var description: String {
        return "Monkey_\(monkeyNumber) {\(items)} Test: \(test) ? \(trueDestination) : \(falseDestination) new = \(operandA) \(operation) \(operandB)"
    }
}

var monkeys: [Monkey] = []

func readMonkeys(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Could not read \(path)")
        return
    }

    var readstate: ReadState = .monkeyNumber

    var monkey = Monkey()
    while let line = readLine() {
        switch readstate {
        case .monkeyNumber:
            monkey.monkeyNumber = Int(String(line
                                             .dropLast(1)
                                             .dropFirst(7)))!
            readstate = .items

        case .items:
            monkey.items = line
            .dropFirst(17)
            .split(separator: ",")
            .map { Int($0.dropFirst(1))! }
            readstate = .operation
           
        case .operation:
            let parts = line.dropFirst(19).split(separator: " ").map { String.init($0) }
            monkey.operandA = (parts[0] == "old") ? .old : .literal(Int(parts[0])!)
            monkey.operandB = (parts[2] == "old") ? .old : .literal(Int(parts[2])!)
            monkey.operation = (parts[1] == "+") ? .add : .multiply
            readstate = .test

        case .test:
            monkey.test = Int(String(line
                                     .dropFirst(21)))!
            readstate = .trueBranch

        case .trueBranch:
            monkey.trueDestination = Int(String(line
                                                .dropFirst(29)))!
            readstate = .falseBranch

        case .falseBranch:
            monkey.falseDestination = Int(String(line
                                                 .dropFirst(30)))!
            readstate = .blank

        case .blank:
            monkeys.append(monkey)
            monkey = Monkey()
            readstate = .monkeyNumber
        }
    }

    // monkeys.forEach { print($0) }
}

func jdi(_ monkey: Monkey) {
    print("Monkey \(monkey.monkeyNumber)'s turn")
    monkey.items.forEach { item in
                           print("Inspecting item \(item)")
                           monkey.inspections += 1
                           let operandA: Int
                           if case let .literal(val) = monkey.operandA {
                               operandA = val
                           } else {
                               operandA = item
                           }
                           let operandB: Int
                           if case let .literal(val) = monkey.operandB {
                               operandB = val
                           } else {
                               operandB = item
                           }

                           print("A \(operandA) B \(operandB)")

                           if monkey.operation == .add {
                               // a + b is divisble by number if both a and b are
                               if (operandA % monkey.test == 0) && (operandB % monkey.test == 0) {
                                   monkeys[monkey.trueDestination].items.append(newItem)
                               } else {
                                   monkeys[monkey.falseDestination].items.append(newItem)
                               }
                               
                           } else {
                               // a * b is divisble bu number if either a or b are
                               if (operandA % monkey.test == 0) || (operandB % monkey.test == 0) {
                                   monkeys[monkey.trueDestination].items.append(newItem)
                               } else {
                                   monkeys[monkey.falseDestination].items.append(newItem)
                               }
                           }
                           
                           // var newItem = 0
                           // if monkey.operation == .add {
                           //     print("+")
                           //     newItem = operandA + operandB
                           // } else {
                           //     print("*")
                           //     newItem = operandA * operandB
                           // }
                           // print("Worry level goes to \(newItem)")

                           // For part 2, we don't do this ... newItem = Int(Double(newItem)/3.0)
                           // print("Worry level drops to \(newItem)")

                           // if newItem % monkey.test == 0 {
                           //     print("Tosses to monkey \(monkey.trueDestination)")
                           //     monkeys[monkey.trueDestination].items.append(newItem)
                           //     //print("Tossed")
                           // } else {
                           //     print("Tosses to monkey \(monkey.falseDestination)")
                           //     monkeys[monkey.falseDestination].items.append(newItem)
                           //     //print("Tossed")
                           // }
                           
}

    monkey.items = []
}

func part1() {
    var round = 0
    while round < 10_000 {
        print("Round \(round)")
        monkeys.forEach { monkey in jdi(monkey) }
        round += 1
    }

    monkeys.sort { $0.inspections < $1.inspections }

    let most = monkeys[monkeys.count - 1].inspections
    let secondMost = monkeys[monkeys.count - 2].inspections
    print("Level of monkey business -> \(most) * \(secondMost) = \(most * secondMost)")
}

if CommandLine.arguments.count > 1 {
    readMonkeys(path: CommandLine.arguments[1])
    part1()
} else {
    print("usage: ./d11.swift <inputfile>")
}
