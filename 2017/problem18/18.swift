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

typealias Register = Character
var registers: [Register : Int] = [:]
var lastFrequency = Int.min
var pc = 0

enum Operand {
    case register(Character)
    case value(Int)
}

enum Command {
    case sound(Operand)
    case set(Register, Operand)
    case add(Register, Operand)
    case multiply(Register, Operand)
    case modulo(Register, Operand)
    case recover(Operand)
    case jump(Operand, Operand)

    static func from(opcode: String, operand1: Operand, operand2: Operand?) -> Command? {
        switch (opcode, operand1, operand2) {
        case ("snd", let op, _):
            return Command.sound(op)
        case ("set", .register(let r), let .some(op)):
            return Command.set(r, op)
        case ("add", .register(let r), let .some(op)):
            return Command.add(r, op)
        case ("mul", .register(let r), let .some(op)):
            return Command.multiply(r, op)
        case ("mod", .register(let r), let .some(op)):
            return Command.modulo(r, op)
        case ("rcv", let op, _):
            return Command.recover(op)
        case ("jgz", let op1, let .some(op2)):
            return Command.jump(op1, op2)
        default:
            return nil
        }
    }

    func execute() {
        switch self {
        case .sound(let op):
            switch op {
            case .register(let r):
                lastFrequency = registers[r] ?? 0
            case .value(let v):
                lastFrequency = v
            }
        case .set(let r, let op):
            switch op {
            case .register(let r2):
                registers[r] = registers[r2] ?? 0
            case .value(let v):
                registers[r] = v
            }
        case .add(let r, let op):
            switch op {
            case .register(let r2):
                registers[r] = (registers[r] ?? 0) + (registers[r2] ?? 0)
            case .value(let v):
                registers[r] = (registers[r] ?? 0) + v
            }
        case .multiply(let r, let op):
            switch op {
            case .register(let r2):
                registers[r] = (registers[r] ?? 0) * (registers[r2] ?? 0)
            case .value(let v):
                registers[r] = (registers[r] ?? 0) * v
            }
        case .modulo(let r, let op):
            switch op {
            case .register(let r2):
                registers[r] = (registers[r] ?? 0) % (registers[r2] ?? 0)
            case .value(let v):
                registers[r] = (registers[r] ?? 0) % v
            }
        case .recover(let op):
            switch op {
            case .register(let r):
                if (registers[r] ?? 0) != 0 {
                    print(lastFrequency)
                    exit(0)
                }
            case .value(let v):
                if v != 0 {
                    print(lastFrequency)
                    exit(0)
                }
            }
        case .jump(let op1, let op2):
            // - 1 from the offset because at the end of the "cycle" we increment the pc by 1
            switch (op1, op2) {
            case (.register(let r1), .register(let r2)):
                if (registers[r1] ?? 0) > 0 {
                    pc += ((registers[r2] ?? 0) - 1)
                }
            case (.register(let r), .value(let v)):
                if (registers[r] ?? 0) > 0 {
                    pc += (v - 1)
                }
            case (.value(let v), .register(let r)):
                if v > 0 {
                    pc += ((registers[r] ?? 0) - 1)
                }
            case (.value(let v1), .value(let v2)):
                if v1 > 0 {
                    pc += (v2 - 1)
                }
            }
        }
    }
}

var commands: [Command] = []
func parse(_ line: String) {
    let components = line.split(separator: " ")
    let op1: Operand
    if let val = Int(components[1]) {
        op1 = .value(val)
    } else {
        op1 = .register(Character(String(components[1])))
    }
    let op2: Operand?
    if components.count < 3 {
        op2 = nil
    } else if let val = Int(components[2]) {
        op2 = .value(val)
    } else {
        op2 = .register(Character(String(components[2])))
    }
    if let command = Command.from(opcode: String(components[0]),
                               operand1: op1,
                               operand2: op2) {
        commands.append(command)
    } else {
        print("Error with '\(line)'")
        exit(1)
    }
}

var lines = data.split(separator: "\n")
_ = lines.map { parse(String($0)) }

while pc > -1 && pc < commands.count {
    commands[pc].execute()
    pc += 1
}

