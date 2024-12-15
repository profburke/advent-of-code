use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let args = &env::args().collect::<Vec<String>>();
    let filename: &String = &args[1];

    part1(filename);
}

fn part1(filename: &String) {
    let mut row = 0;
    let mut symbols = Vec::new();
    let mut numbers = Vec::new();
    let mut start_col = 0;
    let mut val_str = "";
    let mut building = false;
    
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(line) = line {
                for (idx, c) in line.chars().enumerate() {
                    match c {
                        '.' => {
                            if building {
                                let n = Number { val: val_str.parse().unwrap(), row: row, start_col: start_col,
                                                 end_col: i64::try_from(idx).unwrap() - 1 };
                                numbers.push(n);
                                val_str = "";
                                building = false;
                            }
                        },
                        '0'..='9' => {
                            if !building {
                                building = true;
                                start_col = i64::try_from(idx).unwrap();
                            }
                            val_str ..= &c.to_string();
                        },
                        _ => {
                            let s = Symbol { c: c, row: row, col: i64::try_from(idx).unwrap()};
                            symbols.push(s);
                        },
                    }
                }
            }
            
            row += 1;
        }

        println!("{:?}", symbols);
        println!("{:?}", numbers);
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }

#[derive(Debug)]
struct Number {
    val: i64,
    row: i64,
    start_col: i64,
    end_col: i64
}

#[derive(Debug)]
struct Symbol {
    c: char,
    row: i64,
    col: i64
}
