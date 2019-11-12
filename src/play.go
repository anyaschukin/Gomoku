package play

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

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

// SwapPlayers swaps Players, clears the message and iterates move
func SwapPlayers(G *Game) {
	G.Player = Opponent(G.Player)
	if G.Won == false {
		G.Message = ""
	}
	G.Move++

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

// gameLoop runs one move
func gameLoop(coordinate Coordinate, G *Game) {
	validated := PlaceIfValid(coordinate, G)
	if validated == true {
		Capture(coordinate, G)
		CheckWin(coordinate, G)
		G.LastMove = coordinate
		SwapPlayers(G)
	}
	SuggestMove(G)
}

// humanLoop listens for a click on the goban runs gameloop with clicked coordinate
func humanLoop(G *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		if clickGoban(x, y) == true {
			coordinate := Coordinate{-1, -1}
			coordinate.X = int8((float64(x) - (zeroX * scale)) / (positionWidth * scale))
			coordinate.Y = int8((float64(y) - (zeroY * scale)) / (positionWidth * scale))
			gameLoop(coordinate, G)
		}
	}
}

// aiLoop listens runs gameloop with suggested coordinate
func aiLoop(G *Game) {
	coordinate := G.Ai0.Suggest
	if G.Player == true {
		coordinate = G.Ai1.Suggest
	}
	gameLoop(coordinate, G)
}

func (G *Game) updateGame() { ////listen for input, update struct
	input(G)
	if G.NewGame == false && G.Won == false {
		if IsPlayerHuman(G) == true || IsPlayerHotseat(G) == true {
			humanLoop(G)
		} else {
			aiLoop(G)
		}
	}
}

// Play initializes a new game and launches the GUI (Ebiten)
func Play() {
	G := NewGame()
	G.Ai1.Depth = 3 ////////
	RunEbiten()
}
