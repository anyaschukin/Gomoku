package play //gui

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func clickHuman0(x, y int) bool {
	if x > int(newGameBlack2*0.12) && x < 940 &&
		y > int(1550*0.12) && y < 277 {
		return true
	}
	return false
}

func clickHuman1(x, y int) bool {
	if x > int(newGameBlack2*0.12)+1092 && x < 940+1092 &&
		y > int(1550*0.12) && y < 277 {
		return true
	}
	return false
}

func clickHotseat0(x, y int) bool {
	if x > int(newGameBlack2*0.12) && x < 940 &&
		y > 286 && y < 378 {
		return true
	}
	return false
}

func clickHotseat1(x, y int) bool {
	if x > int(newGameBlack2*0.12)+1092 && x < 940+1092 &&
		y > 286 && y < 378 {
		return true
	}
	return false
}

func clickAI0(x, y int) bool {
	if x > int(950*0.12) && x < 454 &&
		y > 186 && y < 1378 {
		return true
	}
	return false
}

func clickAI1(x, y int) bool {
	if x > int(950*0.12)+1092 && x < 454+1092 &&
		y > 186 && y < 1378 {
		return true
	}
	return false
}

func clickNewGame(x, y int) bool {
	if x > int(newGameX*0.6) && x < 2492 &&
		y > int(newGameY*0.6) && y < 1240 {
		return true
	}
	return false
}

func clickExit(x, y int) bool {
	if x > int(exitX*scale) && x < 2492 &&
		y > int(exitY*scale) && y < 1372 {
		return true
	}
	return false
}

func clickGoban(x, y int) bool {
	if x > int(zeroX*scale) && x < int((zeroX*scale)+(positionWidth*float64(19)*scale)) &&
		y > int(zeroY*scale) && y < int((zeroY*scale)+(positionWidth*float64(19)*scale)) {
		return true
	}
	return false
}

func input(G *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true { // quit if press escape
		os.Exit(0) ////// is this exiting properly?
	}
	// click exit or new game
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		// fmt.Printf("mouse pressed x: %d, y: %d\n", x, y) ////////
		if clickExit(x, y) == true {
			os.Exit(0) ////// is this exiting properly?
		}
		if clickNewGame(x, y) == true {
			if G.newGame == false {
				G := NewGame()
				G.ai0.aiplayer = true
				G.newGame = true
			} else {
				G.newGame = false
			}
		}
		if G.newGame == true {
			if clickHuman0(x, y) == true {
				G.ai0.aiplayer = false
			}
			if clickAI0(x, y) == true {
				G.ai0.aiplayer = true
				G.ai0.depth = uint8((y - 186) / (1201 / 12))
			}
			if clickHotseat0(x, y) == true {
				if G.ai0.hotseat == false {
					G.ai0.hotseat = true
				} else {
					G.ai0.hotseat = false
				}
			}
			if clickHuman1(x, y) == true {
				G.ai1.aiplayer = false
			}
			if clickAI1(x, y) == true {
				G.ai1.aiplayer = true
				G.ai1.depth = uint8((y - 186) / (1201 / 12))
			}
			if clickHotseat1(x, y) == true {
				if G.ai1.hotseat == false {
					G.ai1.hotseat = true
				} else {
					G.ai1.hotseat = false
				}
			}
		}
	}
}
