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

// aiSuggestMove, if player is AI call AI to suggest a move
func aiSuggestMove(g *game) {
	if isPlayerHuman(g) == false || isPlayerHotseat(g) == true {
		artificialIdiot(g)
	}
}
