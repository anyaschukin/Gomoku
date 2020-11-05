package gomoku

import (
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func clickGoban(x, y int) bool {
	if x > int(gobanX*scale) && x < int((gobanX*scale)+(positionWidth*float64(19)*scale)) &&
		y > int(gobanY*scale) && y < int((gobanY*scale)+(positionWidth*float64(19)*scale)) {
		return true
	}
	return false
}

func clickNewGame(x, y int) bool {
	if x > int(newGameX*newGameScale) && x < 2492 &&
		y > int(newGameY*newGameScale) && y < 1240 {
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

func clickUndoButton(x, y int) bool {
	if x > 83 && x < 241 &&
		y > 1085 && y < 1232 {
		return true
	}
	return false
}

func inColumnX(x, column int) bool {
	if x > newGameColumnBlack-25+column && x < newGameColumnBlack+340+column {
		return true
	}
	return false
}

func inRowY(y, rowI int) bool {
	if y > rowI*row+25 && y < rowI*row+134 {
		return true
	}
	return false
}

func clickHuman(x, y, column int) bool {
	if inColumnX(x, column) == true &&
		inRowY(y, 3) {
		return true
	}
	return false
}

func clickHotseat(x, y, column int) bool {
	if inColumnX(x, column) == true &&
		inRowY(y, 5) {
		return true
	}
	return false
}

func clickAI(x, y, column int) bool {
	if inColumnX(x, column) == true &&
		inRowY(y, 7) {
		return true
	}
	return false
}

func clickUp(x, y, column int) bool {
	if x > 495+column && x < 550+column &&
		y > 664 && y < 723 {
		return true
	}
	return false
}

func clickDown(x, y, column int) bool {
	if x > 495+column && x < 550+column &&
		y > 837 && y < 898 {
		return true
	}
	return false
}

func clickLastMove(x, y int) bool {
	if inColumnX(x, column1*2) == true &&
		inRowY(y, 3) {
		return true
	}
	return false
}

func clickWinMove(x, y int) bool {
	if inColumnX(x, column1*2) == true &&
		inRowY(y, 5) {
		return true
	}
	return false
}

func clickCapture(x, y int) bool {
	if inColumnX(x, column1*2) == true &&
		inRowY(y, 7) {
		return true
	}
	return false
}

func clickUndo(x, y int) bool {
	if inColumnX(x, column1*3) == true &&
		inRowY(y, 3) {
		return true
	}
	return false
}

func clickTips(x, y int) bool {
	if inColumnX(x, column1*3) == true &&
		inRowY(y, 5) {
		return true
	}
	return false
}

func clickDoge(x, y int) bool {
	if inColumnX(x, column1*3) == true &&
		inRowY(y, 7) {
		return true
	}
	return false
}

func clickColor(x, y int) bool {
	if inColumnX(x, column1*2) == true &&
		y > 1000 && y < 1357 {
		return true
	}
	return false
}

func clickPlayer(x, y int, player bool) {
	column := 0
	p := &g.ai0
	if player == true {
		column = column1
		p = &g.ai1
	}
	if clickHuman(x, y, column) == true {
		p.aiPlayer = false
		p.hotseat = false
	}
	if clickAI(x, y, column) == true {
		p.aiPlayer = true
		p.hotseat = false
	}
	if clickHotseat(x, y, column) == true {
		if p.hotseat == false {
			p.hotseat = true
			p.aiPlayer = true
		} else {
			p.hotseat = false
			p.aiPlayer = false
		}
	}
	if clickUp(x, y, column) == true {
		p.depth++
	}
	if clickDown(x, y, column) == true {
		if p.depth > 0 {
			p.depth--
		}
	}
}

func swapBool(boolean *bool) {
	*boolean = !*boolean
}

func swapFullscreen() {
	if ebiten.IsFullscreen() == true {
		ebiten.SetFullscreen(false)
	} else {
		ebiten.SetFullscreen(true)
	}
}

func swapDogeMode() {
	if dogeMode == true {
		imgBlack = imgBlackStone
		imgWhite = imgWhiteStone
		dogeMode = false
	} else {
		imgBlack = imgDoge
		imgWhite = imgCorg
		dogeMode = true
	}
}

func inputNewGame(g *game, x, y int) {
	clickPlayer(x, y, false)
	clickPlayer(x, y, true)
	if clickLastMove(x, y) == true {
		swapBool(&g.gui.drawLastMove)
	}
	if clickWinMove(x, y) == true {
		swapBool(&g.gui.drawWinMove)
	}
	if clickCapture(x, y) == true {
		swapBool(&g.gui.drawCapture)
	}
	if clickUndo(x, y) == true {
		swapBool(&g.gui.undo)
	}
	if clickTips(x, y) == true {
		swapBool(&g.gui.tips)
	}
	if clickDoge(x, y) == true {
		swapDogeMode()
	}
}

// input listens for keyboard and mouse input
func input(g *game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true {
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) == true {
		swapFullscreen()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) == true {
		swapDogeMode()
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		// fmt.Printf("mouse pressed x: %d, y: %d\n", x, y) // debug tool
		if clickExit(x, y) == true {
			os.Exit(0)
		}
		if clickNewGame(x, y) == true {
			if g.gui.newGame == false {
				g = newGame()
			}
			swapBool(&g.gui.newGame)
		}
		if clickUndoButton(x, y) == true {
			undo(g)
		}
		if g.gui.newGame == true {
			inputNewGame(g, x, y)
		}
	}
}
