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
	blackMove       = `Black to Move`
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)
var whiteMove = `White to Move`
var humanUpper = `HUMAN`
var humanLower = `Human`
var hotseat = `(Hotseat)`
var artificial = `AI`
var depth = `- depth:`
var captured = `captured:`
var timer = `timer:`
var move = `move:`

/// goban position
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
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func drawGoban(screen *ebiten.Image, g *game) {
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
				opStone.GeoM.Translate((zeroX + (float64(coordinate.x) * positionWidth)), (zeroY + (float64(coordinate.y) * positionWidth)))
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
		text.Draw(screen, hotseat, mplusNormalFont, columnBlack, row*4, color.Black)
	}
	if g.ai0.aiPlayer == true {
		text.Draw(screen, artificial, mplusBigFont, columnBlack, row*5, color.Black)
		ebitenutil.DrawRect(screen, float64(columnBlack-8), 520, 90, 6, color.Black)
		text.Draw(screen, depth, mplusNormalFont, columnBlack+100, row*5-9, color.Black)
		text.Draw(screen, strconv.Itoa(int(g.ai0.depth)), mplusNormalFont, columnBlack+320, row*5-9, color.Black)
	} else {
		text.Draw(screen, humanUpper, mplusBigFont, columnBlack, row*5, color.Black)
		ebitenutil.DrawRect(screen, float64(columnBlack-8), 520, 290, 6, color.Black)
	}
	if g.ai1.hotseat == true {
		text.Draw(screen, hotseat, mplusNormalFont, columnWhite, row*4, color.White)
	}
	if g.ai1.aiPlayer == true {
		text.Draw(screen, artificial, mplusBigFont, columnWhite, row*5, color.White)
		ebitenutil.DrawRect(screen, float64(columnWhite-8), 520, 90, 6, color.White)
		text.Draw(screen, depth, mplusNormalFont, columnWhite+100, row*5-9, color.White)
		text.Draw(screen, strconv.Itoa(int(g.ai1.depth)), mplusNormalFont, columnWhite+320, row*5-9, color.White)
	} else {
		text.Draw(screen, humanUpper, mplusBigFont, columnWhite, row*5, color.White)
		ebitenutil.DrawRect(screen, float64(columnWhite-8), 520, 290, 6, color.White)
	}
}

func drawCaptured(screen *ebiten.Image, g *game) {
	text.Draw(screen, captured, mplusNormalFont, columnBlack, row*6, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.capture0)), mplusNormalFont, columnBlack+270, row*6, color.Black)

	text.Draw(screen, captured, mplusNormalFont, columnWhite, row*6, color.White)
	text.Draw(screen, strconv.Itoa(int(g.capture1)), mplusNormalFont, columnWhite+270, row*6, color.White)
}

func drawTimer(screen *ebiten.Image, g *game) {
	if g.ai0.aiPlayer == true || g.ai0.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnBlack, row*6+75, color.Black)
		elapsed, err := time.ParseDuration(g.ai0.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := elapsed.Truncate(time.Nanosecond).String()
		if elapsed >= 1000000 {
			truncated = elapsed.Truncate(time.Millisecond).String()
		} else if elapsed >= 1000 {
			truncated = elapsed.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnBlack+180, row*6+75, color.Black)
	}
	if g.ai1.aiPlayer == true || g.ai1.hotseat == true {
		text.Draw(screen, timer, mplusNormalFont, columnWhite, row*6+75, color.White)
		elapsed, err := time.ParseDuration(g.ai1.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := elapsed.Truncate(time.Nanosecond).String()
		if elapsed >= 1000000 {
			truncated = elapsed.Truncate(time.Millisecond).String()
		} else if elapsed >= 1000 {
			truncated = elapsed.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, columnWhite+180, row*6+75, color.White)
	}
}

func drawMessage(screen *ebiten.Image, g *game) {
	if g.won == false {
		if g.player == false {
			text.Draw(screen, blackMove, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, g.message, mplusNormalFont, columnBlack, row*2, color.Black)
		} else {
			text.Draw(screen, whiteMove, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, g.message, mplusNormalFont, columnWhite, row*2, color.White)
		}
	} else {
		if g.message == "Black Wins!" {
			text.Draw(screen, g.message, mplusBigFont, columnBlack, row*2, color.Black)
		} else {
			text.Draw(screen, g.message, mplusBigFont, columnWhite, row*2, color.White)
		}
	}
}

func drawMove(screen *ebiten.Image, g *game) {
	text.Draw(screen, move, mplusNormalFont, columnBlack, row*13, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.move)), mplusNormalFont, columnBlack+160, row*13, color.Black)
}

func drawText(screen *ebiten.Image, g *game) {
	drawPlayerInfo(screen, g)
	drawCaptured(screen, g)
	drawTimer(screen, g)
	drawMessage(screen, g)
	drawMove(screen, g)
}

func drawLastMove(screen *ebiten.Image, g *game) {
	if g.drawLastMove == true && g.move > 0 {
		opLastMove := &ebiten.DrawImageOptions{}
		opLastMove.GeoM.Translate((zeroX + (float64(g.lastMove.x) * positionWidth)), (zeroY + (float64(g.lastMove.y) * positionWidth)))
		opLastMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opLastMove)
		// screen.DrawImage(imgBlue, opLastMove)
	}
}

func drawHotseatSuggestion(screen *ebiten.Image, g *game) {
	if isPlayerHotseat(g) == true && g.won == false {
		coordinate := g.ai0.suggest
		if g.player == true {
			coordinate = g.ai1.suggest
		}
		opSuggestion := &ebiten.DrawImageOptions{}
		opSuggestion.GeoM.Translate((zeroX + (float64(coordinate.x) * positionWidth)), (zeroY + (float64(coordinate.y) * positionWidth)))
		opSuggestion.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opSuggestion)
	}
}

func drawWinMove(screen *ebiten.Image, g *game) {
	if g.won == true {
		opWinMove := &ebiten.DrawImageOptions{}
		opWinMove.GeoM.Translate((zeroX + (float64(g.winMove.x) * positionWidth)), (zeroY + (float64(g.winMove.y) * positionWidth)))
		opWinMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opWinMove)
		screen.DrawImage(imgBlue, opWinMove)
	}
}

func drawNewGame(screen *ebiten.Image, g *game) {
	opnewGame := &ebiten.DrawImageOptions{}
	opnewGame.GeoM.Translate(newGameX, newGameY)
	opnewGame.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(imgnewGame, opnewGame)
}

func drawExit(screen *ebiten.Image, g *game) {
	opExit := &ebiten.DrawImageOptions{}
	opExit.GeoM.Translate(exitX, exitY)
	opExit.GeoM.Scale(scale, scale)
	screen.DrawImage(imgExit, opExit)
}

func draw(screen *ebiten.Image, g *game) {
	screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff}) /// Draw background
	if g.newGame == true {
		drawNewGameOptions(screen, g)
	} else {
		drawGoban(screen, g)
		drawStones(screen, g)
		drawText(screen, g)
		drawLastMove(screen, g)
		drawHotseatSuggestion(screen, g)
		drawWinMove(screen, g)
	}
	drawNewGame(screen, g)
	drawExit(screen, g)
}
