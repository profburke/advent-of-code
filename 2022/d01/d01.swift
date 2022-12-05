#!/usr/bin/env swift

import Darwin

func part1(path: String) {
    var max = Int.min

    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    var elfTotal = 0
    while let line = readLine() {
        if line.isEmpty {
            if elfTotal > max {
                max = elfTotal
            }

            elfTotal = 0
        } else {
            let val = Int(line)!
            elfTotal += val
        }
    }
    if elfTotal > max {
        max = elfTotal
    }
    
    print("Max is \(max)")
}


func part2(path: String) {
    var totals: [Int] = []

    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    var elfTotal = 0
    while let line = readLine() {
        if line.isEmpty {
            totals.append(elfTotal)
            elfTotal = 0
        } else {
            let val = Int(line)!
            elfTotal += val
        }
    }
    totals.append(elfTotal) // don't forget the last elf
    
    totals.sort()
    let index = totals.count - 3
    let top3 = totals[index...]
    print("Sum of top 3 is \(top3.reduce(0, +))")
}


part1(path: CommandLine.arguments[1])
// part2(path: CommandLine.arguments[1])
