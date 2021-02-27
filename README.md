# Gomoku

#### Final Score 125/100

## Challenge

Create a Gomoku game GUI, with an AI player.

[Gomoku](http://en.wikipedia.org/wiki/Gomoku) is a strategy board game traditionally played on a Go board with black and white stones.
Two players take turns placing their stones on an intersection of the 19x19 board.
The game ends when one player manages to align 5 or more stones.

#### Challenge Requirements:

* AI is Minimiax algorithm-powered
* Timer: AI must move within 0.5 seconds
* Human vs AI
* Human vs Human
* Human vs Hotseat (AI suggests move, human chooses)

## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder with ```git clone https://github.com/anyashuka/Gomoku.git```

Move into Gomoku folder then download dependencies with ```go get -d ./...```

Finally, to run, go run main.go directly:
```go run main.go```

Alternatively build and run the binary:
```go build; ./Gomoku```

## Example

![Example](https://github.com/anyashuka/Gomoku/blob/master/img/example.gif)

## Additional Rules

#### Capture

You can remove a pair of your opponent’s stones by flanking them. In the following scenario, by playing in a, Blue captures the red pair and removes the stones from the game. One can only capture PAIRS, not 1 or >2 stones in a row.

![capture](https://github.com/anyashuka/Gomoku/blob/master/img/capture.png?raw=true){ width=250px }

#### Capture 10

If you capture ten stones, you win.

#### Game-ending capture

A player that aligns 5 stones only wins if the opponent can not break this alignment by capturing, or win by capturing 10.

#### No Double-threes

A free-three is an alignement of three stones that, if not immediately blocked, allows for an indefendable alignment of four stones (an alignment of four stones with two unobstructed extremities). Both of these scenarios are free-threes:

![freeThree](https://github.com/anyashuka/Gomoku/blob/master/img/freeThree.png?raw=true)
![freeThree2](https://github.com/anyashuka/Gomoku/blob/master/img/freeThree2.png?raw=true)

A double-three is a move that introduces two simultaneous free-three alignments. This is an indefendable scenario.

In the following scenario, by playing in a, Red would introduce a double-three, therefore this is a forbidden move. However, if there were a blue stone in b, one of the three-aligned would be obstructed, therefore the move in a would be legal. Furthermore, it is not forbidden to introduce a double-three by capturing a pair.

![doubleFreeThree](https://github.com/anyashuka/Gomoku/blob/master/img/doubleFreeThree.png?raw=true)

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
