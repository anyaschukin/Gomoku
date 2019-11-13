package play //gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var newGamecolumnBlack = 140     //100
var newGamecolumnWhite = 1230    // 1000
var newGameBlack2 float64 = 5000 //4700

func drawBlackStone(screen *ebiten.Image, g *game) {
	opBlack := &ebiten.DrawImageOptions{}
	opBlack.GeoM.Translate(float64(newGamecolumnBlack)+340, 50)
	screen.DrawImage(imgBlack, opBlack)
}

func drawWhiteStone(screen *ebiten.Image, g *game) {
	opWhite := &ebiten.DrawImageOptions{}
	opWhite.GeoM.Translate(float64(newGamecolumnWhite)+340, 50)
	screen.DrawImage(imgWhite, opWhite)
}

func drawSelectHuman(screen *ebiten.Image, g *game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 1550)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectAI(screen *ebiten.Image, g *game, shift float64) {
	depth := g.ai0.depth
	if shift != 0 {
		depth = g.ai1.depth
	}
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(950+shift, 1550+(float64(depth)*833))
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectHotseat(screen *ebiten.Image, g *game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 2390)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, g *game, player bool) {
	p := g.ai0
	var shift float64
	if player == true {
		p = g.ai1
		shift = 9100 //7500
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
}

func drawAI(screen *ebiten.Image, g *game, i int) {
	text.Draw(screen, artificial, mplusNormalFont, newGamecolumnBlack, row*(i+2)+50, color.Black)
	text.Draw(screen, strconv.Itoa(i), mplusNormalFont, newGamecolumnBlack+240, row*(i+2)+50, color.Black)
	text.Draw(screen, artificial, mplusNormalFont, newGamecolumnWhite, row*(i+2)+50, color.White)
	text.Draw(screen, strconv.Itoa(i), mplusNormalFont, newGamecolumnWhite+240, row*(i+2)+50, color.White)
}

func drawHuman(screen *ebiten.Image, g *game) {
	text.Draw(screen, humanLower, mplusNormalFont, newGamecolumnBlack+520, row*2+50, color.Black)
	text.Draw(screen, humanLower, mplusNormalFont, newGamecolumnWhite+520, row*2+50, color.White)
}

func drawHotseat(screen *ebiten.Image, g *game) {
	text.Draw(screen, hotseat, mplusNormalFont, newGamecolumnBlack+520, row*3+50, color.Black)
	text.Draw(screen, hotseat, mplusNormalFont, newGamecolumnWhite+520, row*3+50, color.White)
}

func drawNewGameOptions(screen *ebiten.Image, g *game) {
	drawBlackStone(screen, g)
	drawWhiteStone(screen, g)
	drawSelect(screen, g)
	for i := 0; i <= 11; i++ {
		drawAI(screen, g, i)
	}
	drawHuman(screen, g)
	drawHotseat(screen, g)
}
