def read(filename):
    with open(filename) as f:
        return [ [ int(char) for char in row[:-1] ] for row in f ]

def find_largest1(bank):
    high, low = bank[0], bank[1]
    for i in range(1, len(bank)):
        if i < len(bank) - 1 and bank[i] > high:
            high, low = bank[i], bank[i+1]
        elif bank[i] > low:
            low = bank[i]
    return high*10 + low

def find_largest2(bank):
    digits = _find_largest2(bank, size=12, digits=[])
    sum = 0
    for pos in range(len(digits)):
        n = digits[pos]
        sum += n*(10**(11-pos))
    return sum

def _find_largest2(bank, size=12, digits=[]):
    if size == 0:
        return digits
    if size == 1:
        digits.append(max(bank))
        return digits
    high, index = bank[0], 0
    for i in range(len(bank[1:-size+1])):
        if bank[i+1] > high:
            high, index = bank[i+1], i+1
    digits.append(high)
    # print(bank, bank[:-size+1], high, index, size)
    return _find_largest2(bank[index+1:], size-1, digits)

def puzzle(banks, fn):
    sum = 0
    for bank in banks:
        sum += fn(bank)
    return sum

data = read("input/day3/sample.txt")
# print(data)
print(puzzle(data, find_largest1))
print(puzzle(data, find_largest2))
# print(find_largest2(data[1]))


# 3121910778619
# 3121190777711


# 987654321111
# 811111111119
# 434234234278
# 888911112111

# 987654321111
# 811111111119
# 434234234278
# 888911112111