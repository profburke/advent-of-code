#!/usr/bin/swift

import Darwin

var currentJet = 0
var currentRock = 0
var currentHeight = 0

func readData(path: String) -> [String] {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    // there's only one line...
    guard let line = readLine() else {
        print("Still cannot read \(path)")
        exit(1)
    }
    
    return line.map{ String($0) }
}

struct P: Hashable, CustomStringConvertible {
    var x: Int
    var y: Int

    init(_ x: Int, _ y: Int) { self.x = x; self.y = y }
    
    var description: String {
        return "(\(x), \(y))"
    }
}

struct Rock {
    var rt: Int
    var left: Int
    var right: Int
    var bottom: Int
    var top: Int
    var points: [P]
}

func canDrop() -> Bool {
    return false
}

func dropRock(_ n: Int, _ jets: [String]) {
    let b = currentHeight + 3

    var rock: Rock
    switch currentRock {
    case 0:
        rock = Rock(rt: 0, left: 2, right: 5, bottom: b, top: b, 
                    points: [P(2, b), P(3, b), P(4, b), P(5, b),])
    case 1:
        rock = Rock(rt: 1, left: 2, right: 4, bottom: b, top: b + 2,
                    points: [P(2, b + 1), P(3, b), P(3, b + 1), P(3, b + 2), P(4, b + 1)])
    case 2:
        rock = Rock(rt: 2, left: 2, right: 4, bottom: b, top: b + 2, 
                    points: [P(2, b), P(3, b), P(4, b), P(4, b + 1), P(4, b + 2)])
    case 3:
        rock = Rock(rt: 3, left: 2, right: 2, bottom: b, top: b + b + 3,
                    points: [P(2, b), P(2, b + 1), P(2, b + 2), P(2, b + 3)])
    case 4:
        rock = Rock(rt: 4, left: 2, right: 3, bottom: b, top: b + 1,
                    points: [P(2, b), P(3, b), P(2, b + 1), P(3, b + 1)])
    default:
        print("Can't get here")
        exit(1)
    }
    currentRock = (currentRock + 1) % 5
    
    while true {
        let jet = jets[currentJet]
        if jet == "<" {
            if rock.left > 0 {
                rock.left -= 1
                rock.right -= 1
                for var p in rock.points {
                    p.x -= 1
                }
            }
        } else { // jet is ">"
            if rock.right < 6 {
                rock.right += 1
                rock.left += 1
                for var p in rock.points {
                    p.x += 1
                }
            }
        }
        currentJet = (currentJet + 1) % jets.count

        // drop
        if canDrop() {
            rock.bottom -= 1
            rock.top -= 1
            for var p in rock.points {
                p.y -= 1
            }
        } else {
            break
        }
    }

    currentHeight = max(currentHeight, rock.top)
}

func part1(_ jets: [String]) {
    for n in 0 ..< 2022 {
        print("drop \(n)")
        dropRock(n, jets)
    }

    print(currentHeight)
}

if CommandLine.arguments.count > 1 {
    let jets = readData(path: CommandLine.arguments[1])
    part1(jets)
} else {
    print("usage: d17 <filename>")
}
