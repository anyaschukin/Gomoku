package gomoku

var identity int

type captures struct {
	capture0 uint8
	capture1 uint8
}

type node struct {
	id					int
	value				int
	goban				[19][19]position
	coordinate			coordinate
	lastMove			coordinate
	player				bool // black or white
	maximizingPlayer	bool // used by miniMax algo
	captures			captures
	parent				*node
	children			[]*node
	bestMove			*node
	depth				uint8
}

func newNode(id int, value int, newGoban *[19][19]position, coordinate coordinate, lastMove coordinate, newPlayer bool, maximizingPlayer bool, capture0, capture1 uint8, parent *node, depth uint8) *node {
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
		depth: depth,
	}
}

// Generates every move for a board, assigns value, and adds to tree
func generateBoards(current *node, lastMove coordinate, y, x int8) {
	var value int
	coordinate := coordinate{y, x}
	if hasNeigbours(y, x, &current.goban) == true && isMoveValid2(coordinate, &current.goban, current.player) == true {
		identity++
		newGoban := current.goban
		placeStone(coordinate, !current.player, &newGoban)
		if current.maximizingPlayer == true {
			value = current.value - int(float64(evaluateMove(coordinate, &newGoban, !current.player, current.captures)) / float64(current.depth))
		} else {
			value = current.value + int(float64(evaluateMove(coordinate, &newGoban, !current.player, current.captures)) / float64(current.depth))
		}
		captureTheory(coordinate, &newGoban, opponent(current.player))
		child := newNode(identity, value, &newGoban, coordinate, lastMove, !current.player, !current.maximizingPlayer, current.captures.capture1, current.captures.capture1, current, current.depth + 1)
		current.children = append(current.children, child)
	}
}

// Generates a tree of moves for a given player
// Only explores moves within a threatSpace around 2 last moves
func generateTree(current *node, lastMove, lastMove2 coordinate) {
	var y int8
	var x int8
	var threatSpace int8 = 4	// Depth 10 works with threatSpace = 1

	for y = lastMove.y - threatSpace; y <= lastMove.y+threatSpace; y++ {
		for x = lastMove.x - threatSpace; x <= lastMove.x+threatSpace; x++ {
			generateBoards(current, lastMove, y, x)
		}
	}
	for y = lastMove2.y - threatSpace; y <= lastMove2.y+threatSpace; y++ {
		for x = lastMove2.x - threatSpace; x <= lastMove2.x+threatSpace; x++ {
			if !(y >= lastMove.y-threatSpace && y <= lastMove.y+threatSpace && x >= lastMove.x-threatSpace && x <= lastMove.x+threatSpace) {
				// optimized so the threat-space searches don't overlap
				generateBoards(current, lastMove2, y, x)
			}
		}
	}
}
