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
    var minVal = Int.max
    var maxVal = Int.min
    var currentVal = 0
    while scanner.scanInt(&currentVal) {
        if currentVal < minVal { minVal = currentVal }
        if currentVal > maxVal { maxVal = currentVal }
    }
    return maxVal - minVal
}

let lines = data.split(separator: "\n")
print(lines.reduce(0) { $0 + process(String($1)) })
