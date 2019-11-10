package play //gui

import (
	"image/color"
	"strconv"
	// play "Gomoku/play"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

var newGamecolumnBlack = 140     //100
var newGamecolumnWhite = 1230    // 1000
var newGameBlack2 float64 = 5000 //4700

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
	Depth := G.Ai0.Depth
	if shift != 0 {
		Depth = G.Ai1.Depth
	}
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(950+shift, 1550+(float64(Depth)*833))
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectHotseat(screen *ebiten.Image, G *Game, shift float64) {
	opSelect := &ebiten.DrawImageOptions{}
	opSelect.GeoM.Translate(newGameBlack2+shift, 2390)
	opSelect.GeoM.Scale(0.12, 0.12)
	screen.DrawImage(imgSelect, opSelect)
}

func drawSelectPlayer(screen *ebiten.Image, G *Game, Player bool) {
	p := G.Ai0
	var shift float64
	if Player == true {
		p = G.Ai1
		shift = 9100 //7500
	}
	if p.Hotseat == true {
		drawSelectHotseat(screen, G, shift)
		drawSelectHuman(screen, G, shift)
		drawSelectAI(screen, G, shift)
	} else if p.AiPlayer == false {
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
	text.Draw(screen, Hotseat, mplusNormalFont, newGamecolumnBlack+520, row*3+50, color.Black)
	text.Draw(screen, Hotseat, mplusNormalFont, newGamecolumnWhite+520, row*3+50, color.White)
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
