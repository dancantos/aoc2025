def read(filename):
    with open(filename) as f:
        return [ [ 0 if c == '.' else 1 for c in row[:-1] ] for row in f ]

def puzzle1(grid):
    cols = len(grid)
    rows = len(grid[0])
    result = 0
    for y in range(cols):
        for x in range(rows):
            if grid[y][x] == 0:
                continue
            count = 0
            if x > 0:
                count += grid[y][x-1]
            if x < rows-1:
                count += grid[y][x+1]
            if y > 0:
                count += grid[y-1][x]
            if y < cols-1:
                count += grid[y+1][x]

            if x > 0 and y > 0:
                count += grid[y-1][x-1]
            if x < rows-1 and y > 0:
                count += grid[y-1][x+1]
            if x > 0 and y < cols-1:
                count += grid[y+1][x-1]
            if x < rows-1 and y < cols-1:
                count += grid[y+1][x+1]

            if count < 4:
                result += 1

    return result

def puzzle2(grid):
    cols = len(grid)
    rows = len(grid[0])

    def update(grid):
        remCount = 0
        newGrid = [ [ d for d in row ] for row in grid ]
        for y in range(cols):
            for x in range(rows):
                if grid[y][x] == 0:
                    continue
                count = 0
                if x > 0:
                    count += grid[y][x-1]
                if x < rows-1:
                    count += grid[y][x+1]
                if y > 0:
                    count += grid[y-1][x]
                if y < cols-1:
                    count += grid[y+1][x]

                if x > 0 and y > 0:
                    count += grid[y-1][x-1]
                if x < rows-1 and y > 0:
                    count += grid[y-1][x+1]
                if x > 0 and y < cols-1:
                    count += grid[y+1][x-1]
                if x < rows-1 and y < cols-1:
                    count += grid[y+1][x+1]

                if count < 4:
                    newGrid[y][x] = 0
                    remCount += 1
        return remCount, newGrid

    finalCount = 0
    while True:
        remCount, grid = update(grid)
        if remCount == 0:
            break
        finalCount += remCount
    return finalCount

def printgrid(grid):
    for row in grid:
        print(''.join([ '.' if digit == 0 else '0' for digit in row ]))


grid = read("input/day4/puzzle.txt")
# print(grid)
print(puzzle1(grid))
print(puzzle2(grid))