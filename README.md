# Gomoku

#### Final Score 125/100

## Challenge

Create a Gomoku game GUI, with an AI player.

[Gomoku](http://en.wikipedia.org/wiki/Gomoku) is a strategy board game traditionally played on a Go board with black and white stones.
Two players take turns placing their stones on an intersection of the 19x19 board.
A player wins by aligning 5 or more stones.

#### Requirements:

* AI uses Minimiax algorithm
* Timer: AI must move within 0.5 seconds
* Human vs AI
* Human vs Human
* Human vs Hotseat (AI suggests move, human chooses)

## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder.

```git clone https://github.com/anyashuka/Gomoku.git; cd Gomoku```

Download dependencies.

```go get -d ./...```

Finally, to run.

```go run main.go```

Alternatively build and run the binary.

```go build; ./Gomoku```

![Example](https://github.com/anyashuka/Gomoku/blob/master/img/example.gif)

## Additional Rules

#### Capture

You can remove a pair of your opponent’s stones by flanking them. In the following scenario, by playing in a, Blue captures the red pair and removes the stones from the game. One can only capture PAIRS, not 1 or >2 stones in a row.

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/capture.png" width="20%">

#### Capture 10

If you capture ten stones, you win.

#### Game-ending capture

A player that aligns 5 stones only wins if the opponent can not break this alignment by capturing, or win by capturing 10.

#### No Double-threes

A free-three is an alignement of three stones that, if not immediately blocked, allows for an indefendable alignment of four stones (an alignment of four stones with two unobstructed extremities). Both of these scenarios are free-threes:

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/freeThree.png" width="20%">

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/freeThree2.png" width="25%">

A double-three is a move that introduces two simultaneous free-three alignments. This is an indefendable scenario.

In the following scenario, by playing in a, Red would introduce a double-three, therefore this is a forbidden move. However, if there were a blue stone in b, one of the three-aligned would be obstructed, therefore the move in a would be legal. Furthermore, it is not forbidden to introduce a double-three by capturing a pair.

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/doubleFreeThree.png" width="37%">

## Approach

Written in Golang for speed and elegance.

### Heuristic

For each move considered by the AI ```evaluateMove()``` checks each vertex for alignments, blocks, and captures.
```aiPriority.go``` contains values for each type of alignment, block and capture. Alignments are checked for freedom to expand on both sides, vs flanked on one side, double flanked are ignored.

For any considered move the sum of values is calculated.

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/aiPriority.png" width="37%">

#### Depth > 1:
+ If considering the players move, add the value of that move.
- If considering the opponents move, minus the value of that move.

Divide the value added/subtracted by the depth. This solves the problem of needing to be defensive (valuing blocking an opponents free-three more than placing ones own free-three, failure to do so will lose the game), while simultanously overcoming this pessimistic viewpoint (what's the point of me aligning if they are probably then going to block, and blocking is worth more than aligning).

### Optimization

The goban is represented efficiently by a (19 x 19) 2D array of positions, each position made of 2 bools (2 bits) occupied and player.

The high branching factor of this problem makes it difficult to reach a deep enough depth for an intelligent AI, while also returning suggested moves within a short enough time. There are several ways in which we can reduce branching:

#### Alpha-beta pruning

While searching the minimax tree for the best move, we can cut off branches which need not be searched because there already exists a better move available.

We keep track of two values: Alpha and Beta.
* Alpha = the minimum score that the player is assured of.
* Beta = the maximum score that the opponent is assured of.

Initially, alpha is negative infinity and beta is positive infinity, i.e. both players start with their worst possible score.
Whenever the maximum score that the opponent is assured of becomes less than the minimum score that the player is assured of (i.e. beta < alpha), the player need not consider further descendants of this node.
Search can be limited to 'more promising' subtrees, so a deeper search can be performed in the same time.

For example: Move "A" will improve the player's position. The player continues to look for a better move. Move "B" is also a good move, but the player realizes that it will allow the opponent to win in 2 moves. Thus, other outcomes from playing move "B" no longer need to be considered since the opponent can force a win. The maximum score that the opponent could force after move "B" is negative infinity: a loss for the player. This is less than the minimum position that was previously found; move "A" does not result in a loss in 2 moves.

#### Threat space

A threat space of of 4 spaces around the last two moves reduces the search space. In the following example, everything within the two red squares is within the threat space, and so is considered for the next move.

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/threatSpace.png" width="640">

On my system, with the default threatspace of 4, an AI depth of 4 is possible in under 0.5 seconds. A depth of 10 is possible by reducing the threat space to 1, not that this makes a better AI player.

#### hasNeigbours()

Branching can be further reduced by excluding all moves which do not have have immediate neighbours, i.e. unconnected to anything. In the following example, everything within the two red squares has an immediate neighbour, and so is considered for the next move.

<img src="https://github.com/anyashuka/Gomoku/blob/master/img/hasNeighbours.png" width="640">

## Options

### Hotseat

The AI suggests a move with a pulsing stone, but a human player must choose (click).

![Hotseat](https://github.com/anyashuka/Gomoku/blob/master/img/hotseat.gif)

### Doge-mode

Press key ```d```, or click the hidden doge-mode button in the new game screen, to toggle doge-mode.

![Doge](https://github.com/anyashuka/Gomoku/blob/master/img/doge.gif)

### Fullscreen

Press key ```f``` to toggle fullscreen.

## Dependencies

Thankfully, running ```go get -d ./...``` should take care of all dependencies for you.

The GUI uses [Ebiten](https://github.com/hajimehoshi/ebiten), a dead simple open source 2D game library for Go.
