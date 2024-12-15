import Foundation
import RegexBuilder

enum Direction: Character {
    case left = "L"
    case right = "R"
}

class Node: CustomStringConvertible {
    let name: String
    var left: Node?
    var right: Node?

    init(name: String) {
        self.name = name
    }

    var description: String {
        "<\(name): l:\(left?.name ?? "") r:\(right?.name ?? "")>"
    }
}

// Using regex literal would be simpler, but I wanted to get hang of
// regex builder ...
let nodeNameRef = Reference(Substring.self)
let leftNameRef = Reference(Substring.self)
let rightNameRef = Reference(Substring.self)

let nodeMatcher = Regex {
    Capture(as: nodeNameRef) {
        OneOrMore(.word)
    }
    
    " = ("

    Capture(as: leftNameRef) {
        OneOrMore(.word)
    }

    ", "

    Capture(as: rightNameRef) {
        OneOrMore(.word)
    }
}

func part1() {
    var steps = 0
    var current = nodes["AAA"]!
    var dIndex = 0

    // print("At \(current.name)")

    while current.name != "ZZZ" {
        switch directions[dIndex] {
        case .left: current = nodes[current.left?.name ?? ""]!
        case .right: current = nodes[current.right?.name ?? ""]!
        }

        steps += 1
        dIndex = (dIndex + 1) % directions.count
        
        // print("Step: \(steps) - At \(current.name)")
    }

    print("Part 1 - \(steps) steps")
}

func atEnd(nodes: [Node]) -> Bool {
    nodes.allSatisfy { $0.name.suffix(1) == "Z" }
}

func part2() {
    var steps = 0
    var dIndex = 0

    // collect all the start points
    var current: [Node] = []
    for (_, node) in nodes {
        if node.name.suffix(1) == "A" { current.append(node) }
    }
    print("Start nodes: \(current)")

    
    while !atEnd(nodes: current)  {
        var temp: [Node] = []

        current.forEach { node in
                          switch directions[dIndex] {
                          case .left: temp.append(nodes[node.left?.name ?? ""]!)
                          case .right: temp.append(nodes[node.right?.name ?? ""]!)
                          }
        }

        current = temp
        steps += 1
        dIndex = (dIndex + 1) % directions.count
        
        // print("Step: \(steps) - At \(current.name)")
        if steps % 1_000_000 == 0 { print("Steps \(steps)") }
    }

    print("Part2 - \(steps) steps")
}

var directions: [Direction] = []
var nodes: [String: Node] = [:]

// read input

readLine(strippingNewline: true)?.forEach { c in
                                            directions.append(Direction(rawValue: c)!)
}

_ = readLine(strippingNewline: true) // skip blank line

while true {
    guard let line = readLine(strippingNewline: true) else { break }
    guard let result = line.firstMatch(of: nodeMatcher) else {
        print("Couldn't parse '\(line)'")
        continue
    }

    let nn = String(result[nodeNameRef])
    let ln = String(result[leftNameRef])
    let rn = String(result[rightNameRef])

    let l = nodes[ln, default: Node(name: ln)]
    let r = nodes[rn, default: Node(name: rn)]
    nodes[ln] = l; nodes[rn] = r
    
    let n = nodes[nn, default: Node(name: nn)]
    n.left = l; n.right = r
    nodes[nn] = n
}

// and process it ...

part1()
part2()
