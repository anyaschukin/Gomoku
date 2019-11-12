package play //gui

import (
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

/// Images
var imggoban *ebiten.Image
var imgBlack *ebiten.Image
var imgWhite *ebiten.Image
var imgRed *ebiten.Image
var imgBlue *ebiten.Image
var imgExit *ebiten.Image
var imgnewGame *ebiten.Image
var imgSelect *ebiten.Image

/// Text
var (
	captured        = `captured:`
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var blackMove = `Black to Move`
var whiteMove = `White to Move`
var human = `Human`
var artificial = `AI Depth`
var hotseat = `Hotseat`
var timer = `Timer:`
var move = `Move:`

/// goban position
var positionWidth float64 = 104.6
var zerox float64 = 838 // Left
var zeroy float64 = 34  // Top
var scale float64 = 0.7

/// New Game position
var newGamex float64 = 3405
var newGamey float64 = 1914

/// Exit position
var exitx float64 = 3210
var exity float64 = 1814

/// Text rows and columns
var row = 100
var columnBlack = 80
var columnWhite = 2050

func init() {
	/// Initialize images
	var err error
	imggoban, _, err = ebitenutil.NewImageFromFile("src/img/goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlack, _, err = ebitenutil.NewImageFromFile("src/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgWhite, _, err = ebitenutil.NewImageFromFile("src/img/white.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgRed, _, err = ebitenutil.NewImageFromFile("src/img/red.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlue, _, err = ebitenutil.NewImageFromFile("src/img/blue.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgExit, _, err = ebitenutil.NewImageFromFile("src/img/exit.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgnewGame, _, err = ebitenutil.NewImageFromFile("src/img/newGame.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgSelect, _, err = ebitenutil.NewImageFromFile("src/img/select.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	/// Initialize text
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    52,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	// mplusBigFont = truetype.NewFace(tt, &truetype.Options{
	// 	Size:    72,
	// 	DPI:     dpi,
	// 	Hinting: font.HintingFull,
	// })
}

func drawgoban(screen *ebiten.Image, g *game) {
	opgoban := &ebiten.DrawImageOptions{}
	opgoban.GeoM.Translate(885, 80)
	opgoban.GeoM.Scale(scale, scale)
	screen.DrawImage(imggoban, opgoban)
}

func drawStones(screen *ebiten.Image, g *game) {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if positionOccupied(coordinate, &g.goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((zerox + (float64(coordinate.x) * positionWidth)), (zeroy + (float64(coordinate.y) * positionWidth)))
				opStone.GeoM.Scale(scale, scale)
				if positionOccupiedByPlayer(coordinate, &g.goban, false) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}
}

func drawPlayerInfo(screen *ebiten.Image, g *game) {
	if g.ai0.hotseat == true {
		text.Draw(screen, hotseat, mplusNormalFont, columnBlack, row*2, color.Black)
	}
	if g.ai0.aiPlayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row*3, color.Black)
		text.Draw(screen, strconv.Itoa(int(g.ai0.depth)), mplusNormalFont, columnBlack+230, row*3, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnBlack, row*3, color.Black)
	}
	if g.ai1.hotseat == true {
		text.Draw(screen, hotseat, mplusNormalFont, columnWhite, row*2, color.White)
	}
	if g.ai1.aiPlayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row*3, color.White)
		text.Draw(screen, strconv.Itoa(int(g.ai1.depth)), mplusNormalFont, columnWhite+230, row*3, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnWhite, row*3, color.White)
	}
}

func drawcaptured(screen *ebiten.Image, g *game) {
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row*4, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.capture0)), mplusNormalFont, 340, row*4, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row*4, color.White)
	text.Draw(screen, strconv.Itoa(int(g.capture1)), mplusNormalFont, 2310, row*4, color.White)
}

