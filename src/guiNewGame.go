package play //gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var newGameColumnBlack = 140
var column1 = 640
var scaleSelect = 0.21

func drawStone(screen *ebiten.Image, g *game, column int, stone *ebiten.Image) {
	drawImage(screen, stone, float64(newGameColumnBlack+column)+130, 150, 1)
}

func drawSelectHuman(screen *ebiten.Image, g *game, shift float64) {
	drawImage(screen, imgSelect, 90/scaleSelect+shift, float64(row*3+3)/scaleSelect, scaleSelect)
}

func drawSelectHotseat(screen *ebiten.Image, g *game, shift float64) {
	drawImage(screen, imgSelect, 90/scaleSelect+shift, float64(row*5+3)/scaleSelect, scaleSelect)
}

func drawSelectAI(screen *ebiten.Image, g *game, shift float64) {
	drawImage(screen, imgSelect, 90/scaleSelect+shift, float64(row*7+3)/scaleSelect, scaleSelect)
}

func drawSelectPlayer(screen *ebiten.Image, g *game, player bool) {
	p := g.ai0
	var shift float64
	if player == true {
		p = g.ai1
		shift = float64(column1) / scaleSelect
	}
	if p.hotseat == true {
		drawSelectHuman(screen, g, shift)
		drawSelectHotseat(screen, g, shift)
		drawSelectAI(screen, g, shift)
	} else if p.aiPlayer == false {
		drawSelectHuman(screen, g, shift)
	} else {
		drawSelectAI(screen, g, shift)
	}
}

func drawSelectLastMove(screen *ebiten.Image, g *game) {
	if g.gui.drawLastMove == true {
		drawImage(screen, imgSelect, (float64(newGameColumnBlack+column1*2)-50)/scaleSelect, float64(row*3+3)/scaleSelect, scaleSelect)
	}
}

func drawSelectWinMove(screen *ebiten.Image, g *game) {
	if g.gui.drawWinMove == true {
		drawImage(screen, imgSelect, (float64(newGameColumnBlack+column1*2)-50)/scaleSelect, float64(row*5+3)/scaleSelect, scaleSelect)
	}
}

func drawSelectCapture(screen *ebiten.Image, g *game) {
	if g.gui.drawCapture == true {
		drawImage(screen, imgSelect, (float64(newGameColumnBlack+column1*2)-50)/scaleSelect, float64(row*7+3)/scaleSelect, scaleSelect)
	}
}

func drawSelectUndo(screen *ebiten.Image, g *game) {
	if g.gui.undo == true {
		drawImage(screen, imgSelect, (float64(newGameColumnBlack+column1*3)-50)/scaleSelect, float64(row*3+3)/scaleSelect, scaleSelect)
	}
}

func drawSelectTips(screen *ebiten.Image, g *game) {
	if g.gui.tips == true {
		drawImage(screen, imgSelect, (float64(newGameColumnBlack+column1*3)-50)/scaleSelect, float64(row*5+3)/scaleSelect, scaleSelect)
	}
}

func drawSelect(screen *ebiten.Image, g *game) {
	drawSelectPlayer(screen, g, false)
	drawSelectPlayer(screen, g, true)
	drawSelectLastMove(screen, g)
	drawSelectWinMove(screen, g)
	drawSelectCapture(screen, g)
	drawSelectUndo(screen, g)
	drawSelectTips(screen, g)
}

func drawHuman(screen *ebiten.Image, g *game) {
	text.Draw(screen, `Human`, mplusMediumFont, newGameColumnBlack, row*4, color.Black)
	text.Draw(screen, `Human`, mplusMediumFont, newGameColumnBlack+column1, row*4, color.White)
}

func drawHotseatPulse(screen *ebiten.Image, g *game, alpha float64) {
	if g.ai0.hotseat == true {
		drawImagePulse(screen, imgBlack, float64(newGameColumnBlack+370), 525, 1, alpha)
	}
	if g.ai1.hotseat == true {
		drawImagePulse(screen, imgWhite, float64(newGameColumnBlack+column1+370), 525, 1, alpha)
	}
}

