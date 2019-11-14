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
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

/// goban position
var positionWidth = 104.6
var gobanX float64 = 838 // Left
var gobanY float64 = 34  // Top
var scale = 0.7

/// New Game position
var newGameX float64 = 3405
var newGameY float64 = 1914
var newGameScale = 0.6

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
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func drawGoban(screen *ebiten.Image, g *game) {
	opGoban := &ebiten.DrawImageOptions{}
	opGoban.GeoM.Translate(885, 80)
	opGoban.GeoM.Scale(scale, scale)
	screen.DrawImage(imgGoban, opGoban)
}

func drawStones(screen *ebiten.Image, g *game) {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if positionOccupied(coordinate, &g.goban) == true {
				opStone := &ebiten.DrawImageOptions{}
				opStone.GeoM.Translate((gobanX + (float64(coordinate.x) * positionWidth)), (gobanY + (float64(coordinate.y) * positionWidth)))
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

func drawPlayerInfo(screen *ebiten.Image, g *game, p ai, column int, color color.Color) {
	if p.hotseat == true {
		text.Draw(screen, `(Hotseat)`, mplusNormalFont, column, row*4, color)
	}
	if p.aiPlayer == true {
		text.Draw(screen, `AI`, mplusBigFont, column, row*5, color)
		ebitenutil.DrawRect(screen, float64(column-8), 520, 90, 6, color)
		text.Draw(screen, `- depth:`, mplusNormalFont, column+100, row*5-9, color)
		text.Draw(screen, strconv.Itoa(int(p.depth)), mplusNormalFont, column+320, row*5-9, color)
	} else {
		text.Draw(screen, `HUMAN`, mplusBigFont, column, row*5, color)
		ebitenutil.DrawRect(screen, float64(column-8), 520, 290, 6, color)
	}
}

func drawCaptured(screen *ebiten.Image, g *game, captured uint8, column int, color color.Color) {
	text.Draw(screen, `captured:`, mplusNormalFont, column, row*6, color)
	text.Draw(screen, strconv.Itoa(int(captured)), mplusNormalFont, column+270, row*6, color)
}

func drawTimer(screen *ebiten.Image, g *game, p ai, column int, color color.Color) {
	if p.aiPlayer == true || p.hotseat == true {
		text.Draw(screen, `timer:`, mplusNormalFont, column, row*6+75, color)
		elapsed, err := time.ParseDuration(p.timer.String())
		if err != nil {
			panic(err)
		}
		truncated := elapsed.Truncate(time.Nanosecond).String()
		if elapsed >= 1000000 {
			truncated = elapsed.Truncate(time.Millisecond).String()
		} else if elapsed >= 1000 {
			truncated = elapsed.Truncate(time.Microsecond).String()
		}
		text.Draw(screen, truncated, mplusNormalFont, column+180, row*6+75, color)
	}
}

func drawPlayerText(screen *ebiten.Image, g *game, player bool) {
	var c color.Color
	column := columnBlack
	p := g.ai0
	captured := g.capture0
	c = color.Black
	if player == true {
		column = columnWhite
		p = g.ai1
		captured = g.capture1
		c = color.White
	}
	drawPlayerInfo(screen, g, p, column, c)
	drawCaptured(screen, g, captured, column, c)
	drawTimer(screen, g, p, column, c)
}

func drawMessage(screen *ebiten.Image, g *game) {
	if g.won == false {
		if g.player == false {
			text.Draw(screen, `Black to Move`, mplusNormalFont, columnBlack, row, color.Black)
			text.Draw(screen, g.message, mplusNormalFont, columnBlack, row*2, color.Black)
		} else {
			text.Draw(screen, `White to Move`, mplusNormalFont, columnWhite, row, color.White)
			text.Draw(screen, g.message, mplusNormalFont, columnWhite, row*2, color.White)
		}
	} else {
		if g.message == "Black Wins!" {
			text.Draw(screen, g.message, mplusBigFont, columnBlack, row*1+50, color.Black)
		} else {
			text.Draw(screen, g.message, mplusBigFont, columnWhite, row*1+50, color.White)
		}
	}
}

func drawMove(screen *ebiten.Image, g *game) {
	text.Draw(screen, `move:`, mplusNormalFont, columnBlack, row*13, color.Black)
	text.Draw(screen, strconv.Itoa(int(g.move)), mplusNormalFont, columnBlack+160, row*13, color.Black)
}

func drawText(screen *ebiten.Image, g *game) {
	drawPlayerText(screen, g, false)
	drawPlayerText(screen, g, true)
	drawMessage(screen, g)
	drawMove(screen, g)
}

func drawLastMove(screen *ebiten.Image, g *game) {
	if g.drawLastMove == true && g.move > 0 {
		opLastMove := &ebiten.DrawImageOptions{}
		opLastMove.GeoM.Translate((gobanX + (float64(g.lastMove.x) * positionWidth)), (gobanY + (float64(g.lastMove.y) * positionWidth)))
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
		opSuggestion.GeoM.Translate((gobanX + (float64(coordinate.x) * positionWidth)), (gobanY + (float64(coordinate.y) * positionWidth)))
		opSuggestion.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opSuggestion)
	}
}

func drawWinMove(screen *ebiten.Image, g *game) {
	if g.won == true && g.drawWinMove == true {
		opWinMove := &ebiten.DrawImageOptions{}
		opWinMove.GeoM.Translate((gobanX + (float64(g.winMove.x) * positionWidth)), (gobanY + (float64(g.winMove.y) * positionWidth)))
		opWinMove.GeoM.Scale(scale, scale)
		screen.DrawImage(imgRed, opWinMove)
		screen.DrawImage(imgBlue, opWinMove)
	}
}

func drawNewGame(screen *ebiten.Image, g *game) {
	opnewGame := &ebiten.DrawImageOptions{}
	opnewGame.GeoM.Translate(newGameX, newGameY)
	opnewGame.GeoM.Scale(newGameScale, newGameScale)
	screen.DrawImage(imgNewGame, opnewGame)
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
