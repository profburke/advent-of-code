#!/usr/bin/env swift

import Darwin
import class Foundation.NSString

enum Value: CustomStringConvertible {
    case integer(Int)
    case list([Value])

    var description: String {
        switch self {
        case .integer(let val):
            return "\(val)"
        case .list(let l):
            return "[\(l)]"
        }
    }
}

typealias Packet = [Value]
var packetPairs: [[Packet]] = []

func parseNumber(_ tokens: [String], _ ct: Int) -> (Int, Int) {
    return (ct + 1, 12)
}

func parseList(_ tokens: [String], _ ct: Int) -> (Int, [Value]) {
    var ct = ct
    while tokens[ct] != "]" { ct += 1 }
    return (ct + 1, [])
}

func parsePacket(_ tokens: [String]) -> Packet {
    if tokens.count == 0 { return Packet() }
    let nTokens = tokens.count
    var ct = 0
    var val: Int
    var l: [Value]
    
    while ct < nTokens {
        switch tokens[ct] {
        case "[":
            (ct, l) = parseList(tokens, ct + 1)
        case "]":
            fatalError("Should not see a ']' in parsePacket")
        default:
            (ct, val) = parseNumber(tokens, ct)
        }
    }
    
    
    return [.integer(12)]
}

func readPacketPairs(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        return
    }

    var pair: [Packet] = []
    while let line = readLine() {
        if line.isEmpty {
            packetPairs.append(pair)
            pair = []
            continue
        }

        let packet = parsePacket(line.replacingOccurrences(of: "[", with: "[,")
                                 .replacingOccurrences(of: "]", with: ",]")
                                 .split(separator: ",")
                                 .dropFirst().dropLast() // packets always start with '[' and end with ']'
                                 .map { String($0) })
        print(packet)
        pair.append(packet)
    }
}

func part1() {
}

if CommandLine.arguments.count > 1 {
    readPacketPairs(path: CommandLine.arguments[1])
    part1()
} else {
    print("usage: ./d13.swift <inputfile>")
}
