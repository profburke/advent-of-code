// Advent of Code - Day 1, Parts 1 and 2
// Matthew M. Burke
import Foundation

var lines: [String] = []
while true {
    guard let line = readLine(strippingNewline: true) else { break }
    lines.append(line)
}

func part1(_ lines: [String]) {
    var sum = 0
    
    lines.forEach { line in
                    let l = line.filter { c in c.isASCII && c.isNumber }
                    if !l.isEmpty { 
                        let value = Int("\(l.first!)\(l.last!)") ?? 0
                        sum += value
                    }
    }

    print("Part 1 - \(sum)")
}

func part2(_ lines: [String]) {
    var sum = 0

    lines.forEach { line in
                    var line = line
                    var l = ""

                    while !line.isEmpty {
                        let fc = line[line.startIndex]
                        if fc.isNumber {
                            l.append(fc)
                            line = String(line.dropFirst(1))
                        // READ ME -- if there is a possible overlap between this spelled out #
                        // and the next, then don't drop the overlapping letters :shrug:
                        } else if line.hasPrefix("one") {
                            l.append("1")
                            line = String(line.dropFirst(2))
                        } else if line.hasPrefix("two") {
                            l.append("2")
                            line = String(line.dropFirst(2))
                        } else if line.hasPrefix("three") {
                            l.append("3")
                            line = String(line.dropFirst(4))
                        } else if line.hasPrefix("four") {
                            l.append("4")
                            line = String(line.dropFirst(4))
                        } else if line.hasPrefix("five") {
                            l.append("5")
                            line = String(line.dropFirst(3))
                        } else if line.hasPrefix("six") {
                            l.append("6")
                            line = String(line.dropFirst(3))
                        } else if line.hasPrefix("seven") {
                            l.append("7")
                            line = String(line.dropFirst(5))
                        } else if line.hasPrefix("eight") {
                            l.append("8")
                            line = String(line.dropFirst(4))
                        } else if line.hasPrefix("nine") {
                            l.append("9")
                            line = String(line.dropFirst(3))
                        } else {
                            line = String(line.dropFirst(1))
                        }
                    }

                    let value = Int("\(l.first!)\(l.last!)") ?? 0
                    sum += value
    }

    print("Part 2 - \(sum)")
}

part1(lines)
part2(lines)