func drawTimer(screen *ebiten.Image, g *game) {
	if g.ai0.aiPlayer == true || g.ai0.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnBlack, row*5, color.Black)
		timer, err := time.ParseDuration(g.ai0.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := timer.Truncate(time.Nanosecond).String()
		if timer >= 1000000 {
			truncated = timer.Truncate(time.Millisecond).String()
		} else if timer >= 1000 {
			truncated = timer.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnBlack+180, row*5, color.Black)
	}
	if g.ai1.aiPlayer == true || g.ai1.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnWhite, row*5, color.White)
		timer, err := time.ParseDuration(g.ai1.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := timer.Truncate(time.Nanosecond).String()
		if timer >= 1000000 {
			truncated = timer.Truncate(time.Millisecond).String()
		} else if timer >= 1000 {
			truncated = timer.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnWhite+180, row*5, color.White)
	}
}

func drawMessage(screen *ebiten.Image, g *game) {
	if g.Won == false {
		if g.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, g.Message, mplusNormalFont, columnBlack, row*6, color.Black)
		} else {
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, g.Message, mplusNormalFont, columnWhite, row*6, color.White)
		}
	} else {
		if g.Message == "Black Wins!" {
			text.Draw(screen, g.Message, mplusNormalFont, columnBlack, row*6, color.Black)
		} else {
			text.Draw(screen, g.Message, mplusNormalFont, columnWhite, row*6, color.White)
		}
	}
}

func drawMove(screen *ebiten.Image, g *game) {
	text.Draw(screen, move, mplusNormalFont, columnBlack, row*13, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.move)), mplusNormalFont, columnBlack+160, row*13, color.Black)
}

func drawText(screen *ebiten.Image, g *game) {
	drawPlayerInfo(screen, g)
	drawcaptured(screen, g)
	drawTimer(screen, g)
	drawMessage(screen, g)
	drawMove(screen, g)
}

func drawLastMove(screen *ebiten.Image, g *game) {
	if g.DrawLastMove == true && g.move > 0 {
		opLastMove := &ebiten.DrawImageOptions{}
		opLastMove.GeoM.Translate((zerox + (float64(g.LastMove.x) * positionWidth)), (zeroy + (float64(g.LastMove.y) * positionWidth)))
		opLastMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opLastMove)
		// screen.DrawImage(imgBlue, opLastMove)
	}
}

func drawHotseatSuggestion(screen *ebiten.Image, g *game) {
	if isPlayerHotseat(g) == true && g.Won == false {
		coordinate := g.ai0.suggest
		if g.player == true {
			coordinate = g.ai1.suggest
		}
		opSuggestion := &ebiten.DrawImageOptions{}
		opSuggestion.GeoM.Translate((zerox + (float64(coordinate.x) * positionWidth)), (zeroy + (float64(coordinate.y) * positionWidth)))
		opSuggestion.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opSuggestion)
	}
}

func drawWinMove(screen *ebiten.Image, g *game) {
	if g.Won == true {
		opWinMove := &ebiten.DrawImageOptions{}
		opWinMove.GeoM.Translate((zerox + (float64(g.Winmove.x) * positionWidth)), (zeroy + (float64(g.Winmove.y) * positionWidth)))
		opWinMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opWinMove)
		screen.DrawImage(imgBlue, opWinMove)
	}
}

func drawNewGame(screen *ebiten.Image, g *game) {
	opnewGame := &ebiten.DrawImageOptions{}
	opnewGame.GeoM.Translate(newGamex, newGamey)
	opnewGame.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(imgnewGame, opnewGame)
}

func drawExit(screen *ebiten.Image, g *game) {
	opExit := &ebiten.DrawImageOptions{}
	opExit.GeoM.Translate(exitx, exity)
	opExit.GeoM.Scale(scale, scale)
	screen.DrawImage(imgExit, opExit)
}

func draw(screen *ebiten.Image, g *game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background
	if g.newGame == true {
		drawNewGameOptions(screen, g)
	} else {
		drawgoban(screen, g)
		drawStones(screen, g)
		drawText(screen, g)
		drawLastMove(screen, g)
		drawHotseatSuggestion(screen, g)
		drawWinMove(screen, g)
	}
	drawNewGame(screen, g)
	drawExit(screen, g)
}
