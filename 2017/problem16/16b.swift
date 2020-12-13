#!/usr/bin/swift
import Foundation

/*

Not sure why this doesn't work:

Take the arrangement at the start and the arrangement at the end of one "cycle"
and figure out a single permutation that does that transformation. Then just repeat that
rather than doing the individual dance steps.

But it doesn't work...so...

oh! wait. figured it out. It's because "pair" is based on program names,
not positions



NOTE: "naive" approach takes forever....
going to rewrite so it doesn't keep "parsing" every time...


*/

if CommandLine.argc != 2 {
     print("Usage: \(CommandLine.arguments[0]) <inputfile>")
     exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
     print("Couldn't open \(fileName)")
     exit(1)
}


func step(_ programs: inout [String]) {
    let temp = programs[programs.count - 1]
    for i in (1..<programs.count).reversed() {
        programs[i] = programs[i - 1]
    }
    programs[0] = temp
}

func _spin(n: Int, _ programs: inout [String]) {
    for _ in 0..<n {
        step(&programs)
    }
}

func _exchange(_ a: Int, _ b: Int, _ programs: inout [String]) {
    (programs[b], programs[a]) = (programs[a], programs[b])
}

func _partner(_ m: String, _ n: String, _ programs: inout [String]) {
    let a = programs.index(of: m)!
    let b = programs.index(of: n)!
    _exchange(a, b, &programs)
}

enum Operand {
    case index(Int)
    case name(String)
}

enum Command {
    case spin(Int)
    case exchange(Int, Int)
    case partner(String, String)

    static func from(code: String, op1: Operand, op2: Operand?) -> Command? {
        switch (code, op1, op2) {
        case ("s", .index(let n), _):
            return Command.spin(n)
        case ("x", .index(let a), .some(.index(let b))):
            return Command.exchange(a, b)
        case ("p", .name(let m), .some(.name(let n))):
            return Command.partner(m, n)
        default:
            return nil
        }
    }
    
    func execute(programs: inout [String]) {
        switch self {
        case .spin(let n):
            _spin(n: n, &programs)
        case .exchange(let a, let b):
            _exchange(a, b, &programs)
        case .partner(let m, let n):
            _partner(m, n, &programs)
        }
    }
}

var commands: [Command] = []
func parse(_ line: String) {
    var line = line
    let code = line.removeFirst()
    let op1: Operand
    let op2: Operand?
    if let slashdex = line.index(of: "/") {
        let first = line.prefix(through: line.index(slashdex, offsetBy: -1))
        let second = line.suffix(from: line.index(slashdex, offsetBy: 1))
        if let val = Int(first) {
            op1 = .index(val)
        } else {
            op1 = .name(String(first))
        }
        if let val = Int(second) {
            op2 = .index(val)
        } else {
            op2 = .name(String(second))
        }
    } else {
        op1 = .index(Int(line)!)
        op2 = nil
    }
    let command = Command.from(code: String(code), op1: op1, op2: op2)!
    commands.append(command)
}

var programs: [String] = ["a", "b", "c", "d",
                          "e", "f", "g", "h",
                          "i", "j", "k", "l",
                          "m", "n", "o", "p"]

var seen = Set<String>()
var recordings: [String] = []
var lines = data.split(separator: ",")
let start = Date()
_ = lines.map { parse(String($0).filter { $0 != "\n" }) }
let fingerprint = programs.joined()
seen.insert(fingerprint)
recordings.append(fingerprint)
var modulus = 0
for step in 1..<100 {
    _ = commands.map { $0.execute(programs: &programs) }
    let fingerprint = programs.joined()
    if seen.contains(fingerprint) {
        modulus = step
        print("\(step): Seen it: ", terminator: "")
        print(fingerprint)
        break
    }
    print("\(step): \(fingerprint)")
    seen.insert(fingerprint)
    recordings.append(fingerprint)
}
let finish = Date()
print(finish.timeIntervalSince(start))
print(modulus)
let index = 1_000_000_000 % modulus
print(index)
print(recordings[index])

