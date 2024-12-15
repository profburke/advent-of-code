import Foundation

let seeds = readLine(strippingNewline: true)!.dropFirst(7).split(separator: " ").map { Int($0)! }
_ = readLine(strippingNewline: true) // skip blank line

var categories: [Category] = []
var maps: [Map] = []
var c: Category?

while true {
    guard let line = readLine(strippingNewline: true) else { break }

    if line.isEmpty {
        c = Category(maps: maps)
        categories.append(c!)
        maps = []
    } else if line[line.startIndex].isDigit {
        let ints = line.split(separator: " ")
        let m = Map(destination: Int(ints[0])!, source: Int(ints[1])!, length: Int(ints[2])!)
        maps.append(m)
    } // we don't really care about the lines that read x-to-y map, so skip then
}

// add last caetegory
c = Category(maps: maps)
categories.append(c!)

// Now process...

func plant(_ seeds: [Int]) -> Int {
    var minLocation = Int.max

    seeds.forEach { seed in
                    var r = seed
                    
                    //print(r, terminator: "")
                    
                    categories.forEach { c in
                                         r = c.map(r)
                                         //print(" -> \(r)", terminator: "")
                    }
                    //print()
                    
                    if r < minLocation { minLocation = r }
    }

    return minLocation
}

func part1() {
    let n = plant(seeds)
    print("Part 1 - \(n)")
}

func part2() {
    var allSeeds: [Int] = []

    for i in stride(from: 0, to: seeds.count, by: 2) {
        let base = seeds[i]
        let range = seeds[i+1]

        for j in  0 ... range {
            allSeeds.append(base + j)
        }
    }
    
    let n = plant(allSeeds)
    print("Part 2 - \(n)")
}

//part1()
part2()

// Regex keeps crashing program. Why? :shrug:
//let mapParser = #/\d+ \d+ \d+/#

struct Map: CustomStringConvertible {
    let destination: Int
    let source: Int
    let length: Int

    var description: String {
        "\(source):\(length) -> \(destination)"
    }

    func contains(_ n: Int) -> Bool {
        source <= n && n < source + length
    }

    // assumes n is in source range
    func map(_ n: Int) -> Int {
        destination + (n - source)
    }
}

struct Category {
    let maps: [Map]

    func map(_ n: Int) -> Int {
        guard let m = maps.first(where: { $0.contains(n) }) else { return n }

        return m.map(n)
    }
}

extension Character {
    var isDigit: Bool {
        self.isASCII && self.isNumber
    }
}
