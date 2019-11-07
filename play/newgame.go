package play //gui

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var newGamecolumnBlack = 140
var newGamecolumnWhite = 1230

func drawBlackStone(screen *ebiten.Image, G *Game) {
	opBlack := &ebiten.DrawImageOptions{}
	opBlack.GeoM.Translate(float64(newGamecolumnBlack)+340, 50)
	screen.DrawImage(imgBlack, opBlack)
}

func drawWhiteStone(screen *ebiten.Image, G *Game) {
	opWhite := &ebiten.DrawImageOptions{}
	opWhite.GeoM.Translate(float64(newGamecolumnWhite)+340, 50)
	screen.DrawImage(imgWhite, opWhite)
}

func drawSelectHuman(screen *ebiten.Image, G *Game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 1550)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectAI(screen *ebiten.Image, G *Game, shift float64) {
	depth := G.ai0.depth
	if shift != 0 {
		depth = G.ai1.depth
	}
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(950+shift, 1550+(float64(depth)*833))
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectHotseat(screen *ebiten.Image, G *Game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 2390)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, G *Game, player bool) {
	p := G.ai0
	var shift float64
	if player == true {
		p = G.ai1
		shift = 9100
	}
	if p.hotseat == true {
		drawSelectHotseat(screen, G, shift)
		drawSelectHuman(screen, G, shift)
		drawSelectAI(screen, G, shift)
	} else if p.aiplayer == false {
		drawSelectHuman(screen, G, shift)
	} else {
		drawSelectAI(screen, G, shift)
	}
}

func drawSelect(screen *ebiten.Image, G *Game) {
	drawSelectPlayer(screen, G, false)
	drawSelectPlayer(screen, G, true)
}

func drawAI(screen *ebiten.Image, G *Game, i int) {
	text.Draw(screen, artificial, mplusNormalFont, newGamecolumnBlack, row*(i+2)+50, color.Black)
	text.Draw(screen, strconv.Itoa(i), mplusNormalFont, newGamecolumnBlack+230, row*(i+2)+50, color.Black)
	text.Draw(screen, artificial, mplusNormalFont, newGamecolumnWhite, row*(i+2)+50, color.White)
	text.Draw(screen, strconv.Itoa(i), mplusNormalFont, newGamecolumnWhite+230, row*(i+2)+50, color.White)
}

func drawHuman(screen *ebiten.Image, G *Game) {
	text.Draw(screen, human, mplusNormalFont, newGamecolumnBlack+520, row*2+50, color.Black)
	text.Draw(screen, human, mplusNormalFont, newGamecolumnWhite+520, row*2+50, color.White)
}

func drawHotseat(screen *ebiten.Image, G *Game) {
	text.Draw(screen, hotseat, mplusNormalFont, newGamecolumnBlack+520, row*3+50, color.Black)
	text.Draw(screen, hotseat, mplusNormalFont, newGamecolumnWhite+520, row*3+50, color.White)
}

func drawNewGameOptions(screen *ebiten.Image, G *Game) {
	drawBlackStone(screen, G)
	drawWhiteStone(screen, G)
	drawSelect(screen, G)
	for i := 0; i <= 11; i++ {
		drawAI(screen, G, i)
	}
	drawHuman(screen, G)
	drawHotseat(screen, G)
}
