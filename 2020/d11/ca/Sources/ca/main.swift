import Foundation

public class CA {
    public enum CellState: String, CaseIterable {
    case Floor = "."
    case Occupied = "#"
    case Empty = "L"
    }

    public let rows: Int
    public let columns: Int
    
    private var currentGrid = 0
    private var nextGrid = 1
    
    private var world: [[[CellState]]]

    private func onGrid(_ r: Int, _ c: Int) -> Bool {
        return r >= 0 && c >= 0 && r < rows && c < columns
    }
    
    private func improvedNeighbors(_ r: Int, _ c: Int) -> Int {
        var n = 0
        let deltas = [ (1, 0), (1, 1), (1, -1),(0, 1),
                        (0, -1), (-1, 1), (-1, 0), (-1, -1)]

        for delta in deltas {
            var x = r
            var y = c
            
            while true {
                (x, y) = (x + delta.0, y + delta.1)
                if !onGrid(x, y) {
                    break
                }

                if get(x, y) == .Empty {
                    break
                }
                
                if get(x, y) == .Occupied {
                    n += 1
                    break
                }
            }
        }
        
        return n
    }
    
    private func neighbors(_ r: Int, _ c: Int) -> Int {
        var n = 0
        
        for dr in -1...1 {
            for dc in -1...1 {
                if (dr == 0) && (dc == 0) { continue }
                if (r + dr < 0) || (r + dr >= rows) { continue }
                if (c + dc < 0) || (c + dc >= columns) { continue }
                let current = get(r + dr, c + dc)
                if (current == .Occupied) {
                    n += 1
                }
            }
        }

        return n
    }

    private func improvedGetNewState(current: CellState, n: Int) -> CellState {
        switch current {
        case .Floor:
            return .Floor
        case .Occupied:
            if n >= 5 {
                return .Empty
            } else {
                return .Occupied
            }
        case .Empty:
            if n == 0 {
                return .Occupied
            } else {
                return .Empty
            }
        }
    }
    
    private func getNewState(current: CellState, n: Int) -> CellState {
        switch current {
        case .Floor:
            return .Floor
        case .Occupied:
            if n >= 4 {
                return .Empty
            } else {
                return .Occupied
            }
        case .Empty:
            if n == 0 {
                return .Occupied
            } else {
                return .Empty
            }
        }
    }

    private func newvalue(_ r: Int, _ c: Int, _ n: Int) {
        let curval = get(r, c)

        let newState = improvedGetNewState(current: curval, n: n)
        set(r, c, newState)
        
}

    private func swapGrids() {
        (currentGrid, nextGrid) = (nextGrid, currentGrid)
    }

    public func step() {
        for r in 0..<rows {
            for c in 0..<columns {
                let n = improvedNeighbors(r, c)
                newvalue(r, c, n)
            }
        }

        swapGrids()
    }

    public var stable: Bool {
        for r in 0..<rows {
            for c in 0..<columns {
                if world[currentGrid][r][c] != world[nextGrid][r][c] {
                    return false
                }
            }
        }
        
        return true
    }
    
    public func get(_ r: Int, _ c: Int) -> CellState {
        return world[currentGrid][r][c]
    }

    public func set(_ r: Int, _ c: Int, _ v: CellState) {
        world[nextGrid][r][c] = v
    }

    public func count(state: CellState) -> Int {
        var count = 0
        
        for r in 0..<rows {
            for c in 0..<columns {
                if get(r, c) == state {
                    count += 1
                }
            }
        }

        return count
    }
    
    
    public init(_ rows: Int, _ columns: Int) {
        self.rows = rows
        self.columns = columns
        world = Array(repeating: Array(repeating: Array(repeating: .Floor, count: columns), count: rows), count: 2)
    }

    public static func fromFile(_ fileName: String) -> CA? {
        guard let data = try? String(contentsOfFile: fileName) else {
            return nil
        }

        let lines = data.split(separator: "\n")
        let rows = lines.count
        let columns = String(lines[0]).count

        let ca = CA(rows, columns)

        for (row, line) in lines.enumerated() {
            for (column, char) in Array(line).enumerated() {
                guard let state = CellState(rawValue: String(char)) else {
                    return nil
                }
                ca.set(row, column, state)
            }
        }
        ca.swapGrids()
        
        return ca
    }
}

extension CA: CustomStringConvertible {
    public var description: String {
        var result = ""
        
        for r in 0..<rows {
            for c in 0..<columns {
                result += "\(get(r, c).rawValue)"
            }
            result += "\n"
        }
        
        return result
    }
}

//guard let g = CA.fromFile("test") else {
guard let g = CA.fromFile("input.txt") else {
    print("blech")
    exit(1)
}

var gen = 0

while true {
    print("Gen: \(gen)")
    //print(g)
    g.step()
    if g.stable {
        print("stable")
        break
    }

    gen += 1
}

print(g.count(state: .Occupied))



