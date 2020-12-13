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

var location = 0
var steps = 0
var jumps = data.split(separator: "\n").map { Int($0) ?? 0 }

while location > -1 && location < jumps.count {
    let offset = jumps[location]
    jumps[location] += 1
    steps += 1
    location += offset
    
}
print(steps)
