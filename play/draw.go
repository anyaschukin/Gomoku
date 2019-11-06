package play //gui

import (
	"image/color"
	"log"
	"strconv"
	"time"
	// "fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

/// Images
var imgGoban *ebiten.Image
var imgBlack *ebiten.Image
var imgRed *ebiten.Image
var imgBlue *ebiten.Image
var imgWhite *ebiten.Image
var imgExit *ebiten.Image
var imgNewGame *ebiten.Image
var imgSelect *ebiten.Image

/// Text
var (
	captured        = `Captured: `
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var blackMove = `Black to Move`
var whiteMove = `White to Move`
var human = `Human`
var artificial = `AI depth`
var hotseat = `Hotseat`
var timer = `Timer:`

/// Goban positions
var positionWidth float64 = 104.6
var zeroX float64 = 838 // Left
var zeroY float64 = 34  // Top
var scale float64 = 0.7

/// New Game position
var newGameX float64 = 3405
var newGameY float64 = 1914
var newGameBlack2 float64 = 5000

/// Exit position
var exitX float64 = 3210
var exitY float64 = 1814

/// Text rows
var row = 100

/// New Game Columns
var columnBlack = 80   // screen indent
var columnWhite = 2050 // screen indent

func init() {
	/// Init images
	var err error
	imgGoban, _, err = ebitenutil.NewImageFromFile("GUI/img/goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlack, _, err = ebitenutil.NewImageFromFile("GUI/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgWhite, _, err = ebitenutil.NewImageFromFile("GUI/img/white.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgRed, _, err = ebitenutil.NewImageFromFile("GUI/img/red.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlue, _, err = ebitenutil.NewImageFromFile("GUI/img/blue.png", ebiten.FilterDefault)////////// use it or lose it////////
	if err != nil {
		log.Fatal(err)
	}
	imgExit, _, err = ebitenutil.NewImageFromFile("GUI/img/exit.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgNewGame, _, err = ebitenutil.NewImageFromFile("GUI/img/newGame.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgSelect, _, err = ebitenutil.NewImageFromFile("GUI/img/select.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	/// Init text
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

func drawGoban(screen *ebiten.Image, G *Game) {
	opGoban := &ebiten.DrawImageOptions{}
	opGoban.GeoM.Translate(885, 80)
	opGoban.GeoM.Scale(scale, scale)
	screen.DrawImage(imgGoban, opGoban)
}

func drawStones(screen *ebiten.Image, G *Game) {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if PositionOccupied(coordinate, &G.goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((zeroX + (float64(coordinate.x) * positionWidth)), (zeroY + (float64(coordinate.y) * positionWidth)))
				opStone.GeoM.Scale(scale, scale)
				if PositionOccupiedByPlayer(coordinate, &G.goban, false) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}
}

func drawPlayerInfo(screen *ebiten.Image, G *Game) {
	if G.ai0.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row*2, color.Black)
		text.Draw(screen, strconv.Itoa(int(G.ai0.depth)), mplusNormalFont, columnBlack+230, row*2, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnBlack, row*2, color.Black)
	}
	if G.ai1.aiplayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row*2, color.White)
		text.Draw(screen, strconv.Itoa(int(G.ai1.depth)), mplusNormalFont, columnWhite+230, row*2, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnWhite, row*2, color.White)
	}
}

func drawCaptured(screen *ebiten.Image, G *Game) {
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row*3, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.capture0)), mplusNormalFont, 340, row*3, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row*3, color.White)
	text.Draw(screen, strconv.Itoa(int(G.capture1)), mplusNormalFont, 2310, row*3, color.White)
}

func drawTimer(screen *ebiten.Image, G *Game) {
	if G.ai0.aiplayer == true || G.ai0.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnBlack, row*4, color.Black)
		timer, err := time.ParseDuration(G.ai0.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := timer.Truncate(time.Nanosecond).String()
		if timer >= 1000000 {
			truncated = timer.Truncate(time.Millisecond).String()
		} else if timer >= 1000 {
			truncated = timer.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnBlack+180, row*4, color.Black)
	}
	if G.ai1.aiplayer == true || G.ai1.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnWhite, row*4, color.White)
		timer, err := time.ParseDuration(G.ai1.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := timer.Truncate(time.Nanosecond).String()
		if timer >= 1000000 {
			truncated = timer.Truncate(time.Millisecond).String()
		} else if timer >= 1000 {
			truncated = timer.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnWhite+180, row*4, color.White)
	}
}

func drawMessage(screen *ebiten.Image, G *Game) {
	if G.won == false {
		if G.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row*5, color.Black)
		} else {
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row*5, color.White)
		}
	} else {
		if G.message == "Black Wins!" {
			text.Draw(screen, G.message, mplusNormalFont, columnBlack, row*5, color.Black)
		} else {
			text.Draw(screen, G.message, mplusNormalFont, columnWhite, row*5, color.White)
		}
	}
}

func drawText(screen *ebiten.Image, G *Game) {
	drawPlayerInfo(screen, G)
	drawCaptured(screen, G)
	drawTimer(screen, G)
	drawMessage(screen, G)
}

func drawNewGame(screen *ebiten.Image, G *Game) {
	opNewGame := &ebiten.DrawImageOptions{}
	opNewGame.GeoM.Translate(newGameX, newGameY)
	opNewGame.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(imgNewGame, opNewGame)
}

func drawExit(screen *ebiten.Image, G *Game) {
	opExit := &ebiten.DrawImageOptions{}
	opExit.GeoM.Translate(exitX, exitY)
	opExit.GeoM.Scale(scale, scale)
	screen.DrawImage(imgExit, opExit)
}

func drawHotseatSuggestion(screen *ebiten.Image, G *Game) {
	if isPlayerHotseat(G) == true {
		coordinate := G.ai0.suggest
		if G.player == true {
			coordinate = G.ai1.suggest
		}
		opSuggestion := &ebiten.DrawImageOptions{}
		opSuggestion.GeoM.Translate((zeroX + (float64(coordinate.x) * positionWidth)), (zeroY + (float64(coordinate.y) * positionWidth)))
		opSuggestion.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opSuggestion)
	}
}

func drawWinMove(screen *ebiten.Image, G *Game) {
	if G.won == true {
		if capturedTen(G) == false {
			opWinMove := &ebiten.DrawImageOptions{}
			opWinMove.GeoM.Translate((zeroX + (float64(G.align5.winmove.x) * positionWidth)), (zeroY + (float64(G.align5.winmove.y) * positionWidth)))
			opWinMove.GeoM.Scale(scale, scale)
			screen.DrawImage(imgRed, opWinMove)
			screen.DrawImage(imgBlue, opWinMove)
		}
	}
}

func draw(screen *ebiten.Image, G *Game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background
	if G.newGame == true {
		drawNewGameOptions(screen, G)
	} else {
		drawGoban(screen, G)
		drawStones(screen, G)
		drawText(screen, G)
		drawHotseatSuggestion(screen, G)
		drawWinMove(screen, G)
	}
	drawNewGame(screen, G)
	drawExit(screen, G)
}
