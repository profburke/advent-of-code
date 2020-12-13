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

class Queue {
    private var _queue: [Int] = []

    var isEmpty: Bool {
        return _queue.isEmpty
    }

    func enque(_ item: Int) {
        _queue.append(item)
    }

    func deque() -> Int {
        return _queue.removeFirst()
    }
    
}

var connections: [Int : [Int]] = [:]

func parse(_ line: String) {
    let components = line.split(separator: " ")
    let key = Int(components[0])!
    var destinations: [Int] = []
    for i in 2..<components.count {
        let destination = Int(components[i].filter { $0 != "," })!
        destinations.append(destination)
    }
    connections[key] = destinations
}

var seen = Set<Int>()
var lines = data.split(separator: "\n")
_ = lines.map { parse(String($0)) }

var groups = 0

for i in 0..<connections.keys.count {
    if !seen.contains(i) {
        let queue = Queue()
        queue.enque(i)
        seen.insert(i)
        groups += 1
        
        while !queue.isEmpty {
            let current = queue.deque()
            seen.insert(current)
            let endpoints = connections[current]!
            for endpoint in endpoints {
                if !seen.contains(endpoint) {
                    seen.insert(endpoint)
                    let nextEndpoints = connections[endpoint]!
                    for nextEndpoint in nextEndpoints {
                        queue.enque(nextEndpoint)
                    }
                }
            }
        }
    }
}

print(groups)

