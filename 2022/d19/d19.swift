#!/usr/bin/env swift

import Darwin

struct Element: String, Hashable {
case ore
case clay
case obsidian
case geode
}

struct RobotCost {
    let costs: [Element:Int]
}

struct Blueprint {
    let costs: [Element:RobotCost]
}

func readPlans(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Cannot read \(plan)")
        exit(1)
    }

    while let line = readLine() {
        
    }
}
