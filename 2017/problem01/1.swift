#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

// get file contents, removing newline at end
let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName).dropLast() else {
    print("Couldn't open \(fileName)")
    exit(1)
}

// YUCK!  Can we do something nicer with map or reduce?

extension String {
    func char(at position: Int) -> Character {
        let idx = self.index(self.startIndex, offsetBy: position%self.count) 
        return self[idx]
    }
}

let str = String(data)
var sum = 0
for i in 0..<str.count {
    let currentChar = str.char(at: i)
    let nextChar = str.char(at: i + 1)
    if currentChar == nextChar {
        sum += Int("\(currentChar)")!
    }
}
print(sum)
