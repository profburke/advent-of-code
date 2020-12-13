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

func process(_ line: String) -> Int {
    let scanner = Scanner(string: line)
    var vals = [Int]()
    var currentVal = 0
    while scanner.scanInt(&currentVal) {
        vals.append(currentVal)
    }
    for i in 0..<vals.count {
        for j in 0..<vals.count {
            if i == j { continue }
            let a = vals[i]
            let b = vals[j]
            let quotient = a/b
            if quotient * b == a { return quotient }
        }
    }
    return Int.min // should never get here
}

let lines = data.split(separator: "\n")
print(lines.reduce(0) { $0 + process(String($1)) })
