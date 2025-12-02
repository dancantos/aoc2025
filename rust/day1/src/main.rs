use aocio::{read_lines};

fn main() {
    let moves = read_input("input/day1/puzzle.txt");
    println!("puzzle 1: {}", count_zero(&moves));
    println!("puzzle 2: {}", count_zero_passes(&moves));
}

fn count_zero(moves: &[i32]) -> i32 {
    let mut current: i32 = 50;
    let mut count: i32 = 0;
    for mv in moves {
        current = (current + mv) % 100;
        if current == 0 {
            count += 1;
        }
    }
    return count;
}

fn count_zero_passes(moves: &[i32]) -> i32 {
    let mut current: i32 = 50;
    let mut count: i32 = 0;

    for mv in moves {
        // prepare increment
        let next = current+mv;

        // main update
        count = match next.signum() {
            1 => count + next / 100,        // number of times past +100
            -1 => count - (next / 100) + 1, // number of times past -100 including 0
            0 => count + 1,                 // landed on 0
            _ => panic!("uh oh")
        };

        // started from 0 and went negative (overcounted)
        if current == 0 && *mv < 0 {
            count = count-1;
        }

        // complete increment
        current = next.rem_euclid(100);
    }
    return count;
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
