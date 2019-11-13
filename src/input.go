package play //gui

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func clickHuman0(x, y int) bool {
	if x > 108 && x < 514 &&
		y > 306 && y < 441 {
		return true
	}
	return false
}

func clickHuman1(x, y int) bool {
	if x > 108+860 && x < 514+860 &&
		y > 306 && y < 441 {
		return true
	}
	return false
}

func clickHotseat0(x, y int) bool {
	if x > 108 && x < 514 &&
		y > 530 && y < 623 {
		return true
	}
	return false
}

func clickHotseat1(x, y int) bool {
	if x > 108+860 && x < 514+860 &&
		y > 530 && y < 623 {
		return true
	}
	return false
}

func clickAI0(x, y int) bool {
	if x > 108 && x < 514 &&
		y > 715 && y < 849 {
		return true
	}
	return false
}

func clickAI1(x, y int) bool {
	if x > 108+860 && x < 514+860 &&
		y > 715 && y < 849 {
		return true
	}
	return false
}

func clickUp0(x, y int) bool {
	if x > 535 && x < 593 &&
		y > 643 && y < 705 {
		return true
	}
	return false
}

func clickUp1(x, y int) bool {
	if x > 535+860 && x < 593+860 &&
		y > 643 && y < 705 {
		return true
	}
	return false
}

func clickDown0(x, y int) bool {
	if x > 535 && x < 593 &&
		y > 848 && y < 908 {
		return true
	}
	return false
}

func clickDown1(x, y int) bool {
	if x > 535+860 && x < 593+860 &&
		y > 848 && y < 908 {
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

func clickLastMove(x, y int) bool {
	if x > 1791 && x < 2172 &&
		y > 333 && y < 415 {
		return true
	}
	return false
}

func clickWinMove(x, y int) bool {
	if x > 1791 && x < 2172 &&
		y > 533 && y < 621 {
		return true
	}
	return false
}

func inputNewGame(g *game, x, y int) {
	if clickHuman0(x, y) == true {
		g.ai0.aiPlayer = false
		g.ai0.hotseat = false
	}
	if clickAI0(x, y) == true {
		g.ai0.aiPlayer = true
		g.ai0.hotseat = false
	}
	if clickHotseat0(x, y) == true {
		if g.ai0.hotseat == false {
			g.ai0.hotseat = true
			g.ai0.aiPlayer = true
		} else {
			g.ai0.hotseat = false
			g.ai0.aiPlayer = false
		}
	}
	if clickHuman1(x, y) == true {
		g.ai1.aiPlayer = false
		g.ai1.hotseat = false
	}
	if clickAI1(x, y) == true {
		g.ai1.aiPlayer = true
		g.ai1.hotseat = false
	}
	if clickHotseat1(x, y) == true {
		if g.ai1.hotseat == false {
			g.ai1.hotseat = true
			g.ai1.aiPlayer = true
		} else {
			g.ai1.hotseat = false
			g.ai1.aiPlayer = false
		}
	}
	if clickUp0(x, y) == true {
		g.ai0.depth++
	}
	if clickDown0(x, y) == true {
		g.ai0.depth--
	}
	if clickUp1(x, y) == true {
		g.ai1.depth++
	}
	if clickDown1(x, y) == true {
		g.ai1.depth--
	}
	if clickLastMove(x, y) == true {
		if g.drawLastMove == false {
			g.drawLastMove = true
		} else {
			g.drawLastMove = false
		}
	}
	if clickWinMove(x, y) == true {
		if g.drawWinMove == false {
			g.drawWinMove = true
		} else {
			g.drawWinMove = false
		}
	}
}

func input(g *game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true {
		os.Exit(0)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		// fmt.Printf("mouse pressed x: %d, y: %d\n", x, y) ////////
		if clickExit(x, y) == true {
			os.Exit(0)
		}
		if clickNewGame(x, y) == true {
			if g.newGame == false {
				g := newGame()
				g.newGame = true
			} else {
				g.newGame = false
			}
		}
		if g.newGame == true {
			inputNewGame(g, x, y)
		}
	}
}
