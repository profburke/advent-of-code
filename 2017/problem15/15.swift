#!/usr/bin/swift
import Foundation

struct Generator {
    static let divisor = 2147483647
    let factor: Int
    var previous: Int

    init(factor: Int, seed: Int) {
        self.factor = factor
        previous = seed
    }
    
    mutating func next() -> Int {
        let product = factor * previous
        previous = product % Generator.divisor
        return previous
    }
}


var generatorA = Generator(factor: 16807, seed: 618)
var generatorB = Generator(factor: 48271, seed: 814)

var matches = 0
for i in 0..<40_000_000 {
    let valA = generatorA.next()
    let valB = generatorB.next()
    let match = (valA & 0xFFFF) == (valB & 0xFFFF)
    if match { matches += 1 }
    //print("\(valA) \(valB) \(match)")
    if i != 0 && i%100_000 == 0 { print(".", terminator: "") }
    if i != 0 && i%1_000_000 == 0 { print("") }
}
print("matches: \(matches)")
