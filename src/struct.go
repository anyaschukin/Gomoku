package play

import "time"

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
	aiPlayer bool          // is Player 1 human or Ai
	hotseat  bool          // Ai Player only suggests moves, human must choose move
	depth    uint8         // how many moves in advance do we examine
	timer    time.Duration // How long did the Ai think for
	suggest  coordinate    // Ai suggested move
}

// Game struct contAins all information about currnt game state
type game struct {
	Goban        [19][19]position // game board
	player       bool             // whose move is it? (Player 0 - black first)
	ai0          ai               // is black human or Ai?
	ai1          ai               // is white human or Ai?
	capture0     uint8            // capture 10 and win
	capture1     uint8            // capture 10 and win
	align5       align5           // one Player has aligned 5, however it can be broken. The other Player must break it, capture 10, or lose.
	Move         uint32           // how many moves have been played in total (is this desirable/necessary?)
	DrawLastMove bool             // Higlight the last move played
	LastMove     coordinate       // What was the last move played
	newGame      bool             // New Game button has been pressed, show new game options
	Won          bool             // game finished
	Winmove      coordinate       // how many moves have been played in total
	Message      string           // game feeback (invalid move, win)
}

// G contAins all game state info
var g *game
