#!/usr/bin/env swift

import Foundation

func chomp(_ s: String) -> Set<Character> {
    var r = Set<Character>()

    s.forEach { c in
                if !r.contains(c) {
                    r.insert(c)
                }
    }
    
    return r
}

let asciiA = Int(Character("A").asciiValue!)
let asciia = Int(Character("a").asciiValue!)

func priority(c: Character) -> Int {
    guard let v = c.asciiValue else {
        print("whoops - priority: \(c)")
        return -100
    }

    let vv = Int(v)
    
    if c.isUppercase {
        return vv - asciiA + 27
    } else {
        return vv - asciia + 1
    }
}

func part1(path: String) {
    var score = 0
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        let halfIndex = line.count / 2
        let compartment1 = line.prefix(halfIndex)
        let compartment2 = line.suffix(halfIndex)

        let set1 = chomp(String(compartment1))
        let set2 = chomp(String(compartment2))

        let r = set1.intersection(set2)
        guard let c = r.first else {
            print("whoops")
            return
        }

        let p = priority(c: c)
        score += p
    }

    print("Score is \(score)")
}


func sharedItem(_ sets: [Set<Character>]) -> Character {
    let r1 = sets[0].intersection(sets[1])
    let r2 = r1.intersection(sets[2])
    guard let c = r2.first else {
         print("whoops")
         return Character(" ")
    }
    return c
}

func part2(path: String) {
    var score = 0
    var index = 0
    var sets = [
        Set<Character>(),
        Set<Character>(),
        Set<Character>(),
        ]
    
    guard let file = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    defer {
        fclose(file)
    }

    while let line = readLine() {
        sets[index] = chomp(String(line))
        index += 1

        if index == 3 {
            index = 0

            let c = sharedItem(sets)
            score += priority(c: c)
        }
    }

    print("Score is \(score)")
}

part1(path: CommandLine.arguments[1])
part2(path: CommandLine.arguments[1])