func drawHotseat(screen *ebiten.Image, g *game, alpha float64) {
	text.Draw(screen, `Hotseat`, mplusMediumFont, newGameColumnBlack, row*6, color.Black)
	text.Draw(screen, `Hotseat`, mplusMediumFont, newGameColumnBlack+column1, row*6, color.White)
	drawHotseatPulse(screen, g, alpha)
}

func drawAIplayer(screen *ebiten.Image, g *game, depth uint8, column int, color color.Color) {
	text.Draw(screen, `AI - depth`, mplusMediumFont, newGameColumnBlack+column, row*8, color)
	text.Draw(screen, strconv.Itoa(int(depth)), mplusMediumFont, newGameColumnBlack+column+360, row*8, color)
	text.Draw(screen, string('⇧'), mplusBigFont, newGameColumnBlack+column+345, row*7+20, color)
	text.Draw(screen, string('⇩'), mplusBigFont, newGameColumnBlack+column+345, row*9-10, color)
}

func drawAI(screen *ebiten.Image, g *game) {
	drawAIplayer(screen, g, g.ai0.depth, 0, color.Black)
	drawAIplayer(screen, g, g.ai1.depth, column1, color.White)
}

func drawLastMovePulse(screen *ebiten.Image, g *game, alpha float64, blue *ebiten.Image) {
	drawImagePulse(screen, blue, float64(newGameColumnBlack+(column1*2)+370), 325, 1, alpha*2)
}

func drawLastMovePulses(screen *ebiten.Image, g *game, second, pulse, alpha float64) {
	if g.gui.drawLastMove == true {
		drawLastMovePulse(screen, g, alpha4(second, pulse), imgBlue)
		drawLastMovePulse(screen, g, alpha3(second, alpha), imgBlue2)
		drawLastMovePulse(screen, g, alpha2(second, pulse), imgBlue3)
		drawLastMovePulse(screen, g, alpha1(second, alpha), imgBlue4)
	}
}

func drawWinMovePulse(screen *ebiten.Image, g *game, alpha float64) {
	if g.gui.drawWinMove == true {
		drawImagePulse(screen, imgRed, float64(newGameColumnBlack+(column1*2)+370), 525, 1, alpha)
	}
}

func drawCapturedPulse(screen *ebiten.Image, g *game, alpha float64) {
	if g.gui.drawCapture == true {
		drawImagePulse(screen, imgCapture, float64(newGameColumnBlack+(column1*2)+370), 725, 1, alpha)
	}
}

func drawHighlight(screen *ebiten.Image, g *game, second, pulse, alpha float64) {
	ebitenutil.DrawRect(screen, float64(newGameColumnBlack+column1*2), 242, 320, 6, color.Black)
	text.Draw(screen, `Highlight`, mplusBigFont, newGameColumnBlack+column1*2, row*2+22, color.Black)
	text.Draw(screen, `Last Move`, mplusMediumFont, newGameColumnBlack+column1*2, row*4, color.Black)
	drawLastMovePulses(screen, g, second, pulse, alpha)
	text.Draw(screen, `Win Move`, mplusMediumFont, newGameColumnBlack+column1*2, row*6, color.Black)
	drawWinMovePulse(screen, g, alpha)
	text.Draw(screen, `Captured`, mplusMediumFont, newGameColumnBlack+column1*2, row*8, color.Black)
	drawCapturedPulse(screen, g, alpha)

	text.Draw(screen, `Cheats`, mplusBigFont, newGameColumnBlack+column1*3, row*2+22, color.Black)
	ebitenutil.DrawRect(screen, float64(newGameColumnBlack+column1*3), 242, 250, 6, color.Black)
	text.Draw(screen, `Undo`, mplusMediumFont, newGameColumnBlack+column1*3, row*4, color.Black)
	text.Draw(screen, `Tips`, mplusMediumFont, newGameColumnBlack+column1*3, row*6, color.Black)
}

func drawNewGameOptions(screen *ebiten.Image, g *game, second, pulse, alpha float64) {
	drawStone(screen, g, 0, imgBlack)
	drawStone(screen, g, column1, imgWhite)
	drawSelect(screen, g)
	drawHuman(screen, g)
	drawHotseat(screen, g, alpha)
	drawAI(screen, g)
	drawHighlight(screen, g, second, pulse, alpha)
}
