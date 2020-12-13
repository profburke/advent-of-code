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

input.forEach { char in
                let num: Int = Int(char.unicodeScalars.first!.value)
                lengths.append(num)
}
lengths.append(17)
lengths.append(31)
lengths.append(73)
lengths.append(47)
lengths.append(23)




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

for _ in 0..<64 {
    for l in lengths {
        myswap(l: l)
        pos += (skip + l)%256
        skip += 1
    }
}

var denseHash: [Int] = []

for i in 0..<16 {
    let base = 16 * i
    var result = numbers[base]
    for j in 1..<16 {
        result = result ^ numbers[base+j]
    }
    denseHash.append(result)
}

for n in denseHash {
    print(String(format: "%02x", n), terminator: "")
}
print("")

