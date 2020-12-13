#!/usr/bin/swift
import Foundation

func myswap(l: Int, pos: Int, numbers: inout [Int]) {
     var temp: [Int] = []
     var rev: [Int] = []
     for i in 0..<l {
         temp.append(numbers[(pos + i)%256])
     }
     for i in 0..<l {
         rev.append(temp[l - i - 1])
     }
     for i in 0..<l {
         numbers[(pos + i)%256] = rev[i]
     }
}


func knotHash(for key: String) -> String {
    var pos = 0
    var skip = 0
    var lengths: [Int] = []
    key.forEach { char in
                  let num: Int = Int(char.unicodeScalars.first!.value)
                  lengths.append(num)
    }
    lengths.append(17)
    lengths.append(31)
    lengths.append(73)
    lengths.append(47)
    lengths.append(23)
    
    var numbers: [Int] = Array<Int>(repeating: 0, count: 256)
    for i in 0..<256 {
        numbers[i] = i
    }

    for _ in 0..<64 {
        for l in lengths {
            myswap(l: l, pos: pos, numbers: &numbers)
            pos += (skip + l)%256
            skip += 1
        }
    }

    var denseHash: [Int] = []

    for i in 0..<16 {
        let base = 16 * i
        var result = numbers[base]
        for j in 1..<16 {
            result = result ^ numbers[base+j]
        }
        denseHash.append(result)
    }

    var result = ""
    for n in denseHash {
        result += String(format: "%02x", n)
    }
    return result
}

let lookup: [Character : [Int]] = [
    "0" : [0, 0, 0, 0],
    "1" : [0, 0, 0, 1],
    "2" : [0, 0, 1, 0],
    "3" : [0, 0, 1, 1],
    "4" : [0, 1, 0, 0],
    "5" : [0, 1, 0, 1],
    "6" : [0, 1, 1, 0],
    "7" : [0, 1, 1, 1],
    "8" : [1, 0, 0, 0],
    "9" : [1, 0, 0, 1],
    "a" : [1, 0, 1, 0],
    "b" : [1, 0, 1, 1],
    "c" : [1, 1, 0, 0],
    "d" : [1, 1, 0, 1],
    "e" : [1, 1, 1, 0],
    "f" : [1, 1, 1, 1],
    ]

let input = "hwlqcszp"
//let input = "flqrgnkx"

var blockmap = [[Int]]()
for _ in 0..<128 {
    var row = [Int]()
    for _ in 0..<128 {
        row.append(0)
    }
    blockmap.append(row)
}

for row in 0..<128 {
    let key = input + "-\(row)"
    let hash = knotHash(for: key)
    var col = 0
    hash.forEach { char in
                   let blocks = lookup[char]!
                   for offset in 0..<4 {
                       blockmap[row][col + offset] = blocks[offset]
                   }
                   col += 4
    }
}

// create a smaller blockmap for testing

var nbm = [[Int]]()
let size = 12
for _ in 0..<size {
    var row = [Int]()
    for _ in 0..<size {
        row.append(0)
    }
    nbm.append(row)
}

for row in 0..<size {
    let key = input + "-\(row)"
    let hash = knotHash(for: key)
    var col = 0
    for (i, c) in hash.enumerated() {
        if i == (size/4) { break }
        let blocks = lookup[c]!
        for offset in 0..<4 {
            nbm[row][col + offset] = blocks[offset]
        }
        col += 4
    }
}

// blockmap = nbm
/////////////////////////////////////

class Queue {
    private var _queue: [(Int, Int)] = []

    var isEmpty: Bool {
        return _queue.isEmpty
    }

    func enqueue(_ item: (Int, Int)) {
        _queue.append(item)
    }

    func dequeue() -> (Int, Int) {
        return _queue.removeFirst()
    }
}

func paint(row: Int, col: Int) {
    let q = Queue()
    q.enqueue((row, col))
    while !q.isEmpty {
        let (i, j) = q.dequeue()
        // blockmap[i][j] = regions + 1
        blockmap[i][j] = 2
        // add all neighbors to queue if == 1
        if ((i - 1) > -1) && blockmap[i - 1][j] == 1 { q.enqueue((i - 1, j)) }
        if ((i + 1) < 128) && blockmap[i + 1][j] == 1 { q.enqueue((i + 1, j)) }
        if ((j - 1) > -1) && blockmap[i][j - 1] == 1 { q.enqueue((i, j - 1)) }
        if ((j + 1) < 128) && blockmap[i][j + 1] == 1 { q.enqueue((i, j + 1)) }
    }
}

var regions = 0

// for row in blockmap {
//     let out: String = row.map { String(format: "%1x", $0) }.joined(separator: "")
//     print(out)
// }


for row in 0..<blockmap.count {
    for col in 0..<blockmap[0].count {
        if blockmap[row][col] == 1 {
            regions += 1
            paint(row: row, col: col)
        }
    }
}

print(regions)
// for i in 0..<8 {
//     for j in 0..<8 {
//         let c = blockmap[i][j] == 0 ? "." : "#"
//         print(c, terminator: "")
//     }
//     print("")
// }



