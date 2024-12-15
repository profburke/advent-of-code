import Foundation

enum Hands: Int {
    case high
    case pair
    case twoPair
    case three
    case full
    case four
    case five
}
    
var entries: [Entry] = []

while true {
    guard let line = readLine(strippingNewline: true) else { break }
    let parts = line.split(separator: " ")
    let e = Entry(hand: String(parts[0]), bid: Int(parts[1])!)
    entries.append(e)
}

func compare(a: Entry, b: Entry) -> Bool {
    switch (a.type, b.type) {
    case (.five, .five):
        return true
    case (.five, _):
        return false
    case (.four, .four):
        return true
    case (.four, _):
        return false
    case (.full, .full):
        return true
    case (.full, _):
        return false
    case (.three, .three):
        return true
    case (.three, _):
        return false
    case (.twoPair, .twoPair):
        return true
    case (.twoPair, _):
        return false
    case (.pair, .pair):
        return true
    case (.pair, _):
        return false
    case (.high, .high):
        return true
    case (.high, _):
        return true
    }
}

entries.sort(by: { $0.hand < $1.hand })

entries.forEach { e in
                  print(e)
}

var total = 0
for (i, e) in entries.enumerated() {
    total += e.bid * (i+1)
}

print("Part 1 - \(total)")

struct Entry {
    let hand: String
    let bid: Int

    var type: Hands {
        return .five
    }
}
