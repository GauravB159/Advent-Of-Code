from functools import cache
import sys
import numpy as np

sys.setrecursionlimit(1000000)


text = open("./input.txt").read().split("\n")
text = [list(x) for x in text]

def get_points(limit, text, start):
    grid = {}
    for i, row in enumerate(text):
        for j, col in enumerate(row):
            grid[(i, j)] = col
    visited = {}
    queue = [(start, 0)]
    while(len(queue)):
        point, level = queue.pop(0)
        if level > limit:
            continue
        if not grid.get(point, 0):
            continue
        if point in visited:
            continue
        if grid[point] == '#':
            continue
        visited[point] = level
        queue.append(((point[0] + 1, point[1]), level + 1))
        queue.append(((point[0] - 1, point[1]), level + 1))
        queue.append(((point[0], point[1] + 1), level + 1))
        queue.append(((point[0], point[1] - 1), level + 1))

    return len([k for k in visited if visited[k] % 2 == limit % 2])


def plot_with_extended_graph(extend_factor = 1, steps = 1):
    np_text = np.array(text)
    extended_grid = np.concatenate([np_text] * extend_factor, axis=1)
    extended_grid = np.concatenate([extended_grid] * extend_factor, axis=0)
    start = (extended_grid.shape[0] // 2, extended_grid.shape[1] // 2)
    extended_grid[extended_grid.shape[0] // 2][extended_grid.shape[1] // 2] = 'S'
    result = get_points(steps, extended_grid, start)
    return result

def get_obstacle_pixel_count(text):
    odd_count = 0
    even_count = 0
    for i, row in enumerate(text):
        for j, col in enumerate(row):
            if col == '#':
                if (i + j) % 2 == 0:
                    odd_count += 1
                else:
                    even_count += 1
    return odd_count, even_count

extend_factor = 11

for i in range(5):
    steps = 65 + 131 * i
    print(plot_with_extended_graph(extend_factor, steps))