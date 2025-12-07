def read(filename):
    start = 0
    with open(filename) as f:
        return [ [ char for char in line.strip() ] for line in f ]

    # for i in range(len(lines[0])):
    #     if lines[0][i] == 'S':
    #         start = i

    # return start, lines

def puzzle1(lines):
    splitcount = 0
    for i in range(1, len(lines)):
        for j in range(len(lines[i])):
            char = lines[i][j]

            # split
            if char == '^' and lines[i-1][j] == '|':
                lines[i][j-1] = '|'
                lines[i][j+1] = '|'
                splitcount += 1

            # carry
            if char == '.' and (lines[i-1][j] == '|' or lines[i-1][j] == 'S'):
                lines[i][j] = '|'
    return splitcount

def puzzle2(lines):
    multiplicity_prev = [0] * len(lines[0])
    multiplicity_current = [0] * len(lines[0])
    for i in range(1, len(lines)):
        for j in range(len(lines[i])):
            char = lines[i][j]
            if lines[i-1][j] == 'S':
                lines[i][j] = '|'
                multiplicity_current[j] = 1

            # carry
            if (char == '.' or char == '|') and lines[i-1][j] == '|':
                lines[i][j] = '|'
                multiplicity_current[j] += multiplicity_prev[j]

            # split
            if char == '^' and lines[i-1][j] == '|':
                lines[i][j-1] = '|'
                lines[i][j+1] = '|'
                multiplicity_current[j-1] += multiplicity_prev[j]
                multiplicity_current[j+1] += multiplicity_prev[j]
        # print(i, multiplicity_current)
        multiplicity_prev = multiplicity_current
        multiplicity_current = [0] * len(lines[0])

    return sum(multiplicity_prev)

def printlines(lines):
    for line in lines:
        print(''.join(line))

# lines = read("input/day7/sample.txt")
lines = read("input/day7/puzzle.txt")
# for line in lines:
#     print(line)
print(puzzle1(lines))
# print(puzzle2(lines))
# printlines(lines)



# ['.', '.', '|', '.', '|', '.', '|', '|', '|', '.', '|', '.', '|', '.', '.'] [0, 0, 1, 0, 5, 0, 4, 3, 4, 0, 2, 0, 1, 0, 0]
# ['.', '|', '^', '|', '|', '.', '^', '.', '.', '.', '.', '.', '^', '.', '.'] [0, 1, 0, 1, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

# ['.', '.', '|', '.', '|', '.', '|', '|', '|', '.', '|', '.', '|', '.', '.'] [0, 0, 1, 0, 5, 0, 4, 3, 4, 0, 2, 0, 1, 0, 0]
# ['.', '|', '^', '|', '|', '|', '^', '|', '.', '.', '.', '.', '^', '.', '.'] [0, 1, 0, 1, 5, 4, 0, 4, 0, 0, 0, 0, 0, 0, 0]