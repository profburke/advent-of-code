#!/usr/bin/swift
import Foundation

class Node {
    let value: Int
    var next: Node?

    init(value: Int, next: Node? = nil) {
        self.value = value
        self.next = next
    }
}

var buffer = Node(value: 0)
buffer.next = buffer

let stepsize = 359

var current = buffer
var value = 1

func cycle() {
    for _ in 0..<stepsize {
        current = current.next!
    }
    let node = Node(value: value, next: current.next)
    current.next = node
    current = current.next!
    value += 1
}

var startdate = Date()
print("starting")
while value < 50_000_001 {
    cycle()
    if value != 0 && value%500_000 == 0 {
        let now = Date()
        print("step \(value) in \(now.timeIntervalSince(startdate)) seconds")
        startdate = now
    }
}

print("done inserting")
current = buffer
while current.value != 0 {
    current = current.next!
    if current === buffer {
        print("traversed entire buffer; didn't find 0")
        break
    }
}
current = current.next!
print(current.value)

// Ok, this took about 1/2 hour....surely there's a better way
