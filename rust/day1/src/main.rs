use aocio::{read_lines};

use day1::{count_zero, count_zero_passes};

fn main() {
    let moves = read_input("input/day1/puzzle.txt");
    println!("puzzle 1: {}", count_zero(&moves));
    println!("puzzle 2: {}", count_zero_passes(&moves));
}

fn read_input(filename: &str) -> Vec<i32> {
    let mut moves: Vec<i32> = Vec::new();
    if let Ok(lines) = read_lines(filename) {
        for line in lines.map_while(Result::ok) {
            let amount = line[1..].parse::<i32>()
                .expect("should be an integer");
            if line.starts_with('R') {
                moves.push(amount);
            } else {
                moves.push(-amount);
            }
        }
    }
    return moves
}
