#!/usr/bin/swift
import Foundation

var tape: [Int : Int] = [:]
var position = 0

enum State {
    static var current: State  = .A
    
    case A
    case B
    case C
    case D
    case E
    case F

    func next() {
        let curVal = tape[position] ?? 0
        switch State.current {
        case .A:
            if curVal == 0 {
                tape[position] = 1
                position += 1
                State.current = .B
            } else {
                tape[position] = 0
                position += 1
                State.current = .C
            }
        case .B:
            if curVal == 0 {
                tape[position] = 0
                position -= 1
                State.current = .A
            } else {
                tape[position] = 0
                position += 1
                State.current = .D
            }
        case .C:
            if curVal == 0 {
                tape[position] = 1
                position += 1
                State.current = .D
            } else {
                tape[position] = 1
                position += 1
                State.current = .A
            }
        case .D:
            if curVal == 0 {
                tape[position] = 1
                position -= 1
                State.current = .E
            } else {
                tape[position] = 0
                position -= 1
                State.current = .D
            }
        case .E:
            if curVal == 0 {
                tape[position] = 1
                position += 1
                State.current = .F
            } else {
                tape[position] = 1
                position -= 1
                State.current = .B
            }
        case .F:
            if curVal == 0 {
                tape[position] = 1
                position += 1
                State.current = .A
            } else {
                tape[position] = 1
                position += 1
                State.current = .E
            }
        }
    }
}

for _ in 0..<12368930 {
    State.current.next()
}

// calc checksum
var checksum = 0
for key in tape.keys {
    if tape[key]! == 1 { checksum += 1 }
}
print(checksum)
