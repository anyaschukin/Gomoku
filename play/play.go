package play

import (
	"time"
	"math/rand"
	// lib "Gomoku/golib"
	// gui "Gomoku/GUI"
)

type Coordinate struct {
	Y int8
	X int8
}

type position struct {
	occupied bool
	Player   bool
}

type align5 struct { //winning move for checking if opponent breaks it in the next move
	break5   bool
	capture8 bool // is it possible for the opponent to win by capturing 10? (have they already captured 8, and is there an avAilable capture move)
	winner   bool /// rm?
}

type Ai struct {
	AiPlayer bool          // is Player 1 human or Ai
	Hotseat  bool          // Ai Player only suggests moves, human must choose move
	Depth    uint8         // how many moves in advance do we examine
	Timer    time.Duration // How long did the Ai think for
	Suggest  Coordinate    // Ai suggested move
}

// Game struct contAins all information about currnt game state
type Game struct {
	Goban        [19][19]position // game board
	Player       bool             // whose move is it? (Player 0 - black first)
	Ai0          Ai               // is black human or Ai?
	Ai1          Ai               // is white human or Ai?
	Capture0     uint8            // capture 10 and win
	Capture1     uint8            // capture 10 and win
	align5       align5           // one Player has aligned 5, however it can be broken. The other Player must break it, capture 10, or lose.
	Move         uint32           // how many moves have been played in total (is this desirable/necessary?)
	DrawLastMove bool             // Higlight the last move played
	LastMove     Coordinate       // What was the last move played
	NewGame      bool             // New Game button has been pressed, show new game options
	Won          bool             // game finished
	Winmove      Coordinate       // how many moves have been played in total
	Message      string           // game feeback (invalid move, win)
}

// G contAins all game state info
var G *Game

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

func RandomCoordinate() Coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := Coordinate{y, x}
	return random
}

// artificialIdiot suggests a random move
func artificialIdiot(G *Game) { /////// move/remove?
	start := time.Now()
	suggestion := RandomCoordinate()
	// time.Sleep(498 * time.Millisecond) //////////
	elapsed := (time.Since(start))
	if G.Player == false {
		G.Ai0.Suggest = suggestion
		G.Ai0.Timer = elapsed
	} else {
		G.Ai1.Suggest = suggestion
		G.Ai1.Timer = elapsed
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
func Play() {
	G := NewGame()
	G.Ai1.Depth = 3////////
	RunEbiten()
}
