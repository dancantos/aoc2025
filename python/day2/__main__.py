import numpy as np
from typing import Callable

def read(filename):
    ranges = []
    with open(filename, "r") as f:
        contents = f.read()
        for id_range in contents.split(","):
            sep_index = id_range.index("-")
            ranges.append((int(id_range[:sep_index]), int(id_range[sep_index+1:])))
    return ranges

def invalid1(n):
    # does the number contain an even number of digits?
    digits = int(np.floor(np.log10(n)) + 1)
    if digits%2 == 1:
        return False

    # split in half and compare
    e = 10 ** (digits//2)
    return (n // e) == (n % e)

def invalid2(n):
    # strategy here is to determine if the
    digits = int(np.floor(np.log10(n)) + 1)
    for d in range(1, digits//2+1):
        # skip if digit count to be repeated doesn't divide total digits evenly
        if digits % d != 0:
            continue

        e = 10**d # magnitude to trim by
        cp = n%e  # trim to least d digits for copying

        # compute the candidate by copying the copy value
        candidate = 0
        for nn in range(0, digits, d):
            candidate += cp * 10**nn
        if candidate == n:
            return True
    return False

def invalid1_seek(r):
    # finds invalud nubmers based on range bounds
    digits = lambda n: int(np.log10(n) + 1)
    low_digits, high_digits = digits(r[0]), digits(r[1])
    for d in range(low_digits, high_digits+1):
        if d%2 == 1:
            continue

        # cut the lower d//2 digits off the lower and upper bounds
        e = 10 ** (d//2)
        low, high = r[0]//e, r[1]//e

        # for every number in this range we just append itself to the end and see if it falls in the original range
        for n in range(max(low, 1), high+1):
            if digits(n)*2 != d:
                continue
            candidate = n*e + n
            if r[0] <= candidate <= r[1]:
                yield candidate

def invalid2_seek(r):
    pass

def puzzle(ranges, invalid_fn: Callable[[int], bool]):
    count=0
    for r in ranges:
        for i in range(r[0], r[1]+1):
            if invalid_fn(i):
                count += i
    return count

def puzzle1_seek(ranges):
    count=0
    for r in ranges:
        for n in invalid1_seek(r):
            count += n
    return count

input = read("input/day2/puzzle.txt")
print("puzzle 1:        ", puzzle(input, invalid1))
print("puzzle 2:        ", puzzle(input, invalid2))
# print("puzzle 1 (fast): ", puzzle1_seek(input))
