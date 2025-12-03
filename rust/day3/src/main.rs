use day3::{read, puzzle, largest_joltage1, largest_joltage2};

fn main() {
    let input = read("input/day3/puzzle.txt".to_string());
    println!("puzzle 1: {}", puzzle(&input, largest_joltage1));
    println!("puzzle 2: {}", puzzle(&input, largest_joltage2));
}
