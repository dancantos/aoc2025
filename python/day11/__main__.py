import numpy as np

def read(filename):
    d = {}
    with open(filename) as f:
        for line in f:
            k, v = process_line(line)
            d[k] = v
    return d

def process_line(line):
    parts = line.strip().split()
    return (parts[0][:-1], [ part for part in parts[1:] ])

def solve1(graph):
    return _dfs(graph, 'you', 'out', set({'you'}))

def solve2(graph):
    return _matsolve(graph, 'svr', 'out')

def _dfs(graph, node, target, path):
    count = 0
    for edge in graph[node]:
        if edge in path:
            continue
        if edge == target:
            return 1
        path.add(edge)
        count += _dfs(graph, edge, target, path)
        path.remove(edge)
    return count

def _matsolve(graph, node, target):
    l = len(graph)+1
    mat = [ [ 0 for _ in range(l) ] for _ in range(l) ]
    svr_index = 0
    dac_index = 0
    fft_index = 0
    out_index = l-1
    i = 0
    for k in graph:
        # record indexes that represent our key interests
        if k == 'svr':
            svr_index = i
        if k == 'dac':
            dac_index = i
        if k == 'fft':
            fft_index = i

        # populate the adjacency matrix
        j = 0
        for v in graph:
            if v in graph[k]:
                mat[i][j] = 1
            if 'out' in graph[k]:
                mat[i][l-1] = 1
            j += 1
        i += 1

    svr_dac_count = 0
    svr_fft_count = 0
    dac_fft_count = 0
    fft_dac_count = 0
    dac_out_count = 0
    fft_out_count = 0
    base = np.array(mat)
    m = base.copy()
    for i in range(len(graph)):
        svr_dac_count += m[svr_index, dac_index]
        svr_fft_count += m[svr_index, fft_index]
        dac_fft_count += m[dac_index, fft_index]
        fft_dac_count += m[fft_index, dac_index]
        dac_out_count += m[dac_index, out_index]
        fft_out_count += m[fft_index, out_index]
        m = np.matmul(m, base)
    return (svr_dac_count * dac_fft_count * fft_out_count) + (svr_fft_count * fft_dac_count * dac_out_count)

# graph = read("input/day11/sample.txt")
# graph = read("input/day11/sample2.txt")
graph = read("input/day11/puzzle.txt")
print(solve1(graph))
print(solve2(graph))
