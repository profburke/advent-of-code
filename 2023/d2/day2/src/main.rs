use std::env;
use std::cmp::max;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let args = &env::args().collect::<Vec<String>>();
    let filename: &String = &args[1];
    
    let game = Game {
        red: args[2].parse().unwrap(),
        green: args[3].parse().unwrap(),
        blue: args[4].parse().unwrap(),
    };
    
    part1(filename.to_string(), game);
    part2(filename.to_string());
}

fn part1(filename: String, game: Game) {
    let mut sum = 0;
    
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            let line = line.unwrap();
            let mut ok = true;

            let parts = line.split(":").collect::<Vec<&str>>();
            let game_phrase = parts[0].replace("Game ", "");
            let game_id: i64 = game_phrase.parse().unwrap();
            
            let demos = parts[1].split(";");

            'demos: for demo in demos {
                let samples = demo.split(",");
                
                for sample in samples {
                    let chunks = sample.trim().split(" ").collect::<Vec<&str>>();
                    let n: i64 = chunks[0].trim().parse().unwrap();
                    
                    match chunks[1] {
                        "red" => {
                            if n > game.red { ok = false; break 'demos; }
                        },
                        "green" => {
                            if n > game.green { ok = false; break 'demos; }
                        },
                        "blue" => {
                            if n > game.blue { ok = false; break 'demos; }
                        },
                        _ => {}
                    }
                }
            }
            
            if ok {
                println!("Adding {}", game_id);
                sum += game_id;
            }
        }
    }

    println!("Part 1 is {}", sum);
}
    
fn part2(filename: String) {
    let mut sum = 0;
    
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            let line = line.unwrap();
            let mut red = 0;
            let mut green = 0;
            let mut blue = 0;
            
            let parts = line.split(":").collect::<Vec<&str>>();
            let demos = parts[1].split(";");

            for demo in demos {
                let samples = demo.split(",");
                
                for sample in samples {
                    let chunks = sample.trim().split(" ").collect::<Vec<&str>>();
                    let n: i64 = chunks[0].trim().parse().unwrap();
                    
                    match chunks[1] {
                        "red" => {
                            red = max(red, n);
                        },
                        "green" => {
                            green = max(green, n);
                        },
                        "blue" => {
                            blue = max(blue, n);
                        },
                        _ => {}
                    }
                }
            }

            let power = red * green * blue;
            sum += power;
        }
    }

    println!("Part 2 is {}", sum);
}
    
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
        let file = File::open(filename)?;
        Ok(io::BufReader::new(file).lines())
    }


struct Game {
    red: i64,
    green: i64,
    blue: i64,
}
