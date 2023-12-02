#!/usr/bin/env swift

import Darwin

class File {
    let name: String
    let size: Int

    init(_ name: String, _ size: Int) {
        self.name = name
        self.size = size
    }
}

class Directory {
    let name: String
    let path: String
    var files: [String:File]
    var directories: [String:Directory]
    let parent: Directory?
    var size: Int = 0
    
    init(_ name: String, _ parent: Directory?) {
        self.name = name
        self.files = [:]
        self.directories = [:]
        self.parent = parent

        if let p = parent {
            path = p.path + name
        } else {
            path = name
        }
    }
}

func computeSize(_ d: Directory) -> Int {
    let fSize = d.files.values.map { $0.size }.reduce(0,  +)
    let dSize = d.directories.values.map { computeSize($0) }.reduce(0, +)
    d.size = fSize + dSize
    return d.size
}

func part1(path: String) {
    guard let _ = freopen(path, "r", stdin) else {
        print("Couldn't open file: \(path)")
        return
    }

    let tree = Directory("/", nil)
    var pwd = tree
    var allDirectories: [String:Directory] = ["/": tree]
    
    while let line = readLine() {
        if line.prefix(5) == "$ cd " {
            let index = line.lastIndex(of: " ")!
            let i = line.index(after: index)
            let dName = String(line[i...])

            if dName == ".." {
                pwd = pwd.parent!
            } else {
                if dName == "/" {
                    pwd = tree
                } else {
                    let d = pwd.directories[dName]!
                    pwd = d
                }
            }
        } else if line == "$ ls" {
            // skip
        } else {
            let parts = line.split(separator: " ")
            if parts[0] == "dir" {
                let n = String(parts[1])
                if pwd.directories[n] == nil {
                    // create the directory entry
                    let d = Directory(n, pwd)
                    pwd.directories[d.name] = d

                    if allDirectories[d.path] == nil {
                        allDirectories[d.path] = d
                    }
                }
            } else {
                let n = String(parts[1])
                if pwd.files[n] == nil {
                    // create the file
                    let f = File(n, Int(parts[0])!)
                    pwd.files[f.name] = f
                }
            }
        }
    }

    _ = computeSize(tree)

    let total = allDirectories.values.filter { $0.size <= 100_000 }.map { $0.size }.reduce(0, +)
    print(total)

    let freeSpace = 70_000_000 - tree.size
    let needed = 30_000_000 - freeSpace

    let allSorted = allDirectories.values.sorted() { f, g in f.size < g.size }

    for dir in allSorted {
        if dir.size >= needed {
            print(needed, dir.size)
            break
        }
    }
}

part1(path: CommandLine.arguments[1])
