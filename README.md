# nPuzzle

## Web-app

![front](https://github.com/42Projects/nPuzzle/blob/master/misc/front.png)

### Installation
Requirements:
 - docker
 - docker-compose

`$ make deploy`

Wait for it...

### Usage
...application is now available on localhost:8080

## CLI

### Installation
Requirements:
 - go (if you need to compile from source)

You can either use the ELF-64 binary ./bin/nPuzzleCli if your arch allows it, or compile a new binary from source with:

`$ make cli`

### Usage

`$ ./bin/nPuzzleCli [options] [file]`

Options:
 - display: display a step-by-step solution (default: `true`)
 - heuristic: choose between available heuristics (default: `manhattan + linear conflicts`)
   - hamming distance: a misplaced tile adds one to the score (very weak)
   - manhattan distance: a move from each tile to it's goal position adds one to the score (average)
   - manhattan distance + linear conflicts: a linear conflict (two tiles in their goal row or column but inverted) adds two to the score (best)
 - search: modify the A* algorithm search variant (default: `uniform-cost`)
   - greedy: at each step, the move that is closest to the goal state will be picked. Very fast algorithm but the solution will not be optimal (use for 5x5+ size grids)
   - uniform-cost search: the lesser overall cost from the beginning will be favored. Will return the best solution but is slow and extremely memory-consuming on 5x5 grids and above
 - timeout: modify the base timeout (default: `1m0s`)

If no file is provided, binary will await entry from standard input.

You can also use the python puzzle generator with:

`$ ./misc/res_npuzzle-gen.py [-s | -u] [size] | ./bin/nPuzzleCli [options]`
