use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let args = &env::args().collect::<Vec<String>>();
    let filename: &String = &args[1];

    if let Ok(mut lines) = read_lines(filename)  {
        let times = match lines.nth(0).unwrap() {
            Ok(line) => line.clone().split(":").collect::<Vec<&str>>()[1],
            _ => "",
        };
        
        let distances = match lines.nth(0).unwrap() {
            Ok(line) => line.split(":").collect::<Vec<&str>>()[1],
            _ => "",
        };
            
        println!("{} {}", times, distances);
    }
}



fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }


struct Race {
    time: i64,
    max_d: i64
}
