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

// If one Player aligns 5, their opponent must break it, capture 10, or lose.
type align5 struct {
	break5   bool // Can the opponent break the aligned 5?
	capture8 bool // Can the opponent win by capturing 10? (have they already captured 8, and is there an available capture move)
}

type ai struct {
	aiPlayer bool          // Is Player Human or Ai
	hotseat  bool          // Ai Player only suggests moves, Human must choose move
	depth    uint8         // How many moves in advance do we examine
	timer    time.Duration // How long did the Ai think for
	suggest  coordinate    // Ai suggested move
}

// captured struct records last captures for display in GUI
type captured struct {
	drawCapture       bool         // Higlight captures
	capturedPositions []coordinate // Positions of captured stones
}

// game struct contains all information about current game state
type game struct {
	goban        [19][19]position // Game board
	player       bool             // Whose move is it? (Player 0 - black first)
	ai0          ai               // Is black human or Ai?
	ai1          ai               // Is white human or Ai?
	captured     captured         // info to display last capture in GUI
	capture0     uint8            // How many stones has Black Captured? 10 to win
	capture1     uint8            // How many stones has White Captured? 10 to win
	align5       align5           // Can an aligned 5 can be broken or trumped by capturing 10?
	move         uint32           // How many moves have been played in total
	drawLastMove bool             // Higlight the last move played
	lastMove     coordinate       // What was the last move played
	newGame      bool             // New Game button has been pressed, show new game options
	won          bool             // Game finished
	drawWinMove  bool             // Higlight the last move played
	winMove      coordinate       // Winning move
	message      string           // Game feeback for display in gui (invalid move, win)
}

// Game var contains all information about current game state
var g *game
