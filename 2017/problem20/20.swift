#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
     print("Usage: \(CommandLine.arguments[0]) <inputfile>")
     exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName) else {
     print("Couldn't open \(fileName)")
     exit(1)
}

struct Coordinates {
    let x: Int
    let y: Int
    let z: Int
}

func +(lhs: Coordinates, rhs: Coordinates) -> Coordinates {
    return Coordinates(x: lhs.x + rhs.x , y: lhs.y + rhs.y, z: lhs.z + rhs.z)
}


struct PointData {
    var p: Coordinates
    var v: Coordinates
    var a: Coordinates
}

func parse(_ line: String) -> PointData {
    let scanner = Scanner(string: line)
    scanner.charactersToBeSkipped = CharacterSet(charactersIn: "apv=<,> ")
    var x = 0
    var y = 0
    var z = 0
    scanner.scanInt(&x)
    scanner.scanInt(&y)
    scanner.scanInt(&z)
    let p = Coordinates(x: x , y: y, z: z)
    scanner.scanInt(&x)
    scanner.scanInt(&y)
    scanner.scanInt(&z)
    let v = Coordinates(x: x , y: y, z: z)
    scanner.scanInt(&x)
    scanner.scanInt(&y)
    scanner.scanInt(&z)
    let a = Coordinates(x: x , y: y, z: z)
    return PointData(p: p, v: v, a: a)
}

func manhattan(_ p: PointData) -> Int {
    return abs(p.p.x) + abs(p.p.y) + abs(p.p.z)
}

var points: [PointData] = []
var lines = data.split(separator: "\n")
_ = lines.map { points.append(parse(String($0))) }

for _ in 0..<10000 {
    for i in 0..<points.count {
        var point = points[i]
        let p1 = point.p + point.v
        let v1 = point.v + point.a
        point.p = p1
        point.v = v1
        points[i] = point
    }
}

var min = Int.max
var mindex = -1

for i in 0..<points.count {
    let m = manhattan(points[i])
    if m < min {
        min = m
        mindex = i
    }
}
print("points[\(mindex)] = \(min)")
