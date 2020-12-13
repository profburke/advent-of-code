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

var maxDepth = 0
var layerData: [Int : Int] = [:]
func parse(_ line: String) {
    let components = line.split(separator: " ")
    let depth = Int(components[0].filter {  $0 != ":" })!
    let range = Int(components[1])!
    layerData[depth] = range
    if depth > maxDepth { maxDepth = depth }
}

var lines = data.split(separator: "\n")
_ = lines.map { parse(String($0)) }

var firewall = Array<Int>(repeating: 0, count: maxDepth + 1)
for i in 0..<(maxDepth + 1) {
    firewall[i] = layerData[i] ?? 0
}

for delay in 0..<10000000 {
    var spotted = false
    var severity = 0

    for i in 0..<(maxDepth + 1) {
        let range = firewall[i]
        if (range != 0) && (((i+delay)%(2 * (range - 1))) == 0) {
            spotted = true
            severity += (i*range)
        }
    }

    if !spotted {
        print("\(delay): \(severity)")
    }
}

