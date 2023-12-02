#!/usr/bin/env swift

import Darwin

struct Node: Equatable {
    let val: Int
    let order: Int
}

var list: [Node] = []

func readFile(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    var order = 0
    while let line = readLine() {
        let n = Node(val: Int(line)!, order: order)
        list.append(n)
        order += 1
    }
}

func part1() {
    let llen = list.count
    
    for i in 0 ..< list.count {
        let i = list.firstIndex { $0.order == i }!
        let item = list[i]
        if item.val != 0 { // 0 doesn't move
            var ni = (i + item.val + llen) % llen
            let adj = (ni > i) ? -1 : 0
            ni += adj
            list.remove(at: i)
            list.insert(item, at: ni)
        }
    }

    list.forEach { print("\($0.val) ", terminator: "") }
    print()
    
    let zi = list.firstIndex { $0.val == 0 }!
    let i1 = (zi + 1000) % llen
    let i2 = (zi + 2000) % llen
    let i3 = (zi + 3000) % llen

    let sum = list[i1].val + list[i2].val + list[i3].val
    print("Sum is \(sum): \(list[i1].val) \(list[i2].val) \(list[i3].val) ")
}

if CommandLine.arguments.count > 1 {
    readFile(path: CommandLine.arguments[1])
    part1()
} else {
    print("usage: d20 <filename>")
}
