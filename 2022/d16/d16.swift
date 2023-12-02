#!/usr/bin/env swift

import Darwin
import class Foundation.NSString

struct ValveData: Hashable, CustomStringConvertible {
    let name: String
    let flow: Int
    let tunnels: [String]
    var open = false

    var description: String {
        return "<\(name): \(open)-\(flow) \(tunnels)>"
    }
}

var valves: [String:ValveData] = [:]
var currentLocation = "AA"

func readData(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    // I just do want to bother with regexes ...
    while let line = readLine() {
        let parts = line.replacingOccurrences(of: "Valve ", with: "")
        .replacingOccurrences(of: "has flow rate=", with: "")
        .replacingOccurrences(of: "; tunnels lead to valves", with: "")
        .replacingOccurrences(of: "; tunnel leads to valve", with: "")
        .split(separator: " ")
        .map { String($0) }

        valves[parts[0]] = ValveData(name: parts[0], flow: Int(parts[1])!, tunnels: Array(parts[2...]))
    }
}

func part1() {
    valves.forEach { print($0) }
}

if CommandLine.arguments.count > 1 {
    readData(path: CommandLine.arguments[1])
    part1()
} else {
    print("usage: d16 <filename>")
}
