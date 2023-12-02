#!/usr/bin/env swift

import Darwin
import class Foundation.NSString

struct Point: Hashable, CustomStringConvertible {
    let x: Int
    let y: Int

    var description: String {
        return "(\(x), \(y))"
    }

    static func +(lhs: Point, rhs: Point) -> Point {
        return Point(x: lhs.x + rhs.x, y: lhs.y + rhs.y)
    }

    func distance(to p: Point) -> Int {
        return abs(self.x - p.x) + abs(self.y - p.y)
    }
}

struct Sensor: CustomStringConvertible {
    let position: Point
    let closestBeacon: Point

    var distance: Int { return position.distance(to: closestBeacon) }

    var description: String {
        return "<\(position)> - \(distance) @<\(closestBeacon)>"
    }
}

var sensors:[Point:Sensor] = [:]
var beaconSpots = Set<Point>()

func readData(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(path)")
        exit(1)
    }

    while let line = readLine() {
        let parts = line.replacingOccurrences(of: "Sensor at ", with: "")
        .replacingOccurrences(of: " closest beacon is at ", with: "")
        .replacingOccurrences(of: " ", with: "")
        .replacingOccurrences(of: "x=", with: "")
        .replacingOccurrences(of: "y=", with: "")
        .split(separator: ":").map { String($0) }
        
        let sCoords = parts[0].split(separator: ",").map { Int(String($0))! }
        let bCoords = parts[1].split(separator: ",").map { Int(String($0))! }

        let sPos = Point(x: sCoords[0], y: sCoords[1])
        let bPos = Point(x: bCoords[0], y: bCoords[1])
        
        sensors[sPos] = (Sensor(position: sPos, closestBeacon: bPos))
        beaconSpots.insert(bPos)
    }
}

func scanRow(row: Int) -> Int {
    var impossibleSpots = Set<Point>()
    
    sensors.values.forEach { sensor in
                             let p0 = sensor.position

                             // if the sensor is on the line, add it
                             if p0.y == row {
                                 // print("+ inserting \(p0)")
                                 impossibleSpots.insert(p0)
                             }
                             // if a beacon is on the line ... while a place where there's a beacon
                             // can't be a place where "there can't be a beacon"
                             
                             // find distance from sensor to line line y=row
                             let d = abs(p0.y - row)

                             if d <= sensor.distance {
                                 // let p be spot on the line closest to sensor
                                 var p = Point(x: p0.x, y: row)
                                 // add p to impossibleSpots UNLESS there is a beacon there
                                 if !beaconSpots.contains(p) {
                                     // print("++ inserting \(p)")
                                     impossibleSpots.insert(p)
                                 }

                                 let delta = sensor.distance - d
                                 for i in 0 ... delta {
                                     //       find two spots on the line that are dp away
                                     //       add them to impossibleSpots if in range

                                     var x = p0.x + i
                                     if x <= 4_000_000 {
                                         p = Point(x: x, y: row)
                                         if p0.distance(to: p) <= sensor.distance && !beaconSpots.contains(p) {
                                             // print("+++ inserting \(p)")
                                             impossibleSpots.insert(p)
                                         }
                                     }

                                     x = p0.x - i
                                     if x >= 0 {
                                         p = Point(x: x, y: row)
                                         if p0.distance(to: p) <= sensor.distance && !beaconSpots.contains(p) {
                                             // print("++++ inserting \(p)")
                                             impossibleSpots.insert(p)
                                         }
                                     }
                                 }
                             }
    }

    return impossibleSpots.count
}

func part1() {
    print(scanRow(row: 2000000))
}

func part2() {
    for row in 0 ... 4000000 {
        let c = scanRow(row: row)
        if c < 10 {
            print(row, c)
        }
    }
}

if CommandLine.arguments.count > 1 {
    readData(path: CommandLine.arguments[1])
    // part1()
    part2()
} else {
    print("usage: ./d15.swift <filename>")
}
