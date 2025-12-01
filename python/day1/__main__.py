import sys

def parse_input(file):
    sig = lambda char : 1 if char == 'R' else -1
    with open(file, "r") as f:
        return [ -1*sign*num for sign, num in [ (sig(line[0]), int(line[1:])) for line in f if len(line.strip()) > 0 ] ]

def count_zero(moves, start=50, count=0):
    for turn in moves:
        start = (start+turn) % 100
        count += start == 0
    return count

def count_zero_passes(moves, start=50, count=0):
    for turn in moves:
        next = start + turn
                      # full turns past 100       landed on 0    started from 0 and went negative (overcounted)
        count +=    abs((next-(next<0))//100)  +  (next == 0)   -   (start == 0 and turn < 0)
        start = next%100
    return count

args = sys.argv[1:]
if len(args) == 0:
    file = "input/day1/puzzle.txt"
else:
    file = args[0]

moves = parse_input(file)
print("part 1:", count_zero(moves))
print("part 2:", count_zero_passes(moves))
