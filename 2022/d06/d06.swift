#!/usr/bin/env swift

import Darwin

func allDifferent(_ buffer: [Character]) -> Bool {
    var s = Set<Character>()
    buffer.forEach { c in
                     s.insert(c)
    }

    return s.count == buffer.count
}

func process(path: String, markerLength: Int) {
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    var buffer: [Character] = []
    var position = 0
    
    while true {
        let c = fgetc(file)

        if c == -1 {
            break
        }

        position += 1
        let ch = Character(UnicodeScalar(UInt8(c)))

        if buffer.isEmpty {
            for _ in 0..<markerLength {
                buffer.append(ch)
            }
        }
        
        buffer.append(ch)
        buffer.removeFirst()
        
        if allDifferent(buffer) {
            print("Marker found a position \(position)")
            break
        }
    }
}

process(path: CommandLine.arguments[1], markerLength: 4)
process(path: CommandLine.arguments[1], markerLength: 14)
