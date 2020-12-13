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

class Entry {
    let name: String
    let weight: Int
    var towerWeight: Int
    var children: [Entry]
    var parent: Entry?

    init(name: String, weight: Int, children: [Entry] = [], parent: Entry? = nil) {
        self.name = name
        self.weight = weight
        self.children = children
        self.parent = parent
        towerWeight = weight
    }
}

var database: [String : Entry] = [:]

func parse(_ line: String) -> Entry {
    let components = line.split(separator: " ")
    let name = String(components[0])
    let weight = Int(String(components[1].dropFirst().dropLast())) ?? -1
    let entry = Entry(name: name, weight: weight)
    return entry
}

func wire(_ line: String) {
    let components = line.split(separator: " ")
    guard components.count > 2 else { return }
    
    let name = String(components[0])
    if let parent = database[name] {
        for index in 3..<components.count { // assuming line is formatted correctly and that words 3, 4, ... are child names
            let childName = components[index].filter { $0 != "," }
            if let child = database[childName] {
                child.parent = parent
                parent.children.append(child)
            } else {
                print("Weird, can't find child \(childName)")
            }
        }
    } else {
        print("Weird, can't find \(name)")
        exit(1)
    }
}

var lines = data.split(separator: "\n", maxSplits: Int.max,
                        omittingEmptySubsequences: true)

// First create all the entries
for line in lines {
    let entry = parse(String(line))
    database[entry.name] = entry
}

// Now wire up the children and parents
for line in lines {
    wire(String(line))
}

var current: Entry? = database.first!.value
while current!.parent != nil { current = current!.parent }
let root = current!
print("Root: \(root.name)")
