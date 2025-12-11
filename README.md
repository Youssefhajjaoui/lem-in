# 🐜 Lem-in: Ant Colony Optimization

## 🌟 Project Overview

**Lem-in** is a program designed to find the **quickest and most efficient way** to move a colony of ants from a starting room (`##start`) to an end room (`##end`) through a complex network of tunnels and rooms.

The core challenge is not just finding a single shortest path, but coordinating the movement of **all $N$ ants** across multiple available paths simultaneously. This approach minimizes the total number of turns required for the *last* ant to reach the exit, effectively solving the problem of traffic jams and optimizing flow.

The project is written entirely in **Go**.

## 🚀 Getting Started

### Prerequisites

You must have the Go programming language installed on your system.

### Usage

1.  **Run the program:**
    The executable takes a single argument: the path to the text file describing the ant colony.

    ```bash
    go run . [input_file_path]
    ```

    **Example:**
    ```bash
    $ go run . colony_map.txt
    ```

## 📝 Input File Format

The program expects a standard input file structure containing three main sections in order:

1.  **Number of Ants:** A single positive integer on the first non-comment line.
2.  **Rooms:** Defined as `name coord_x coord_y`.
    * Special rooms are designated by `##start` and `##end` on the line immediately preceding the room definition.
3.  **Links (Tunnels):** Defined as `name1-name2`.
4.  **Comments:** Any line starting with `#` is ignored, except for the special commands `##start` and `##end`.

### Example Input:
3 ##start Entrance 1 0 B 5 0 ##end Exit 9 0 Entrance-B B-Exit



## 🖥️ Output Format

Upon successful execution, the program first displays the exact content of the input file you provided. Immediately following the input, the ant movement simulation begins.

### Movement Display Format

Each line after the input represents one **turn** of the simulation. It lists all the ants that moved during that turn.

The format for each movement is:

$$L_x - y$$

Where:
* $L$ is a constant prefix.
* $x$ is the **ant's number** (starting from 1 up to $N$).
* $y$ is the **room name** the ant moved into.

Multiple movements in the same turn are separated by spaces.

### Example Output:

[Input file content displayed here]
3 ##start A 1 0 ... A-B

[Ant movements start here]
L1-R2 L2-R3 L1-R4 L3-R2 L2-R5 L1-End L3-R4 L2-R6 L3-End


## ⚠️ Error Handling

The program is designed to validate all input data. If any part of the input is invalid (e.g., missing `##start` or `##end`, invalid coordinates, non-existent links, or impossible paths), the program will exit gracefully and print an error message:

ERROR: invalid data format


(The specific error message may sometimes provide more detail, such as "no start room found").

## 🛠️ Technical Details for Developers

The efficient solution to this flow problem relies on the following core algorithms and data structures:

* **Graph Parsing:** Robust parsing of the input file to construct the ant colony as a graph (rooms as nodes, tunnels as edges).
* **Pathfinding:** Utilizing a combination of **Breadth-First Search (BFS)** for performance on unweighted edges and concepts from **Dijkstra's Algorithm** to find the optimal set of paths.
* **Data Structures:** Heavy use of **Queues**, **Linked Lists**, and custom structures to manage the graph, track residual capacity, and coordinate ant movements turn-by-turn. The use of BFS over simpler DFS helps ensure the paths found are optimal for managing flow.