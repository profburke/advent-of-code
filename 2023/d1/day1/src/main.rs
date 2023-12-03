use regex::Regex;
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
    let re = Regex::new(r"[A-Za-z]").unwrap();
    let mut sum = 0;
    
    for line in lines {
        if let Ok(line) = line {
            let val = part1_val(line, &re);
            sum += val;
        }
    }
    
    println!("Part 1 is {}", sum)
}

fn part1_val(line: String, re: &Regex) -> i64 {
    let line = re.replace_all(&line, "").to_string();
    let f = line.chars().nth(0).unwrap();
    let l = line.chars().nth(line.len()-1).unwrap();
    let val: i64 = vec![f, l].into_iter().collect::<String>().parse().unwrap();

    return val;
}

fn part2(filename: &String) {
    if let Ok(lines) = read_lines(filename) {
        let re = Regex::new(r"[A-Za-z]").unwrap();
        let mut sum = 0;
        
        for line in lines {
            if let Ok(mut line) = line {
                line = line.replace("one", "1");
                line = line.replace("two", "2");
                line = line.replace("three", "3");
                line = line.replace("four", "4");
                line = line.replace("five", "5");
                line = line.replace("six", "6");
                line = line.replace("seven", "7");
                line = line.replace("eight", "8");
                line = line.replace("nine", "9");

                println!("Mangled line '{}'", line);
                
                let all_digits = re.replace_all(&line, "");
                let f = all_digits.chars().nth(0).unwrap();
                let l = all_digits.chars().nth(all_digits.len()-1).unwrap();
                let val: i64 = vec![f, l].into_iter().collect::<String>().parse().unwrap();
                println!("Val is {}", val);
                sum += val;
            }
        }

        println!("Part 2 is {}", sum)
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }
