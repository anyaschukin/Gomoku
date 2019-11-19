package play

import (
	"math/rand"
	"time"
)

func randomCoordinate() coordinate {
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

// artificialIdiot suggests a random move
func artificialIdiot(g *game) {
	start := time.Now()
	suggestion := randomCoordinate()
	elapsed := (time.Since(start))
	if g.player == false {
		g.ai0.suggest = suggestion
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = suggestion
		g.ai1.timer = elapsed
	}
}

func isPlayerHuman(g *game) bool {
	if (g.player == false && g.ai0.aiPlayer == false) ||
		(g.player == true && g.ai1.aiPlayer == false) {
		return true
	}
	return false
}

func isPlayerHotseat(g *game) bool {
	if (g.player == false && g.ai0.hotseat == true) ||
		(g.player == true && g.ai1.hotseat == true) {
		return true
	}
	return false
}

func openingMoves(g *game) {
	start := time.Now()
	openingBlack := coordinate{9, 9}
	if g.move == 0 { // Black opening
		elapsed := (time.Since(start))
		g.ai0.suggest = openingBlack
		g.ai0.timer = elapsed
	}
	if g.move == 1 { // White opening
		openingWhite := openingBlack
		if positionOccupied(openingBlack, &g.goban) == true {
			openingWhite = coordinate{8, 8}
		}
		elapsed := (time.Since(start))
		g.ai1.suggest = openingWhite
		g.ai1.timer = elapsed
	}
}

// aiSuggestMove, if player is AI call AI to suggest a move
func aiSuggestMove(g *game) {
	if isPlayerHuman(g) == false || isPlayerHotseat(g) == true { /////do we need this??!!!
		// artificialIdiot(g)
		if g.move < 2 {
			openingMoves(g)
			return
		}
		minimaxTree(g)
	}
}
