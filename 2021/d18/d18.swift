import Foundation

struct SnailNumber {
    let l: SnailPart
    let r: SnailPart

    var magnitude: Int {
        return 3 * l.magnitude + 2 * r.magnitude
    }

    func reduce() -> SnailNumber {
        var reduced = self
        
//        while true {
//            if explode(redu) 
//        }
        
        return reduced
    }
}

indirect enum SnailPart {
case n(Int)
case sn(SnailNumber)

    var magnitude: Int {
        switch self {
        case .n(let n):
            return n
        case .sn(let s):
            return s.magnitude
        }
    }
}

func +(left: SnailNumber, right: SnailNumber) -> SnailNumber {
    return SnailNumber(l: .sn(left), r: .sn(right))
}

var index = 0
func parseNumber(_ line: [String]) -> Int {
    var s = ""

    var c = Character(line[index])
    while CharacterSet.decimalDigits.contains(c) {
        s += String(c)
        index += 1
        c = Character(line[index])
    }

    
}

func parseSnailNumber(_ line: [String]) -> SnailNumber {
    let l: SnailPart
    let r: SnailPart
    
    // expect [
    if line[index] != "[" { print("[ oops") }
    index += 1

    if line[index] == "[" {
        let sn = parseSnailNumber(line)
        l = .sn(sn)
    } else {
        let v = parseNumber(line)
        l = .n(v)
    }
    
    // expect ,
    if line[index] != "," { print(", oops") }
    index += 1
    
    if line[index] == "[" {
        let sn = parseSnailNumber(line)
        r = .sn(sn)
    } else {
        let v = parseNumber(line)
        r = .n(v)
    }
    
    // expect ]
    if line[index] != "]" { print("] oops") }
    index += 1

    return SnailNumber(l: l, r: r)
}

func readNumbers() -> [SnailNumber] {
    var result: [SnailNumber] = []

    while true {
        if let line = readLine() {
            let n = parseSnailNumber(line.map{ String($0) })
            result.append(n)
        } else {
            break
        }
    }
    
    return result
}

func part1(_ numbers: [SnailNumber]) {
    var sum =  SnailNumber(l: .n(0), r: .n(0))

    for n in numbers {
        sum = sum + n
        sum = sum.reduce()
    }
    
    print(sum.magnitude)
    // let a = SnailNumber(l: .n(9), r: .n(1))
    // let b = SnailNumber(l: .n(1), r: .n(9))
    // let c = SnailNumber(l: .sn(a), r: .sn(b))
    // print(c.magnitude)
}

func main() {
    let numbers = readNumbers()
    part1(numbers)
}

main()

// Local Variables:
// compile-command: "swiftc -o d18 d18.swift"
// End:
