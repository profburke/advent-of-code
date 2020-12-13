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

var max = Int.min

enum Operation: String {
    case inc
    case dec

    func perform(_ initial: Int, _ value: Int) -> Int {
        switch self {
        case .inc:
            let result =  initial + value
            if result > max { max = result }
            return result
        case .dec:
            let result = initial - value
            if result > max { max = result }
            return result
        }
    }
}

enum Relop: String {
    case gt = ">"
    case lt = "<"
    case ge = ">="
    case eq = "=="
    case le = "<="
    case ne = "!="

    func perform(_ a: Int, _ b: Int) -> Bool {
        switch self {
        case .gt:
            return (a > b)
        case .lt:
            return (a < b)
        case .ge:
            return (a >= b)
        case .eq:
            return (a == b)
        case .le:
            return (a <= b)
        case .ne:
            return (a != b)
        }
    }
}

struct Command {
    let targetRegister: String
    let operation: Operation
    let value: Int
    let sourceRegister: String
    let relop: Relop
    let triggerValue: Int

    func execute(with registers: inout [String : Int]) {
        let control = relop.perform(registers[sourceRegister]!, triggerValue)
        if control {
            registers[targetRegister] = operation.perform(registers[targetRegister]!, value)
        }
    }
}

func parse(_ line: String) -> (String, Command) {
    let components = line.split(separator: " ")
    let targetRegister = String(components[0])
    let operation = Operation(rawValue: String(components[1]))!
    let value = Int(components[2])!
    let sourceRegister = String(components[4])
    let relop = Relop(rawValue: String(components[5]))!
    let triggerValue = Int(components[6])!
    let command = Command(targetRegister: targetRegister, operation: operation,
                          value: value, sourceRegister: sourceRegister, relop: relop,
                          triggerValue: triggerValue)
    return (targetRegister, command)
}

var lines = data.split(separator: "\n", maxSplits: Int.max,
                        omittingEmptySubsequences: true)

var registers: [String : Int] = [:]
var commands: [Command] = []

print("parsing...")
for line in lines {
    //print("Line: \(line)")
    let (register, command) = parse(String(line))
    registers[register] = 0
    commands.append(command)
}

print("executing...")
for command in commands {
    command.execute(with: &registers)
}

print(max)

