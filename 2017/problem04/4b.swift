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

func anagrams(of word: String) -> [String] {
    var result: [String] = []

    if word.count < 2 {
        return [word]
    }

    for (n, c) in word.enumerated() {
        let idx = word.index(word.startIndex, offsetBy: n)
        var subWord = word
        subWord.remove(at: idx)
        let subAnagrams = anagrams(of: subWord)
        for partial in subAnagrams {
            result.append("\(c)" + partial)
        }
    }
   
    return result
}

func process(_ line: String) -> Int {
    var seen = Set<String>()
    for wordseq in line.split(separator: " ", maxSplits: Int.max, omittingEmptySubsequences: true) {
        let s = String(wordseq.sorted())
        if seen.contains(s) { return 0 }
        seen.insert(s)
    }
    return 1
}

let start = Date()
let lines = data.split(separator: "\n")
print(lines.reduce(0) { $0 + process(String($1)) })
let finish = Date()
print("Time: \( finish.timeIntervalSinceReferenceDate - start.timeIntervalSinceReferenceDate)")
