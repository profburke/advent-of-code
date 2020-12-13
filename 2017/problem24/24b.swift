#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
     print("Usage: \(CommandLine.arguments[0]) <inputfile>")
     exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
     print("Couldn't open \(fileName)")
     exit(1)
}

struct Component: CustomStringConvertible { 
    static var nextId = 0
    let id: Int
    let portA: Int
    let portB: Int

    init(_ a: Int, _ b: Int) {
        id = Component.nextId
        portA = a
        portB = b
        Component.nextId += 1
    }

    var description: String {
        return "(\(id))-\(portA):\(portB)"
    }
}

enum BridgeError: Error {
    case alreadyContains
    case noMatchingEnd
}

struct Bridge: CustomStringConvertible {
    private(set) var ids: Set<Int>
    private(set) var components: [Component]
    private(set) var freeEnd: Int

    init() {
        ids = Set<Int>()
        components = []
        freeEnd = 0
    }

    init(with bridge: Bridge) {
        ids = Set<Int>(bridge.ids)
        components = []
        for comp in bridge.components {
            components.append(comp)
        }
        freeEnd = bridge.freeEnd
    }
    
    mutating func add(component: Component) throws {
        if ids.contains(component.id) {
            throw BridgeError.alreadyContains
        }

        if component.portA != freeEnd && component.portB != freeEnd {
            throw BridgeError.noMatchingEnd
        }
        
        ids.insert(component.id)
        components.append(component)
        if freeEnd == component.portA {
            freeEnd = component.portB
        } else {
            freeEnd = component.portA
        }
    }

    func accepts(component: Component) -> Bool {
        return (freeEnd == component.portA || freeEnd == component.portB) && !ids.contains(component.id)
    }
    
    var score: Int {
        return components.reduce(0) { $0 + $1.portA + $1.portB }
    }

    var description: String {
        return "[\(components.map { $0.description }.joined(separator: ", "))]"
    }
}


var components: [Component] = []

func parse(_ line: String) {
    let slashdex = line.index(of: "/")!
    let a = Int(line.prefix(through: line.index(slashdex, offsetBy: -1)))!
    let b = Int(line.suffix(from: line.index(slashdex, offsetBy: 1)))!
    components.append(Component(a, b))
}
           
var lines = data.split(separator: "\n")
_ = lines.map { parse(String($0)) }
    
var solutions: [Int : [Bridge]] = [:]
var length = 1

// add the length 1 solutions
var length1Solutions: [Bridge] = []
for component in components {
    if component.portA == 0 || component.portB == 0 {
        var bridge = Bridge()
        try? bridge.add(component: component)
        length1Solutions.append(bridge)
    }
}
solutions[1] = length1Solutions

for length in 2..<components.count {
    var newSolutions: [Bridge] = []
    let previousSolutions = solutions[length - 1]!
    
    for component in components {
        for prevSol in previousSolutions {
            if prevSol.accepts(component: component) {
                var newBridge = Bridge(with: prevSol)
                try? newBridge.add(component: component)
                newSolutions.append(newBridge)
            }
        }
    }
    solutions[length] = newSolutions
}

// find max scoring bridge from longest bridges
var max = Int.min
for length in solutions.keys.sorted().reversed() {
    let sols = solutions[length]!
    if sols.count > 0 {
        for sol in sols {
            let score = sol.score
            if score > max { max = score }
        }
        break
    }
}
print(max)
