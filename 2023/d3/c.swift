import Foundation

//
// Print a list of all unique characters in file fed in via standard input
//
var cs: [Character: Int] = [:]

while true {
    guard let line = readLine(strippingNewline: true) else { break }

    line.forEach { c in
                   let count = cs[c, default: 0]
                   cs[c] = count + 1
    }
}

print(cs.keys.sorted())


    
