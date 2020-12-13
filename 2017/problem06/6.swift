#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName).dropLast() else {
    print("Couldn't open \(fileName)")
    exit(1)
}

func fingerprint(_ blocks: [Int]) -> String {
    return blocks.map { String($0) }.joined(separator:",")
}

func maxBlock(_ blocks: [Int]) -> Int {
    var max = Int.min
    var result = -1
    for (idx, val) in blocks.enumerated() {
        if val > max {
            max = val
            result = idx
        }
    }
    return result
}

func redistribute(_ blocks: [Int], startingWith index: Int) -> [Int] {
    var results = blocks
    var currentIndex = index

    let numToMove = results[currentIndex]
    results[currentIndex] = 0
    for _ in 0..<numToMove {
        currentIndex = (currentIndex + 1) % results.count
        results[currentIndex] += 1
    }
    return results
}

var seen = Set<String>()
var blocks = data.split(separator: "\t", maxSplits: Int.max,
                        omittingEmptySubsequences: true).map { Int($0) ?? 0 }
print(blocks)
let fp = fingerprint(blocks)
seen.insert(fp)
print("\(fp) <-- Start")

var done = false
var steps = 0

while !done {
    let mbIdx = maxBlock(blocks)
    blocks = redistribute(blocks, startingWith: mbIdx)
    steps += 1
    let fp = fingerprint(blocks)

    print(fp)

    if seen.contains(fp) { done = true }
    seen.insert(fp)
}

print(steps)
