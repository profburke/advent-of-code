#!/usr/bin/env swift

import Darwin

class Tree: CustomStringConvertible {
    let height: Int
    var visible = false

    var description: String {
        let start = visible ? "\u{001B}[31m" : ""
        let finish = visible ? "\u{001B}[0m" : ""
        return "\(start)\(height)\(finish)"
    }

    init(height: Int) {
        self.height = height
    }
}

class Grid {
    var trees: [[Tree]] = []

    var width: Int {
        return trees[0].count
    }

    var height: Int {
        return trees.count
    }

    func printGrid() {
        trees.forEach { row in
                        row.forEach { print($0, terminator: "") }
                        print()
        }
    }

}

func slurp(path: String) -> Grid {
    guard let _ = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        exit(1)
    }

    let grove = Grid()
    //var grove: [[Tree]] = []

    while let line = readLine() {
        var row: [Tree] = []

        for d in line {
            row.append(Tree(height: Int(String(d))!))
        }
        
        grove.trees.append(row)
    }

    return grove
}

func part1(_ grove: Grid) {
    let width = grove.width
    let height = grove.height

    var maxHeight = Int.min
    grove.trees.forEach { row in
                          // scan each row from left
                          maxHeight = Int.min
                          row.forEach { tree in
                                        if maxHeight < tree.height { tree.visible = true }
                                        if tree.height > maxHeight { maxHeight = tree.height }
                          }
                          
                          // scan each row from right
                          maxHeight = Int.min
                          row.reversed().forEach { tree in
                                                   if maxHeight < tree.height { tree.visible = true }
                                                   if tree.height > maxHeight { maxHeight = tree.height }
                          }
    }
    
    // scan each column from top
    for x in 0 ..< width { // could optimize by skipping first and last columns, but ...
        maxHeight = Int.min
        for y in 0 ..< height {
            let tree = grove.trees[y][x]
            if maxHeight < tree.height { tree.visible = true }
            if tree.height > maxHeight { maxHeight = tree.height }
        }
    }
    
    // scan each column from bottom
    for x in 0 ..< width { // could optimize by skipping first and last columns, but ...
        maxHeight = Int.min
        for y in (0 ..< height).reversed() {
            let tree = grove.trees[y][x]
            if maxHeight < tree.height { tree.visible = true }
            if tree.height > maxHeight { maxHeight = tree.height }
        }
    }

    grove.printGrid()
    
    let count = grove.trees.map { row in row.map { tree in tree.visible ? 1 : 0}.reduce(0, +) }.reduce(0, +)
    print("Visible count: \(count)")
}

func part2(_ grove: Grid) {
    var maxScore = Int.min
    let width = grove.width
    let height = grove.height
    
    for x in 1 ..< width - 1 { // piddling optimization: don't bother with trees on edge
        for y in 1 ..< height - 1 {
            let h = grove.trees[y][x].height
            // look up
            var upScore = 0
            for yp in (0 ..< y).reversed() {
                let hp = grove.trees[yp][x].height
                upScore += 1
                if hp >= h { break }
            }
            
            // look down
            var downScore = 0
            for yp in ((y + 1) ..< height) {
                let hp = grove.trees[yp][x].height
                downScore += 1
                if hp >= h { break }
            }

            // look left
            var leftScore = 0
            for xp in (0 ..< x).reversed() {
                let hp = grove.trees[y][xp].height
                leftScore += 1
                if hp >= h { break }
            }
            
            // look right
            var rightScore = 0
            for xp in ((x + 1) ..< width) {
                let hp = grove.trees[y][xp].height
                rightScore += 1
                if hp >= h { break }
            }

            let score = upScore * downScore * leftScore * rightScore
            if score > maxScore { maxScore = score }
        }
    }

    print("Max scenic score is \(maxScore)")
}

let grove = slurp(path: CommandLine.arguments[1])
part1(grove)
part2(grove)
