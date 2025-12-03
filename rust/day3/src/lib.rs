use aocio::read_lines;

pub fn puzzle<F : Fn(&[i64]) -> i64>(banks: Vec<Vec<i64>>, f: F) -> i64 {
    let mut sum: i64 = 0;
    for bank in banks {
        sum = sum + f(&bank);
    }
    return sum;
}

pub fn largest_joltage1(bank: &[i64]) -> i64 {
    let mut high: i64 = bank[0];
    let mut low: i64 = bank[1];
    for i in 1..(bank.len()-1) {
        if bank[i] > high {
            (high, low) = (bank[i], bank[i+1]);
        } else if bank[i] > low {
            low = bank[i];
        }
    }
    if bank[bank.len()-1] > low {
        low = bank[bank.len()-1];
    }
    return 10*high + low;
}

pub fn largest_joltage2(bank: &[i64]) -> i64 {
    return _largest_joltage2(bank, 12, 0)
}

fn _largest_joltage2(bank: &[i64], size: u32, result: i64) -> i64 {
    if size == 0 {
        return result
    }
    let mut high = bank[0];
    let mut index = 0;
    for i in 1..=bank.len()-(size as usize) {
        if bank[i] > high {
            high = bank[i];
            index = i;
        }
    }
    return _largest_joltage2(&bank[index+1..], size-1, result + high*(10_i64.pow(size-1)))
}

pub fn read(filename: String) -> Vec<Vec<i64>> {
    let mut result: Vec<Vec<i64>> = Vec::new();

    match read_lines(filename) {
        Err(err) => panic!("{}", err),
        Ok(lines) => {
            for line in lines.map_while(Result::ok) {
                let mut row: Vec<i64> = Vec::new();
                line.trim_end_matches("\n").chars().for_each(|c| {
                    row.push(c.to_digit(10).unwrap() as i64);
                });
                result.push(row);
            }
        },
    };
    return result;
}