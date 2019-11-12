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
var imgGoban *ebiten.Image
var imgBlack *ebiten.Image
var imgWhite *ebiten.Image
var imgRed *ebiten.Image
var imgBlue *ebiten.Image
var imgExit *ebiten.Image
var imgNewGame *ebiten.Image
var imgSelect *ebiten.Image

/// Text
var (
	captured        = `Captured:`
	mplusNormalFont font.Face
	// mpluBigFont     font.Face
)
var blackMove = `Black to Move`
var whiteMove = `White to Move`
var human = `Human`
var artificial = `AI Depth`
var Hotseat = `Hotseat`
var timer = `Timer:`
var move = `Move:`

/// Goban position
var positionWidth float64 = 104.6
var zeroX float64 = 838 // Left
var zeroY float64 = 34  // Top
var scale float64 = 0.7

/// New Game position
var newGameX float64 = 3405
var newGameY float64 = 1914

/// Exit position
var exitX float64 = 3210
var exitY float64 = 1814

/// Text rows and columns
var row = 100
var columnBlack = 80
var columnWhite = 2050

func init() {
	/// Initialize images
	var err error
	imgGoban, _, err = ebitenutil.NewImageFromFile("src/img/goban.png", ebiten.FilterDefault)
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
	imgNewGame, _, err = ebitenutil.NewImageFromFile("src/img/newGame.png", ebiten.FilterDefault)
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
			coordinate := Coordinate{y, x}
			if PositionOccupied(coordinate, &G.Goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((zeroX + (float64(coordinate.X) * positionWidth)), (zeroY + (float64(coordinate.Y) * positionWidth)))
				opStone.GeoM.Scale(scale, scale)
				if PositionOccupiedByPlayer(coordinate, &G.Goban, false) == true {
					screen.DrawImage(imgBlack, opStone)
				} else {
					screen.DrawImage(imgWhite, opStone)
				}
			}
		}
	}
}

func drawPlayerInfo(screen *ebiten.Image, G *Game) {
	if G.Ai0.Hotseat == true {
		text.Draw(screen, Hotseat, mplusNormalFont, columnBlack, row*2, color.Black)
	}
	if G.Ai0.AiPlayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnBlack, row*3, color.Black)
		text.Draw(screen, strconv.Itoa(int(G.Ai0.Depth)), mplusNormalFont, columnBlack+230, row*3, color.Black)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnBlack, row*3, color.Black)
	}
	if G.Ai1.Hotseat == true {
		text.Draw(screen, Hotseat, mplusNormalFont, columnWhite, row*2, color.White)
	}
	if G.Ai1.AiPlayer == true {
		text.Draw(screen, artificial, mplusNormalFont, columnWhite, row*3, color.White)
		text.Draw(screen, strconv.Itoa(int(G.Ai1.Depth)), mplusNormalFont, columnWhite+230, row*3, color.White)
	} else {
		text.Draw(screen, human, mplusNormalFont, columnWhite, row*3, color.White)
	}
}

func drawCaptured(screen *ebiten.Image, G *Game) {
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row*4, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.Capture0)), mplusNormalFont, 340, row*4, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row*4, color.White)
	text.Draw(screen, strconv.Itoa(int(G.Capture1)), mplusNormalFont, 2310, row*4, color.White)
}

func drawTimer(screen *ebiten.Image, G *Game) {
	if G.Ai0.AiPlayer == true || G.Ai0.Hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnBlack, row*5, color.Black)
		timer, err := time.ParseDuration(G.Ai0.Timer.String())
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
	if G.Ai1.AiPlayer == true || G.Ai1.Hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnWhite, row*5, color.White)
		timer, err := time.ParseDuration(G.Ai1.Timer.String())
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

func drawMessage(screen *ebiten.Image, G *Game) {
	if G.Won == false {
		if G.Player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, G.Message, mplusNormalFont, columnBlack, row*6, color.Black)
		} else {
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, G.Message, mplusNormalFont, columnWhite, row*6, color.White)
		}
	} else {
		if G.Message == "Black Wins!" {
			text.Draw(screen, G.Message, mplusNormalFont, columnBlack, row*6, color.Black)
		} else {
			text.Draw(screen, G.Message, mplusNormalFont, columnWhite, row*6, color.White)
		}
	}
}

func drawMove(screen *ebiten.Image, G *Game) {
	text.Draw(screen, move, mplusNormalFont, columnBlack, row*13, color.Black)
	text.Draw(screen, strconv.Itoa(int(G.Move)), mplusNormalFont, columnBlack+160, row*13, color.Black)
}

func drawText(screen *ebiten.Image, G *Game) {
	drawPlayerInfo(screen, G)
	drawCaptured(screen, G)
	drawTimer(screen, G)
	drawMessage(screen, G)
	drawMove(screen, G)
}

func drawLastMove(screen *ebiten.Image, G *Game) {
	if G.DrawLastMove == true && G.Move > 0 {
		opLastMove := &ebiten.DrawImageOptions{}
		opLastMove.GeoM.Translate((zeroX + (float64(G.LastMove.X) * positionWidth)), (zeroY + (float64(G.LastMove.Y) * positionWidth)))
		opLastMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opLastMove)
		// screen.DrawImage(imgBlue, opLastMove)
	}
}

func drawHotseatSuggestion(screen *ebiten.Image, G *Game) {
	if IsPlayerHotseat(G) == true && G.Won == false {
		coordinate := G.Ai0.Suggest
		if G.Player == true {
			coordinate = G.Ai1.Suggest
		}
		opSuggestion := &ebiten.DrawImageOptions{}
		opSuggestion.GeoM.Translate((zeroX + (float64(coordinate.X) * positionWidth)), (zeroY + (float64(coordinate.Y) * positionWidth)))
		opSuggestion.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opSuggestion)
	}
}

func drawWinMove(screen *ebiten.Image, G *Game) {
	if G.Won == true {
		opWinMove := &ebiten.DrawImageOptions{}
		opWinMove.GeoM.Translate((zeroX + (float64(G.Winmove.X) * positionWidth)), (zeroY + (float64(G.Winmove.Y) * positionWidth)))
		opWinMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opWinMove)
		screen.DrawImage(imgBlue, opWinMove)
	}
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

func draw(screen *ebiten.Image, G *Game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background
	if G.NewGame == true {
		drawNewGameOptions(screen, G)
	} else {
		drawGoban(screen, G)
		drawStones(screen, G)
		drawText(screen, G)
		drawLastMove(screen, G)
		drawHotseatSuggestion(screen, G)
		drawWinMove(screen, G)
	}
	drawNewGame(screen, G)
	drawExit(screen, G)
}
