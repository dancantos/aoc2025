def read(filename):
    part = 1
    ranges = []
    available = []
    with open(filename) as f:
        for line in f:
            if line == "\n":
                part = 2
                continue

            if part == 1:
                ranges.append([ int(num) for num in line.split("-") ])

            if part == 2:
                available.append(int(line))
    return ranges, available

def fresh(ranges, item):
    for r in ranges:
        if r[0] <= item and item <= r[1]:
            return True
    return False

def puzzle1(ranges, available):
    count = 0
    for item in available:
        if fresh(ranges, item):
            count+=1
    return count

def puzzle2(ranges):
    merged = [False] * len(ranges)
    def merge(r1, r2):
        # print("TEST", r1, r2)
        if r1[1] < r2[0] or r1[0] > r2[1]:
            return r1, False
        # print("MERGE", r1, r2)
        return [ min(r1[0], r2[0]), max(r1[1], r2[1]) ], True
    for i in range(1, len(ranges)):
        for j in range(i):
            if merged[j]:
                continue
            ranges[i], merged[j] = merge(ranges[i], ranges[j])
    # print(ranges)
    # print(merged)

    sum = 0
    for i in range(len(ranges)):
        r = ranges[i]
        if not merged[i]:
            sum += r[1] - r[0] + 1
    return sum

ranges, available = read("input/day5/puzzle.txt")
# print(ranges, available)
print(puzzle1(ranges, available))
print(puzzle2(ranges))