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

data.forEach { char in
               if skip {
                   skip = false
               } else {
                   switch char {
                   case "{":
                       if !inGarbage {
                           nesting += 1
                           score += nesting
                       }
                   case "}":
                       if !inGarbage {
                           nesting -= 1
                       }
                   case "<":
                       inGarbage = true
                   case ">":
                       inGarbage = false
                   case "!":
                       skip = true
                   default: // do nothing
                       ()
                   }
               }
}

print(score)
