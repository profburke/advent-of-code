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
    var seen = Set<String>()
    for wordseq in line.split(separator: " ", maxSplits: Int.max, omittingEmptySubsequences: true) {
        let word = String(wordseq)
        if seen.contains(word) { return 0 }
        seen.insert(word)
    }
    return 1
}

let lines = data.split(separator: "\n")
print(lines.reduce(0) { $0 + process(String($1)) })
