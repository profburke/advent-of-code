#!/usr/bin/env swift

import Darwin

func readNumbers(path: String) -> Int {
    var sum = 0
    
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    while let line = readLine() {
        var power = -1
        sum += line.reversed().map { c in
                                     power += 1
                                     let place = Int(pow(Double(5), Double(power)))
                                     switch c {
                                     case "0":
                                         return 0
                                     case "1":
                                         return place
                                     case "2":
                                         return 2 * place
                                     case "-":
                                         return -1 * place
                                     case "=":
                                         return -2 * place
                                     default:
                                         print("Error in input: \(c)")
                                         exit(1)
                                     }
        }.reduce(0, +)
    }

    return sum
}

func snafu(_ n: Int) -> Int {
    let digits = String(n)
    let hpow = digits.count
    var result = 0

    
    return result
}

func part1(sum: Int) {
    print(sum, snafu(sum))
}

if CommandLine.arguments.count > 1 {
    let sum = readNumbers(path: CommandLine.arguments[1])
    part1(sum: sum)
} else {
    print("usage: d25.swift <filename>")
}
