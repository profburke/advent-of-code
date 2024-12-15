import Foundation

var sum = 0
let steps = readLine(strippingNewline: true)!.split(separator: ",").map { String($0) }

steps.forEach { step in sum += hash(step) }

print("Part 1 - \(sum)")

func hash(_ s: String) -> Int {
    var result = 0

    s.forEach { c in
                result += Int(c.asciiValue!)
                result *= 17
                result %= 256
    }
    
    return result
}
