package gomoku

import (
	"time"
)

// coordinate pinpoints the location of a goban position
type coordinate struct {
	y	int8
	x	int8
}

// position describes a goban positions state. The goban is 19 × 19 positions
type position struct {
	occupied	bool
	player		bool
}

// ai struct contains info for one player
type ai struct {
	aiPlayer	bool					// Is Player Human or AI
	hotseat		bool					// AI only suggests moves, Human must choose move
	depth		uint8					// How many moves in advance do we examine
	timer		time.Duration			// How long did the AI think for
	suggest		coordinate				// AI suggested move
}

// if a Player aligns 5, their opponent must break it, capture 10, or lose
type align5 struct {
	break5		bool					// Can the opponent break the aligned 5?
	capture8	bool					// Can the opponent win by capturing 10?
}

// gui struct contains info for GUI
type gui struct {
	newGame				bool			// Show new game options
	drawIntro			bool			// Finished drawing the intro
	introTime			time.Time		// When was the game started
	message				string			// Game feeback for display in gui (invalid move, win)
	drawLastMove		bool			// Higlight the last move played
	drawWinMove			bool			// Higlight the last move played
	drawCapture			bool			// Higlight captures
	capturedPositions	[]coordinate	// Positions of captured stones
	capturedPositions2	[]coordinate	// Positions of captured stones from move before last
	undo				bool			// Show undo button
	undoMove			uint32			// Last move undone
	tips				bool			// Display undo button
	align5Positions		[]coordinate	// Positions of align5
	canCapturePositions	[]coordinate	// Positions can capture
	canCaptureByPlaying	[]coordinate	// Positions can capture by playing
}

// game struct contains all information about current game state
type game struct {
	goban		[19][19]position		// Game board
	player		bool					// Whose move is it? (Player 0 - black first)
	ai0			ai						// Is black human or AI?
	ai1			ai						// Is white human or AI?
	capture0	uint8					// How many stones has Black Captured? 10 to win
	capture1	uint8					// How many stones has White Captured? 10 to win
	align5		align5					// Can aligned 5 be broken, or trumped by capturing 10?
	move		uint32					// How many moves have been played in total
	lastMove	coordinate				// Last move played
	lastMove2	coordinate				// Move before last played
	won			bool					// Game finished
	winMove		coordinate				// Winning move
	gui			gui 					// Info to display in GUI
}

// game var contains all information about current game state
var g *game
