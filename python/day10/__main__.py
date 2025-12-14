import numpy as np

def read(filename):
    with open(filename) as f:
        return [ process_line(line) for line in f ]

def process_line(line):
    components = line.strip().split()
    return parse_lights(components[0][1:-1]), [ parse_button(button) for button in components[1:-1] ], parse_joltage(components[-1])

def parse_lights(lights):
    fn = lambda char: 1 if char == '#' else 0
    num = 0
    for i, char in enumerate(lights):
        num += fn(char)<<i
    return num

def parse_button(button):
    result = 0
    for num in button[1:-1].split(','):
        result += 1<<int(num)
    return result

def parse_joltage(joltage):
    result = []
    for num in joltage[1:-1].split(','):
        result.append(int(num))
    return result

def press(button, lights):
    return button ^ lights

def solve_single1(buttons, target):
    # solutions = []
    minimum = len(buttons)
    for i in range(1, 1<<len(buttons)):
        lights = 0
        count = 0
        for j, b in enumerate(buttons):
            if i&(1<<j):
                lights ^= b
                count += 1
        if lights == target and count < minimum:
            minimum = count

    return minimum

def solve_single2(buttons, target_lights, target_joltage):
    print(target_joltage)
    arr = []
    for b in buttons:
        row = [ 1 if b & 1<<pos else 0 for pos in range(len(target_joltage))]
        arr.append(row)
    mat = np.array(arr).transpose()
    target = np.array(target_joltage)
    print(np.linalg.lstsq(mat, target))

def solve(problems, solver):
    return sum([ solver(p[1], p[0]) for p in problems ])

# def count_digits(n):
#     count = 0
#     while n > 0:
#         n &= (n - 1)
#         count += 1
#     return count


p = read("input/day10/sample.txt")
# p = read("input/day10/puzzle.txt")
# print(p)
# print(solve(p, solve_single1))
# print(solve(p, solve_single2))
solve_single2(p[0][1], p[0][0], p[0][2])