#!/usr/bin/env swift

import Darwin

enum Command {
    case noop
    case add(Int)
}

func draw(_ position: Int, _ x: Int) {
    let c: String
    if position == x - 1 || position == x || position == x + 1 {
        c = "#"
    } else {
        c = "."
    }
    print(c, terminator: "")
}

func part1(path: String) {
    var cycle = 1
    var x = 1
    var position = 0
    
    var sumSignalStrengths = 0
    let keyTimes: Set = [20, 60, 100, 140, 180, 220]
    
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot open \(path)")
        return
    }

    while true /*let line = readLine() */ {
        var command: Command = .noop

        if let line = readLine() {
            let parts = line.split(separator: " ")
            if parts[0] == "addx" {
                command = .add(Int(parts[1])!)
            }
        }
        
        // get signal strength if correct cycle
        if keyTimes.contains(cycle) {
            sumSignalStrengths += cycle * x
        }

        draw(position, x)
        
        cycle += 1
        position += 1
        if position == 40 {
            print()
            position = 0
        }
        
        if case let .add(val) = command {
            // get signal strength if correct cycle
            if keyTimes.contains(cycle) {
                sumSignalStrengths += cycle * x
            }

            draw(position, x)
            
            x += val
            cycle += 1

            position += 1
            if position == 40 {
                print()
                position = 0
            }
        }

        if cycle > 220 { break }
    }

    print()
    print("Sum of signal strengths: \(sumSignalStrengths)")
}

if (CommandLine.arguments.count > 1) {
    part1(path: CommandLine.arguments[1])
}
