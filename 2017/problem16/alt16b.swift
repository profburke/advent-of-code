#!/usr/bin/swift
import Foundation

var programs: [String] = ["a", "b", "c", "d",
                          "e", "f", "g", "h",
                          "i", "j", "k", "l",
                          "m", "n", "o", "p"]

var endpoint: [String] = Array<String>(repeating: "", count: 16)

// fnloekigdmpajchb

func dance() {
    endpoint[0] = programs[5]
    endpoint[1] = programs[13]
    endpoint[2] = programs[11]
    endpoint[3] = programs[14]
    endpoint[4] = programs[4]
    endpoint[5] = programs[10]
    endpoint[6] = programs[8]
    endpoint[7] = programs[6]
    endpoint[8] = programs[3]
    endpoint[9] = programs[12]
    endpoint[10] = programs[15]
    endpoint[11] = programs[0]
    endpoint[12] = programs[9]
    endpoint[13] = programs[2]
    endpoint[14] = programs[7]
    endpoint[15] = programs[1]
    for i in 0..<16 {
        programs[i] = endpoint[i]
    }
}


print(programs.joined())
for _ in 0..<10 {
    dance()
    print(programs.joined())
}

