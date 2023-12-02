use regex::Regex;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
}

fn part1() {
    if let Ok(lines) = read_lines("../input.txt") {
        let re = Regex::new(r"[A-Za-z]").unwrap();
        let mut sum = 0;
        
        for line in lines {
            if let Ok(stuff) = line {
                let all_digits = re.replace_all(&stuff, "");
                match all_digits.chars().nth(0) {
                    None => println!("gack 1"),
                    Some(f) => {
                        match all_digits.chars().nth(all_digits.len()-1) {
                            None => println!("gack 2"),
                            Some(l) => {
                                let border_digits: String = vec![f, l].into_iter().collect();
                                let val: i64 = border_digits.parse().unwrap();
                                sum += val;
                            }
                        }
                    }
                }
            }
        }

        println!("Sum is {}", sum)
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }
