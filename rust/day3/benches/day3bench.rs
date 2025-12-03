use criterion::{Criterion, criterion_group, criterion_main, BatchSize};

use day3::{puzzle, largest_joltage1, largest_joltage2, read};

// bench_puzzle1           time:   [21.419 µs 21.946 µs 22.429 µs]
// bench_puzzle2           time:   [48.139 µs 48.498 µs 48.903 µs]
fn bench_puzzle1(c: &mut Criterion) {
    let input = read("../../input/day3/puzzle.txt".to_string());
    c.bench_function("bench_puzzle1", |b| {
        b.iter_batched(|| input.clone(), |input| puzzle(input, largest_joltage1), BatchSize::SmallInput);
    });

    c.bench_function("bench_puzzle2", |b| {
        b.iter_batched(|| input.clone(), |input| puzzle(input, largest_joltage2), BatchSize::SmallInput);
    });
}

criterion_group!(
    benches,
    bench_puzzle1,
);
criterion_main!(benches);
