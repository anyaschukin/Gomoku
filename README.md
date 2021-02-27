# Gomoku

#### Final Score 125/100

## Challenge

Create a Gomoku game GUI, with an AI player.

#### Challenge Requirements:

* AI is Minimiax algorithm-powered
* Timer: AI must move within 0.5 seconds
* Human vs AI
* Human vs Human
* Human vs Hotseat (AI suggests move, human chooses)

## Gomoku Rules

[Gomoku](http://en.wikipedia.org/wiki/Gomoku) is a strategy board game traditionally played on a Go board with black and white stones.
Two players take turns placing their stones on an intersection of the 19x19 board.
The game ends when one player manages to align 5 or more stones.

### Additional Rules

* Capture: You can remove a pair of your opponent’s stones by flanking them.
* Capture 10: If you capture ten stones, you win.
* Game-ending capture: A player that aligns 5 stones only wins if the opponent can not break this alignment by capturing, or win by capturing 10.
* Double-threes: No two free-three alignments, which would guarantee a win by alignment.









## The Project



## Usage

## Approach

Golang for speed and elegance.

### Heuristic

### Depth

On my system an AI depth of 4 is possible in under 0.5 seconds.

A depth of 10 is possible by reducing the threat space to 1, not that this makes a better AI player.

### Optimization

#### Alpha beta pruning

#### threat space

#### has neigbours

Search space
