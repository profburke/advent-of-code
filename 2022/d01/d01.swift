#!/usr/bin/env swift

import Darwin

func process(path: String) {
    var max = Int.min
    var totals: [Int] = []

    func flush() {
            if elfTotal > max {
                max = elfTotal
            }

            totals.append(elfTotal)
            elfTotal = 0
    }
    
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
            flush()
        } else {
            let val = Int(line)!
            elfTotal += val
        }
    }
    
    // Don't forget the last elf!
    flush()
    
    print("Part 1 is \(max)")
    print("Part 2 is \(totals.sorted().suffix(3).reduce(0, +))")
}

process(path: CommandLine.arguments[1])
