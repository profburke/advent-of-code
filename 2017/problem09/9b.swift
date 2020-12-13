#!/usr/bin/swift
import Foundation

if CommandLine.argc != 2 {
    print("Usage: \(CommandLine.arguments[0]) <inputfile>")
    exit(1)
}

let fileName = CommandLine.arguments[1]
guard let data = try? String(contentsOfFile: fileName).dropLast() else {
    print("Couldn't open \(fileName)")
    exit(1)
}

var score = 0
var nesting = 0
var inGarbage = false
var skip = false
var garbageCount = 0

data.forEach { char in
               if skip {
                   skip = false
               } else {
                   switch char {
                   case "{":
                       if !inGarbage {
                           nesting += 1
                           score += nesting
                       } else {
                           garbageCount += 1
                       }
                   case "}":
                       if !inGarbage {
                           nesting -= 1
                       } else {
                           garbageCount += 1
                       }
                   case "<":
                       if inGarbage {
                           garbageCount += 1
                       }
                       inGarbage = true
                   case ">":
                       inGarbage = false
                   case "!":
                       skip = true
                   default: // do nothing
                       if inGarbage { garbageCount += 1 }
                   }
               }
}

print(garbageCount)
