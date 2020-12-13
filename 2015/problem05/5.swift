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

func checkForDouble(_ line: String) -> Bool {
    var previous: Character = "A"
    var found = false
    line.forEach { char in
                   if char == previous {
                       found = true
                   } 
                   previous = char
    }
    return found
}

let vowels = "aeiou"
func process(_ line: String) -> Int {
    let strippedLine: String = line.filter { vowels.index(of: $0) != nil }
    let has3Vowels = strippedLine.count > 2
    let hasDoubleLetter = checkForDouble(line)
    let hasForbiddenSubstring = line.contains("ab") || line.contains("cd") || line.contains("pq") || line.contains("xy")
    return (has3Vowels && hasDoubleLetter && !hasForbiddenSubstring) ? 1 : 0
}

let lines = data.split(separator: "\n")
print(lines.reduce(0) { $0 + process(String($1)) })
