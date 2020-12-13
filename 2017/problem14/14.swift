#!/usr/bin/swift
import Foundation

func myswap(l: Int, pos: Int, numbers: inout [Int]) {
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


func knotHash(for key: String) -> String {
    var pos = 0
    var skip = 0
    var lengths: [Int] = []
    key.forEach { char in
                  let num: Int = Int(char.unicodeScalars.first!.value)
                  lengths.append(num)
    }
    lengths.append(17)
    lengths.append(31)
    lengths.append(73)
    lengths.append(47)
    lengths.append(23)
    
    var numbers: [Int] = Array<Int>(repeating: 0, count: 256)
    for i in 0..<256 {
        numbers[i] = i
    }

    for _ in 0..<64 {
        for l in lengths {
            myswap(l: l, pos: pos, numbers: &numbers)
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

    var result = ""
    for n in denseHash {
        result += String(format: "%02x", n)
    }
    return result
}

let lookup: [Character : Int] = [
    "0" : 0,
    "1" : 1,
    "2" : 1,
    "3" : 2,
    "4" : 1,
    "5" : 2,
    "6" : 2,
    "7" : 3,
    "8" : 1,
    "9" : 2,
    "a" : 2,
    "b" : 3,
    "c" : 2,
    "d" : 3,
    "e" : 3,
    "f" : 4,
    ]

func oneCount(for hash: String) -> Int {
    var result = 0
    hash.forEach { char in
                   result += lookup[char] ?? 0
    }
    return result
}

let input = "hwlqcszp"
var blocksUsed = 0

for i in 0..<128 {
    let key = input + "-\(i)"
    let hash = knotHash(for: key)
    blocksUsed += oneCount(for: hash)
}

print(blocksUsed)
