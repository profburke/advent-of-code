import Foundation

var numbers: [Entity] = []
var symbols: [Entity] = []

extension Character {
    var isDigit: Bool {
        self.isASCII && self.isNumber
    }
}

var row = 0
while true {
    guard let line = readLine(strippingNewline: true) else { break }
    var col = 0
    var inNumber = false
    var token = ""
    
    line.forEach { c in
                   
                   if c.isDigit {
                       inNumber = true
                       token.append(c)
                   } else if c == "." {
                       if inNumber {
                           let n = Entity(token: token, coords: Coordinates(row: row, col: col - 1))
                           numbers.append(n)
                           token = ""
                           inNumber = false
                       }
                   } else {
                       let e = Entity(token: String(c), coords: Coordinates(row: row, col: col))
                       symbols.append(e)

                       if inNumber {
                           let n = Entity(token: token, coords: Coordinates(row: row, col: col - 1))
                           numbers.append(n)
                           token = ""
                           inNumber = false
                       }
                   }

                   col += 1
    }

    if inNumber {
        let n = Entity(token: token, coords: Coordinates(row: row, col: col - 1))
        numbers.append(n)
    }

    row += 1
}

func part1(numbers: [Entity], symbols: [Entity]) {
    var gearish: [Coordinates: [Int]] = [:]
    var symbollocs: [Coordinates: Entity] = [:]
    symbols.forEach { s in
                      symbollocs[s.coords] = s
    }

    var sum = 0
    for n in numbers {
        var potentialSpots = Set<Coordinates>()
        var isPartNumber = false
        let lastCol = n.coords.col
        let firstCol = lastCol - n.token.count + 1
        let row = n.coords.row
        
        // check all adjacent spots (incl diagonal)
        for col in (firstCol - 1) ... (lastCol + 1) {
            let above = Coordinates(row: row - 1, col: col)
            if symbollocs[above] != nil {
                isPartNumber = true
                potentialSpots.insert(above)
            }
            let below = Coordinates(row: row + 1, col: col)
            if symbollocs[below] != nil {
                isPartNumber = true
                potentialSpots.insert(below)
            }
        }

        let before = Coordinates(row: row, col: firstCol - 1)
        if symbollocs[before] != nil {
            isPartNumber = true
            potentialSpots.insert(before)
        }
        let after = Coordinates(row: row, col: lastCol + 1)
        if symbollocs[after] != nil {
            isPartNumber = true
            potentialSpots.insert(after)
        }

        if isPartNumber {
            let val = Int(n.token) ?? 0
            sum += val
        }

        potentialSpots.forEach { spot in
                                 if let s = symbollocs[spot],
                                 s.token == "*" {
                                     var sprockets = gearish[spot, default: []]
                                     sprockets.append(Int(n.token) ?? 0)
                                     gearish[spot] = sprockets
                                 }
        }
    }

    print("Part 1 - \(sum)")

    sum = 0
    for (_, sprockets) in gearish {
        if sprockets.count == 2 {
            sum += sprockets[0] * sprockets[1]
        }
    }

    print("Part 2 - \(sum)")
}

part1(numbers: numbers, symbols: symbols)


struct Coordinates: Hashable, CustomStringConvertible {
    let row: Int
    let col: Int

    var description: String {
        "\(row)-\(col)"
    }
}

struct Entity: CustomStringConvertible {
    let token: String
    let coords: Coordinates

    var description: String {
        "<\(token): \(coords)>"
    }
}
