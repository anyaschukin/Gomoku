package play //gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var newGameColumnBlack = 140
var newGameColumnWhite = 1000
var newGameColumnHighlight = 1860
var newGameBlack2 float64 = 5000 //4700/////
var scaleSelect = 0.23

func drawBlackStone(screen *ebiten.Image, g *game) {
	opBlack := &ebiten.DrawImageOptions{}
	opBlack.GeoM.Translate(float64(newGameColumnBlack)+130, 150)
	screen.DrawImage(imgBlack, opBlack)
}

func drawWhiteStone(screen *ebiten.Image, g *game) {
	opWhite := &ebiten.DrawImageOptions{}
	opWhite.GeoM.Translate(float64(newGameColumnWhite)+130, 150)
	screen.DrawImage(imgWhite, opWhite)
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

func drawSelectLastMove(screen *ebiten.Image, g *game) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(1760/scaleSelect, 295/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectWinMove(screen *ebiten.Image, g *game) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(1760/scaleSelect, 492/scaleSelect)
	opSelect.GeoM.Scale(scaleSelect, scaleSelect)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, g *game, player bool) {
	p := g.ai0
	var shift float64
	if player == true {
		p = g.ai1
		shift = 860 / scaleSelect //7500
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

func drawAI(screen *ebiten.Image, g *game) {
	text.Draw(screen, `AI`, mplusBigFont, newGameColumnBlack, row*8, color.Black)
	text.Draw(screen, `- depth`, mplusBigFont, newGameColumnBlack+100, row*8, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.ai0.depth)), mplusBigFont, newGameColumnBlack+400, row*8, color.Black)
	text.Draw(screen, string('⇧'), mplusBigFont, newGameColumnBlack+387, row*7, color.Black)
	text.Draw(screen, string('⇩'), mplusBigFont, newGameColumnBlack+387, row*9, color.Black)

	text.Draw(screen, `AI`, mplusBigFont, newGameColumnWhite, row*8, color.White)
	text.Draw(screen, `- depth`, mplusBigFont, newGameColumnWhite+100, row*8, color.White)
	text.Draw(screen, strconv.Itoa(int(g.ai1.depth)), mplusBigFont, newGameColumnWhite+400, row*8, color.White)
	text.Draw(screen, string('⇧'), mplusBigFont, newGameColumnWhite+387, row*7, color.White)
	text.Draw(screen, string('⇩'), mplusBigFont, newGameColumnWhite+387, row*9, color.White)
}

func drawHuman(screen *ebiten.Image, g *game) {
	text.Draw(screen, `Human`, mplusBigFont, newGameColumnBlack, row*4, color.Black)
	text.Draw(screen, `Human`, mplusBigFont, newGameColumnWhite, row*4, color.White)
}

func drawHotseat(screen *ebiten.Image, g *game) {
	text.Draw(screen, `Hotseat`, mplusBigFont, newGameColumnBlack, row*6, color.Black)
	text.Draw(screen, `Hotseat`, mplusBigFont, newGameColumnWhite, row*6, color.White)
}

func drawHighlight(screen *ebiten.Image, g *game) {
	ebitenutil.DrawRect(screen, float64(newGameColumnWhite)+800, 242, 320, 6, color.Black)
	text.Draw(screen, `Highlight`, mplusBigFont, newGameColumnWhite+800, row*2+22, color.Black)
	text.Draw(screen, `Last Move`, mplusBigFont, newGameColumnWhite+800, row*4, color.Black)
	text.Draw(screen, `Win Move`, mplusBigFont, newGameColumnWhite+800, row*6, color.Black)
}

func drawNewGameOptions(screen *ebiten.Image, g *game) {
	drawBlackStone(screen, g)
	drawWhiteStone(screen, g)
	drawSelect(screen, g)
	drawHuman(screen, g)
	drawHotseat(screen, g)
	drawAI(screen, g)
	drawHighlight(screen, g)
}
