#!/usr/bin/env swift

import Darwin
import struct Foundation.CharacterSet
import class Foundation.NSString

enum Expr {
    case constant(Float)
    case variable(coef: Float, constant: Float)

    static func +(lhs: Expr, rhs: Expr) -> Expr {
        switch (lhs, rhs) {
        case (.constant(let l), .constant(let r)):
            return .constant(l + r)
        case (.constant(let l), .variable(let k, let c)):
            return .variable(coef: k, constant: l + c)
        case (.variable(let k, let c), .constant(let r)):
            return .variable(coef: k, constant: c + r)
        default:
            print("This can't be happening +")
            exit(1)
        }
    }

    static func -(lhs: Expr, rhs: Expr) -> Expr {
        switch (lhs, rhs) {
        case (.constant(let l), .constant(let r)):
            return .constant(l - r)
        case (.constant(let l), .variable(let k, let c)):
            return .variable(coef: k, constant: l - c)
        case (.variable(let k, let c), .constant(let r)):
            return .variable(coef: k, constant: c - r)
        default:
            print("This can't be happening -")
            exit(1)
        }
    }

    static func *(lhs: Expr, rhs: Expr) -> Expr {
        switch (lhs, rhs) {
        case (.constant(let l), .constant(let r)):
            return .constant(l * r)
        case (.constant(let l), .variable(let k, let c)):
            return .variable(coef: l * k, constant: l * c)
        case (.variable(let k, let c), .constant(let r)):
            return .variable(coef: k * r, constant: c * r)
        default:
            print("This can't be happening *")
            exit(1)
        }
    }

    static func /(lhs: Expr, rhs: Expr) -> Expr {
        switch (lhs, rhs) {
        case (.constant(let l), .constant(let r)):
            return .constant(l / r)
        case (.constant(let l), .variable(let k, let c)): // is this correct?
            print("BOOOM")
            return .variable(coef: l / k, constant: l / c)
        case (.variable(let k, let c), .constant(let r)):
            return .variable(coef: k / r, constant: c / r)
        default:
            print("This can't be happening /")
            exit(1)
        }
    }
}

enum Operation: String, CustomStringConvertible {
    case addition = "+"
    case subtraction = "-"
    case multiplication = "*"
    case division = "/"
    case equals = "="
    
    func jdi(l: Expr, r: Expr) -> Expr {
        switch self {
        case .addition: return l + r
        case .subtraction: return l - r
        case .multiplication: return l * r
        case .division: return l / r
        case .equals:
            print("\(l) = \(r)")
            switch (l, r) {
            case (.variable(let k, let c), .constant(let c2)):
                print( (c2-c)/k )
            case (.constant(let c), .variable(let k, let c2)):
                print( (c-c2)/k )
            default:
                print("Monkey trouble")
                exit(1)
            }
                
                return .constant(Float(Int.min))
        }
    }
    
    var description: String {
        switch self {
        case .addition: return "+"
        case .subtraction: return "-"
        case .multiplication: return "*"
        case .division: return "/"
        case .equals: return "="
        }
    }
}

enum Payload {
    case constant(Float)
    case operation(Operation)
}

class Node: CustomStringConvertible {
    let name: String
    var left: String?
    var right: String?
    let payload: Payload

    init(name: String, payload: Payload) {
        self.name = name
        self.left = nil
        self.right = nil
        self.payload = payload
    }

    func eval() -> Expr {
        switch payload {
        case .constant(let v):
            return .constant(v)
        case .operation(let op):
            let ln = nodes[left!]!
            let rn = nodes[right!]!
            return op.jdi(l: ln.eval(), r: rn.eval())
        }
    }

    func eval2() -> Expr {
        if name == "root" {
            return Operation.equals.jdi(l: nodes[left!]!.eval2(), r: nodes[right!]!.eval2())
        }
        if name == "humn" {
            return .variable(coef: 1, constant: 0)
        }
        switch payload {
        case .constant(let v):
            return .constant(v)
        case .operation(let op):
            let ln = nodes[left!]!
            let rn = nodes[right!]!
            let lv = ln.eval2()
            let rv = rn.eval2()
            print(lv, op, rv)
            return op.jdi(l: lv, r: rv)
        }
    }
    
    var description: String {
        return "(\(name): \(payload) - l: \(left ?? "") r: \(right ?? ""))"
    }
}

var nodes: [String:Node] = [:]

func readData(path: String) {
    let operatorCS = CharacterSet(charactersIn: "+-/*")
    
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    while let line = readLine() {
        let parts = line.replacingOccurrences(of: " ", with: "")
        .split(separator: ":")

        let name = String(parts[0])
        let components = parts[1].components(separatedBy: operatorCS)

        let node: Node
        if components.count == 1 { // leaf
            node = Node(name: name, payload: .constant(Float(components[0])!))
        } else {
            let r = parts[1].rangeOfCharacter(from: operatorCS)!
            let op = Operation(rawValue: String(parts[1][r]))!
            node = Node(name: name, payload: .operation(op))
            node.left = components[0]
            node.right = components[1]
        }

        nodes[name] = node
    }

    nodes.forEach { print($0) }
}

func part1() {
    let n = nodes["root"]!
    print(n.eval())
}

func part2() {
    let n = nodes["root"]!
    print(n.eval2())
}

if CommandLine.arguments.count > 1 {
    readData(path: CommandLine.arguments[1])
    part1()
    part2()
} else {
    print("usage: d21.swift <filename>")
}
