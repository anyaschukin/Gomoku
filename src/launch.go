package play

// NewGame initializes a new game
func NewGame() *Game {
	G = &Game{}
	G.Ai0.AiPlayer = true
	G.Ai0.Depth = 3
	G.DrawLastMove = true /////////// implement in gui!!!!!!!
	SuggestMove(G)
	return G
}

// Opponent returns the opponent of the current Player
func Opponent(Player bool) bool {
	if Player == false {
		return true
	}
	return false
}

// SwapPlayers swaps Players and clears the message
func SwapPlayers(G *Game) {
	G.Player = Opponent(G.Player)
	if G.Won == false {
		G.Message = ""
	}
}

func IsPlayerHuman(G *Game) bool {
	if (G.Player == false && G.Ai0.AiPlayer == false) ||
		(G.Player == true && G.Ai1.AiPlayer == false) {
		return true
	}
	return false
}

func IsPlayerHotseat(G *Game) bool {
	if (G.Player == false && G.Ai0.Hotseat == true) ||
		(G.Player == true && G.Ai1.Hotseat == true) {
		return true
	}
	return false
}

// suggestMove checks if Player is Ai & if so prompts the Ai to suggest a move
func SuggestMove(G *Game) {
	if IsPlayerHuman(G) == false || IsPlayerHotseat(G) == true {
		artificialIdiot(G) /////create move suggestion
	}
}

// Play initializes a new game and launches the GUI (Ebiten)
func Launch() {
	G := NewGame()
	G.Ai1.Depth = 3 ////////
	RunEbiten()
}
