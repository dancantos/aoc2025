def read(filename):
    problems = []
    with open(filename) as f:
        chars = [ [ c for c in line.strip().split() ] for line in f ]

    # separate rows of numbers from operators
    nums = [ [ int(n) for n in row ] for row in chars[:-1] ]
    ops = chars[-1]

    # transpose
    problems = [ [ nums[i][j] for i in range(len(nums)) ] for j in range(len(nums[0])) ]
    return problems, ops

def compute(nums, op):
    if op == '*':
        fn = lambda x, y: x*y
        init = 1
    if op == '+':
        fn = lambda x, y: x+y
        init = 0
    if op == '-':
        fn = lambda x, y: x-y
        init = 0

    for n in nums:
        init = fn(init, n)
    return init

def str_to_num(s):
    n = 0
    size = len(s)
    for i in range(size-1, -1, -1):
        n += int(s[i]) * 10**i
    return n

def read2(filename):
    with open(filename) as f:
        lines = [ line for line in f ]

    # pad lines
    maxlen = max([ len(l) for l in lines ])
    lines = [ line+' '*(maxlen-len(line)) for line in lines ]

    # whitespace between digits now matters, so find the starting indices of each problem
    start_pos = []
    for i in range(1, len(lines[-1])):
        if lines[-1][i] != ' ':
            start_pos.append(i)

    # extract the lines for each problem by iterating over the starting positions
    problems = []
    col_start = 0
    for i in range(len(start_pos)):
        col_end = start_pos[i]-1
        problems.append([ line[col_start:col_end] for line in lines[:-1] ])
        col_start = start_pos[i]
    # last line
    problems.append([ line[col_start:-1] for line in lines[:-1] ])

    ops = lines[-1].strip().split()
    # for each problem: transpose -> reverse -> to_num gives us the numbers for each problem
    # explanation: if we treat each problem array as a matrix of characters, then rotating the matrix gives us the numbers we need
    # the inputs are arrays of strings, so we need to index into them, then reconstruct the results as strings as we go.
    # we then wrap up the computation and sum the results as we go
    return sum([ compute([ str_to_num(''.join([ p[i][j] for i in range(len(p)) ]).strip()[::-1]) for j in range(len(p[0])) ], ops[pi]) for pi, p in enumerate(problems) ])

def puzzle(problems, ops):
    result = 0
    for i in range(len(problems)):
        result += compute(problems[i], ops[i])
    return result

p, o = read("input/day6/puzzle.txt")
# p2, o2 = read2("input/day6/puzzle.txt")
print("Puzzle 1: ", puzzle(p, o))
print("Puzzle 2: ", read2("input/day6/puzzle.txt"))
