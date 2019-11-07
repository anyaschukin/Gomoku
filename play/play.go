package play

import "time"

// import (
// 	"fmt"
// 	// lib "Gomoku/golib"
// 	// gui "Gomoku/GUI"
// )

type coordinate struct {
	y int8
	x int8
}

type position struct {
	occupied bool
	player   bool
}

type align5 struct { //winning move for checking if opponent breaks it in the next move
	break5   bool
	capture8 bool // is it possible for the opponent to win by capturing 10? (have they already captured 8, and is there an available capture move)
	winner   bool /// rm?
}

type ai struct {
	aiplayer bool          // is player 1 human or AI
	hotseat  bool          // AI player only suggests moves, human must choose move
	depth    uint8         // how many moves in advance do we examine
	timer    time.Duration // How long did the ai think for
	suggest  coordinate    // ai suggested move
}

// Game struct contains all information about currnt game state
type Game struct {
	goban        [19][19]position // game board
	player       bool             // whose move is it? (player 0 - black first)
	ai0          ai               // is black human or ai?
	ai1          ai               // is white human or ai?
	capture0     uint8            // capture 10 and win
	capture1     uint8            // capture 10 and win
	align5       align5           // one player has aligned 5, however it can be broken. The other player must break it, capture 10, or lose.
	move         uint32           // how many moves have been played in total (is this desirable/necessary?)
	drawLastMove bool             // Higlight the last move played
	lastMove     coordinate       // What was the last move played
	newGame      bool             // New Game button has been pressed, show new game options
	won          bool             // game finished
	winmove      coordinate       // how many moves have been played in total
	message      string           // game feeback (invalid move, win)
}

// G contains all game state info
var G *Game

// NewGame initializes a new game
func NewGame() *Game {
	G = &Game{}
	return G
}

// Opponent returns the opponent of the current player
func Opponent(player bool) bool {
	if player == false {
		return true
	}
	return false
}

// SwapPlayers swaps players and clears the message
func SwapPlayers(G *Game) {
	G.player = Opponent(G.player)
	if G.won == false {
		G.message = ""
	}
}

// artificialIdiot suggests a random move
func artificialIdiot(G *Game) { /////// move/remove?
	start := time.Now()
	suggestion := RandomCoordinate()
	// time.Sleep(498 * time.Millisecond) //////////
	elapsed := (time.Since(start))
	if G.player == false {
		G.ai0.suggest = suggestion
		G.ai0.timer = elapsed
	} else {
		G.ai1.suggest = suggestion
		G.ai1.timer = elapsed
	}
}

// suggestMove checks if player is AI & if so prompts the AI to suggest a move
func suggestMove(G *Game) {
	if isPlayerHuman(G) == false || isPlayerHotseat(G) == true {
		artificialIdiot(G) /////create move suggestion
	}
}

// Play initializes a new game and launches the GUI (Ebiten)
func Play() {
	G := NewGame()
	G.ai0.aiplayer = true
	G.ai0.depth = 3
	G.drawLastMove = true /////////// implement in gui!!!!!!!
	suggestMove(G)
	RunEbiten()
}
