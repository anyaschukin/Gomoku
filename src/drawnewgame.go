package play //gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var newGameColumnBlack = 140
var column1 = 860
var scaleSelect = 0.23

func drawStone(screen *ebiten.Image, g *game, column int, stone *ebiten.Image) {
	opStone := &ebiten.DrawImageOptions{}
	opStone.GeoM.Translate(float64(newGameColumnBlack+column)+130, 150)
	screen.DrawImage(stone, opStone)
}

func drawSelectHuman(screen *ebiten.Image, g *game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(90/scaleSelect+shift, 292/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectAI(screen *ebiten.Image, g *game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(90/scaleSelect+shift, 698/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectHotseat(screen *ebiten.Image, g *game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(90/scaleSelect+shift, 495/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, g *game, player bool) {
	p := g.ai0
	var shift float64
	if player == true {
		p = g.ai1
		shift = float64(column1) / scaleSelect
	}
	if p.hotseat == true {
		drawSelectHotseat(screen, g, shift)
		drawSelectHuman(screen, g, shift)
		drawSelectAI(screen, g, shift)
	} else if p.aiPlayer == false {
		drawSelectHuman(screen, g, shift)
	} else {
		drawSelectAI(screen, g, shift)
	}
}

func drawSelectLastMove(screen *ebiten.Image, g *game) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate((float64(newGameColumnBlack+column1*2)-45)/scaleSelect, 295/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectWinMove(screen *ebiten.Image, g *game) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate((float64(newGameColumnBlack+column1*2)-45)/scaleSelect, 492/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelect(screen *ebiten.Image, g *game) {
	drawSelectPlayer(screen, g, false)
	drawSelectPlayer(screen, g, true)
	if g.drawLastMove == true {
		drawSelectLastMove(screen, g)
	}
	if g.drawWinMove == true {
		drawSelectWinMove(screen, g)
	}
}

func drawAIplayer(screen *ebiten.Image, g *game, depth uint8, column int, color color.Color) {
	text.Draw(screen, `AI - depth`, mplusBigFont, newGameColumnBlack+column, row*8, color)
	text.Draw(screen, strconv.Itoa(int(depth)), mplusBigFont, newGameColumnBlack+column+400, row*8, color)
	text.Draw(screen, string('⇧'), mplusBigFont, newGameColumnBlack+column+387, row*7, color)
	text.Draw(screen, string('⇩'), mplusBigFont, newGameColumnBlack+column+387, row*9, color)
}

func drawAI(screen *ebiten.Image, g *game) {
	drawAIplayer(screen, g, g.ai0.depth, 0, color.Black)
	drawAIplayer(screen, g, g.ai1.depth, column1, color.White)
}

func drawHuman(screen *ebiten.Image, g *game) {
	text.Draw(screen, `Human`, mplusBigFont, newGameColumnBlack, row*4, color.Black)
	text.Draw(screen, `Human`, mplusBigFont, newGameColumnBlack+column1, row*4, color.White)
}

func drawHotseat(screen *ebiten.Image, g *game) {
	text.Draw(screen, `Hotseat`, mplusBigFont, newGameColumnBlack, row*6, color.Black)
	text.Draw(screen, `Hotseat`, mplusBigFont, newGameColumnBlack+column1, row*6, color.White)
}

func drawHighlight(screen *ebiten.Image, g *game) {
	ebitenutil.DrawRect(screen, float64(newGameColumnBlack+column1*2), 242, 320, 6, color.Black)
	text.Draw(screen, `Highlight`, mplusBigFont, newGameColumnBlack+column1*2, row*2+22, color.Black)
	text.Draw(screen, `Last Move`, mplusBigFont, newGameColumnBlack+column1*2, row*4, color.Black)
	text.Draw(screen, `Win Move`, mplusBigFont, newGameColumnBlack+column1*2, row*6, color.Black)
}

func drawNewGameOptions(screen *ebiten.Image, g *game) {
	drawStone(screen, g, 0, imgBlack)
	drawStone(screen, g, column1, imgWhite)
	drawSelect(screen, g)
	drawHuman(screen, g)
	drawHotseat(screen, g)
	drawAI(screen, g)
	drawHighlight(screen, g)
}
