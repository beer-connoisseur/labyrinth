# Maze Generator and Solver

A command-line application written in Go that generates mazes and finds paths within them. It supports **three maze generation algorithms** (DFS, Prim, Kruskal) and **three pathfinding algorithms** (A*, Dijkstra, Bellman–Ford). Results can be displayed in the console or saved to a file.

## Features

- Generate mazes of arbitrary size using:
    - Depth‑First Search (DFS) algorithm
    - Prim‘s algorithm
    - Kruskal‘s algorithm
- Find the shortest path from a given start point (**O**) to an end point (**X**) using:
    - A* (A-star) algorithm
    - Dijkstra‘s algorithm
    - Bellman–Ford algorithm
- Render the maze and the found path in the console with symbolic graphics:
    - `#` – wall
    - space – passable cell
    - `O` – start
    - `X` – finish
    - `.` – path cells
- Save generated mazes and solved paths to text files (optional).
- Validate input data and handle errors (invalid coordinates, negative dimensions, missing files, etc.).

