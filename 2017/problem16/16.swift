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

enum Command: Character {
    case spin = "s"
    case exchange = "x"
    case partner = "p"

    func execute(arg1: Any, arg2: Any, programs: inout [String]) {
        switch self {
        case .spin:
            let n = arg1 as! Int
            _spin(n: n, &programs)
        case .exchange:
            let a = arg1 as! Int
            let b = arg2 as! Int
            _exchange(a, b, &programs)
        case .partner:
            let m = arg1 as! String
            let n = arg2 as! String
            _partner(m, n, &programs)
        }
    }
}

func parse(_ line: String) {
    var line = line
    let command = Command(rawValue: line.removeFirst())!
    switch command {
    case .spin:
        let arg1 = Int(line)!
        command.execute(arg1: arg1 as Any, arg2: 0, programs: &programs)
    case .exchange: 
        let slash = line.index(of: "/")!
        let a = Int(line.prefix(through: line.index(slash, offsetBy: -1)))
        let b = Int(line.suffix(from: line.index(slash, offsetBy: 1)))
        command.execute(arg1: a as Any, arg2: b as Any, programs: &programs)
    case .partner:
        let slash = line.index(of: "/")!
        let a = String(line.prefix(through: line.index(slash, offsetBy: -1)))
        let b = String(line.suffix(from: line.index(slash, offsetBy: 1)))
        command.execute(arg1: a as Any, arg2: b as Any, programs: &programs)
    }
}

var programs: [String] = ["a", "b", "c", "d",
                          "e", "f", "g", "h",
                          "i", "j", "k", "l",
                          "m", "n", "o", "p"]

var lines = data.split(separator: ",")
_ = lines.map { parse(String($0).filter { $0 != "\n" }) }

print(programs.joined())

