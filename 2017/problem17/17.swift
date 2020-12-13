#!/usr/bin/swift
import Foundation

let stepsize = 359

var buffer = [0]
var current = 0
var value = 1

func cycle() {
    for _ in 0..<stepsize {
        current = (current + 1)%buffer.count
    }
    current += 1
    buffer.insert(value, at: current)
    value += 1
}

while value < 2018 {
    cycle()
    if value != 0 && value%500000 == 0 { print(value) }
}

let index = buffer.index(of: 2017)!
print(buffer[index + 1])
