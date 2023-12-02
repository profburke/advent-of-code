#!/usr/bin/env swift

import Darwin

func part1(path: String) {
    var stacks: [[Character]] = []

    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }
    
    defer {
        fclose(file)
    }


    while let line = readLine() {
        if line.isEmpty { break } // blank line separates stacks from commands

        var stack: [Character] = []
        line.forEach { c in
                       stack.append(c)
        }
        
        stacks.append(stack)
    }

    while let line = readLine() {
        // process command (move # of blocks from n to m
        let parts = line.split(separator: " ")
        let nblocks = Int(parts[0])!
        let source = Int(parts[1])! - 1
        let destination = Int(parts[2])! - 1

        for _ in 0..<nblocks {
            let b = stacks[source].removeLast()
            stacks[destination].append(b)
        }
    }

    for var stack in stacks {
        print(stack.removeLast(), terminator: "")
    }
    print()
}

func part2(path: String) {
    var stacks: [[Character]] = []

    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }
    
    defer {
        fclose(file)
    }


    while let line = readLine() {
        if line.isEmpty { break } // blank line separates stacks from commands

        var stack: [Character] = []
        line.forEach { c in
                       stack.append(c)
        }
        
        stacks.append(stack)
    }

    while let line = readLine() {
        // process command (move # of blocks from n to m
        let parts = line.split(separator: " ")
        let nblocks = Int(parts[0])!
        let source = Int(parts[1])! - 1
        let destination = Int(parts[2])! - 1

        var temp: [Character] = []
        for _ in 0..<nblocks {
            let b = stacks[source].removeLast()
            temp.append(b)
        }
        for _ in 0..<nblocks {
            let b = temp.removeLast()
            stacks[destination].append(b)
        }
    }

    for var stack in stacks {
        print(stack.removeLast(), terminator: "")
    }
    print()
}

// part1(path: CommandLine.arguments[1])
part2(path: CommandLine.arguments[1])
