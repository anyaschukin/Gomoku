package play

func isPlayerHuman(g *game) bool {
	if (g.player == false && g.ai0.aiPlayer == false) ||
		(g.player == true && g.ai1.aiPlayer == false) {
		return true
	}
	return false
}

func isOpponentHuman(g *game) bool {
	if (g.player == false && g.ai1.aiPlayer == false) ||
		(g.player == true && g.ai0.aiPlayer == false) {
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
