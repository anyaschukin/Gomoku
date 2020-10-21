package gomoku

import (
	"fmt"
	"time"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

var identity int

type captures struct {
	capture0 uint8
	capture1 uint8
}

type node struct {
	id               int
	value            int
	goban            [19][19]position
	coordinate       coordinate
	lastMove         coordinate
	player           bool // black or white
	maximizingPlayer bool // used by miniMax algo
	captures         captures
	parent           *node
	children         []*node
	bestMove         *node
}

func newNode(id int, value int, newGoban *[19][19]position, coordinate coordinate, lastMove coordinate, newPlayer bool, maximizingPlayer bool, capture0, capture1 uint8, parent *node) *node {
	return &node{
		id:               id,
		value:            value,
		goban:            *newGoban,
		coordinate:       coordinate,
		lastMove:         lastMove,
		player:           newPlayer,
		maximizingPlayer: maximizingPlayer,
		captures: captures{
			capture0: capture0,
			capture1: capture1,
		},
		parent: parent,
	}
}

// Recursively finds node by ID, and then appends child to node.chilren
func addChild(node *node, parentID int, child *node) {
	if node.id == parentID {
		node.children = append(node.children, child)
	} else {
		for idx, _ := range node.children {
			current := node.children[idx]
			addChild(current, parentID, child)
		}
	}
}

// Generates every move for a board, assigns value, and adds to tree
func generateBoards(current *node, lastMove coordinate, x, y int8) {
	var value int
	coordinate := coordinate{y, x}
	if isMoveValid2(coordinate, &current.goban, current.player) == true { // duplicate of isMoveValid w/o *game
		identity++
		newGoban := current.goban
		placeStone(coordinate, !current.player, &newGoban)
		if current.maximizingPlayer == true {
			value = current.value - int(float64(evaluateMove(coordinate, &newGoban, !current.player, current.captures)) * 0.9)
		} else {
			value = current.value + evaluateMove(coordinate, &newGoban, !current.player, current.captures)
		}
		captureTheory(coordinate, &newGoban, opponent(current.player))
		child := newNode(identity, value, &newGoban, coordinate, lastMove, !current.player, !current.maximizingPlayer, current.captures.capture1, current.captures.capture1, current)
		// fmt.Printf("current.coordinate = %v, child.coordinate = %v, child.parent.coordinate = %v\n", current.coordinate, child.coordinate, child.parent.coordinate)
		addChild(current, current.id, child)
	}
}

// Returns true if given position has immediate neighbor which is occupied
func hasNeigbours(y_orig int8, x_orig int8, goban *[19][19]position) bool {
	possibleMove := coordinate{y_orig, x_orig}
	if coordinateOnGoban(possibleMove) == false {
		return false
	}

	var x int8
	var y int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				neighbour := findNeighbour(possibleMove, y, x, 1)
				if coordinateOnGoban(neighbour) == true {
					if positionOccupied(neighbour, goban) == true {
						return true
					}
				}
			}
		}
	}
	return false
}

func generateChildBoards(current *node, lastMove, lastMove2 coordinate) {
	var y int8
	var x int8
	var threatSpace int8 = 4	// Depth 10 works with threatSpace = 1

	// threat-space search of 4
	for y = lastMove.y - threatSpace; y <= lastMove.y+threatSpace; y++ {
		for x = lastMove.x - threatSpace; x <= lastMove.x+threatSpace; x++ {
			// Optimized so that only populated parts of the board are explored. Standalone/isolated positions are ignored.
			if hasNeigbours(y, x, &current.goban) == true {
				generateBoards(current, lastMove, x, y)
			}
		}
	}
	for y = lastMove2.y - threatSpace; y <= lastMove2.y+threatSpace; y++ {
		for x = lastMove2.x - threatSpace; x <= lastMove2.x+threatSpace; x++ {
			// optimized so the threat-space searches don't overlap
			if !(y >= lastMove.y-threatSpace && y <= lastMove.y+threatSpace && x >= lastMove.x-threatSpace && x <= lastMove.x+threatSpace) {
				if hasNeigbours(y, x, &current.goban) == true {
					generateBoards(current, lastMove2, x, y)
				}
			}
		}
	}
}

func minimaxTree(g *game) {
	start := time.Now()
	limit := g.ai0.depth
	if g.player == true {
		limit = g.ai1.depth
	}

	root := newNode(0, 0, &g.goban, g.lastMove, g.lastMove2, !g.player, false, g.capture0, g.capture1, nil)
	identity = 0
	alpha := minInt
	beta := maxInt
	// minimaxRecursive(root, limit, alpha, beta, true)//////////////!!!!!!!! for test
	value_wtf := minimaxRecursive(root, limit, alpha, beta, true)//////////////!!!!!!!! for test
	fmt.Printf("value_wtf: %v, player = %v, root.bestMove.value = %d\n", value_wtf, root.player, root.bestMove.value) ///////////!!!!!!!!
	
	elapsed := (time.Since(start))
	besty := root.bestMove

	if g.player == false {
		g.ai0.suggest = besty.coordinate
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = besty.coordinate
		g.ai1.timer = elapsed
	}
}

