import Foundation

var sequences: [[Int]] = []

while true {
    guard let line = readLine(strippingNewline: true) else { break }
    let sequence = line.split(separator: " ").map { Int(String($0))! }
    sequences.append(sequence)
}

func doit(_ sequences: [[Int]]) -> Int {
    var sum = 0
    sequences.forEach { seq in
                        var cur = seq
                        var dseqs: [[Int]] = [cur]

                        while !cur.allSatisfy({ $0 == 0 })  {
                            var newSeq: [Int] = []

                            for i in 1..<cur.count {
                                newSeq.append(cur[i] - cur[i-1])
                            }

                            dseqs.append(newSeq)
                            cur = newSeq
                        }

                        for i in (0..<(dseqs.count-1)).reversed() {
                            var seq = dseqs[i]
                            let dseq = dseqs[i+1]
                            seq.append( seq[seq.count-1] + dseq[dseq.count-1])
                            dseqs[i] = seq
                        }

                        let s = dseqs[0]
                        let val = s[s.count-1]
                        sum += val
    }

    return sum
}

func part1() {
    let result = doit(sequences)
    print("Part 1 - \(result)")
}

func part2() {
    var newSeqs: [[Int]] = []

    sequences.forEach { seq in
                        newSeqs.append(seq.reversed())
    }

    let result = doit(newSeqs)
    print("Part 2 - \(result)")
   
}

part1()
part2()
