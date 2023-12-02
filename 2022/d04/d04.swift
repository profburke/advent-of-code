#!/usr/bin/env swift

import Darwin

func part1(path: String) {
    var overlappingPairs = 0
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        let parts = line.split(separator: ",")

        var subParts = parts[0].split(separator: "-")
        let elf1 = Int(subParts[0])! ... Int(subParts[1])!

        subParts = parts[1].split(separator: "-")
        let elf2 = Int(subParts[0])! ... Int(subParts[1])!

        if elf1.mContains(elf2) || elf2.mContains(elf1) {
            overlappingPairs += 1
        }
}

    print(overlappingPairs)
}

func part2(path: String) {
    var overlappingPairs = 0
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        let parts = line.split(separator: ",")

        var subParts = parts[0].split(separator: "-")
        let elf1 = Int(subParts[0])! ... Int(subParts[1])!

        subParts = parts[1].split(separator: "-")
        let elf2 = Int(subParts[0])! ... Int(subParts[1])!

        if elf1.overlaps(elf2) || elf2.overlaps(elf1) {
            overlappingPairs += 1
        }
}

    print(overlappingPairs)
}

// part1(path: CommandLine.arguments[1])
part2(path: CommandLine.arguments[1])

extension ClosedRange {
    func mContains(_ other: ClosedRange) -> Bool {
        if other.lowerBound >= self.lowerBound && other.upperBound <= self.upperBound {
            return true
        } else {
            return false
        }
    }
}
