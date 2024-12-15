use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let args = &env::args().collect::<Vec<String>>();
    let filename = if args.len() == 2 { &args[1] } else { &args[2] };
    let part: i64 = if args.len() == 2 { 1 } else { args[1].parse().unwrap() };

    match part {
        1 => part1(filename),
        _ => part2(filename),
    }
}

fn part1(filename: &String) {
    let lines = read_lines(filename).expect("Couldn't open file");
    let mut points = 0;
    
    for line in lines {
        if let Ok(line) = line {
            let index = line.find(":").unwrap();
            let parts = line[(index+1)..].split("|").collect::<Vec<&str>>();

            let winning_numbers: HashSet<i64> = HashSet::from_iter(to_ivec(parts[0]));
            let my_numbers = HashSet::from_iter(to_ivec(parts[1]));

            let matches: Vec<&i64> = winning_numbers.intersection(&my_numbers).collect();
            let n = matches.len();
            let card_points = if n == 0 { 0 } else { i64::pow(2, (n - 1).try_into().unwrap()) };
            points += card_points;
        }        
    }

    println!("Points = {}", points);
}

fn part2(filename: &String) { println!("TBD {filename}"); }

fn to_ivec(s: &str) -> Vec<i64> {
    let frags = s.trim().split_whitespace().collect::<Vec<&str>>();
    let result: Vec<i64> = frags.into_iter().map( |frag| { frag.parse().unwrap() }).collect();

    return result;
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }
