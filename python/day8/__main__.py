from collections import defaultdict

def read(filename):
    with open(filename) as f:
        return [ [ int(num) for num in line.strip().split(',') ] for line in f ]

def puzzle1(data, connections=1000):
    sorted_dist = []
    pair_tracker = []
    for i in range(1, len(data)):
        for j in range(i):
            dx, dy, dz = (data[i][0] - data[j][0], data[i][1] - data[j][1], data[i][2] - data[j][2])
            dist_sq = dx*dx + dy*dy + dz*dz
            insert_loc = binarySearch(sorted_dist, dist_sq)
            sorted_dist.insert(insert_loc, dist_sq)
            pair_tracker.insert(insert_loc, (i, j))
            if len(sorted_dist) > connections:
                sorted_dist = sorted_dist[:connections]
                pair_tracker = pair_tracker[:connections]
            # print(j + i*len(data))

    # print(sorted_dist)
    # print(pair_tracker)
    groups = {}

    def pair(i, j):
        if j in groups and i not in groups:
            return pair(j, i)
        if i not in groups:
            groups[i] = set({i})
        # print(groups, i, j)
        if j in groups[i]:
            return False
        if j not in groups:
            groups[i].add(j)
        else:
            groups[i].update(groups[j])
        groups[j] = groups[i]
        for n in groups[i]:
            groups[n] = groups[i]
        # print(groups, i, j)
        return True

    counter = 0
    i = 0
    while True:
        if counter == connections:
            break
        if pair(pair_tracker[i][0], pair_tracker[i][1]):
            pass
            # print("PAIR", pair_tracker[i])
        # else:
        #     print("REJECT", pair_tracker[i])
        i += 1
        counter += 1

    # print(groups)

    swapped = defaultdict(set)
    for k, v in groups.items():
        swapped[frozenset(v)].add(k)
    # print(swapped)
    sizes = sorted([ len(swapped[i]) for i in swapped ], reverse=True)[:3]
    mul = 1
    for n in sizes:
        mul *= n

    return mul

def puzzle2(data):
    pairs_needed = len(data)-1
    sorted_dist = []
    pairs = []
    phase = 0
    graph = defaultdict(set)
    for i in range(1, len(data)):
        for j in range(i):
            dx, dy, dz = (data[i][0] - data[j][0], data[i][1] - data[j][1], data[i][2] - data[j][2])
            dist_sq = dx*dx + dy*dy + dz*dz
            insert_loc = binarySearch(sorted_dist, dist_sq)

            # add items until we find a spanning tree
            if phase == 0:
                sorted_dist.insert(insert_loc, dist_sq)
                pairs.insert(insert_loc, (i, j))
                graph[i].add(j)
                graph[j].add(i)
                if len(sorted_dist) >= pairs_needed and is_spanning(graph, len(data)):
                    phase = 1

            # trim to a minimal spanning tree
            if phase == 1:
                graph, pairs = minimal_span_tree(graph, pairs)

            # propose a new pair, find a loop, and find the largest pair in the loop, remove it
            if phase == 2:
                index = find_span_trim(pairs, (i, j))
                if index > insert_loc:
                    sorted_dist.pop(index)
                    sorted_dist.insert(insert_loc)
                    (remi, remj) = pairs.pop(index)
                    pairs.insert(insert_loc)
                    graph[remi].remove(remj)
                    graph[remj].remove(remi)

    last_pair = pairs[-1]
    return last_pair[0][0] * last_pair[1][0]

def is_spanning(graph, pair, expected):
    horizon = set(pair[0], pair[1])
    seen = set()
    while len(horizon) > 0:
        n = horizon.pop()
        seen.add(n)
        for edge in graph[n]:
            if not seen[edge]:
                horizon.append(edge)
    return len(seen) == expected

def minimum_spanning_tree(graph, weights, start = 0):
    seen = set()
    horizon = set(graph)
    result_pairs = []
    cheap = defaultdict(1e10)
    cheap[start] = 0
    cheapest = start
    while len(horizon) > 0:
        current = cheapest
        horizon.remove(current)
        cheap.remove(current)
        seen.add(current)

        for edge in graph[current]:
            if cheap[edge]
            cheap[edge]


def find_span_trim(graph, new_pair, pairs):
    path = _dfs(graph, set(new_pair[0]), new_pair[1], set(new_pair[0]), [new_pair[0], new_pair[1]])
    # find the last pair that appears in the path
    for i in range(len(pairs)-1, -1, -1):
        pair = pairs[i]
        for j in range(1, len(path)):
            if pair[0] == path[j] and pair[1] == path[j-1]:
                # found the pair to remove
                return i

def _dfs(graph, visited, n, pair_stack, path):
    if pair_stack[n]:
        return path
    if visited[n]:
        return None
    visited[n] = True
    pair_stack[n] = True

    for edge in graph[n]:
        result = _dfs(graph, visited, edge, pair_stack, path + [edge])
        if result is not None:
            return result

    pair_stack[n] = False
    return None

def binarySearch(list, item):
    if len(list) == 0:
        return 0
    candidate_index = len(list) // 2
    if item == list[candidate_index]:
        return candidate_index
    if item < list[candidate_index]:
        return binarySearch(list[:candidate_index], item)
    if item > list[candidate_index]:
        return candidate_index+1 + binarySearch(list[candidate_index+1:], item)


data = read("input/day8/puzzle.txt")
# data = read("input/day8/puzzle.txt")
# print(data)
# print(puzzle1(data, connections=1000))
print(puzzle2(data))
# print([ binarySearch([0, 1, 2, 3, 4, 5], i) for i in range(6) ])


