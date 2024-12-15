import Foundation

var line = readLine(strippingNewline: true)!
let times = String(line.dropFirst(5)).trimmingCharacters(in: .whitespaces).split(separator: " ").map { Int($0)! }
line = readLine(strippingNewline: true)!
let distances = String(line.dropFirst(9)).trimmingCharacters(in: .whitespaces).split(separator: " ").map { Int($0)! }


func doit(times: [Int], distances: [Int]) -> Int {
    var margin = 1
    for (tmax, dmin) in zip(times, distances) {
        // Can't win if hold button for 0 or tmax seconds ...
        var wins = 0
        for t in 1..<tmax {
            let d = t*(tmax - t)
            if d > dmin { wins += 1 }
        }
        
        margin *= wins
    }

    return margin
}

func part1() {
    let margin = doit(times: times, distances: distances)
    print("Part 1 - \(margin)")
}


func part2() {
    let times = [Int(times.map { "\($0)" }.joined())!]
    let distances = [Int(distances.map { "\($0)" }.joined())!]

    print(times, distances)
    
    let margin = doit(times: times, distances: distances)
    print("Part 2 - \(margin)")
}

part1()
part2()
