#!/usr/bin/swift
import Foundation

var pos = 0
var skip = 0
var lengths: [Int] = []
var numbers: [Int] = []
for i in 0..<256 {
    numbers.append(i)
}

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let input = try? String(contentsOfFile: fileName).dropLast() else {
    print("Couldn't open \(fileName)")
    exit(1)
}

lengths = input.split(separator: ",").map { Int($0) ?? -1 }

func myswap(l: Int) {
    var temp: [Int] = []
    var rev: [Int] = []
    for i in 0..<l {
        temp.append(numbers[(pos + i)%256])
    }
    for i in 0..<l {
        rev.append(temp[l - i - 1])
    }
    for i in 0..<l {
        numbers[(pos + i)%256] = rev[i]
    }
}

for l in lengths {
    myswap(l: l)
    pos += (skip + l)%256
    skip += 1
}

print(numbers[0]*numbers[1])
