package gomoku

import (
	"image/color"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

var dogeMode bool

/// Goban position
var positionWidth = 104.6
var gobanX float64 = 838 // Left
var gobanY float64 = 34  // Top
var scale = 0.7

/// Text rows and columns
const row = 100
const columnBlack = 80
const columnWhite = 2050

/// New Game position
const newGameX float64 = 3405
const newGameY float64 = 1914

var newGameScale = 0.6

/// Exit position
const exitX float64 = 3210
const exitY float64 = 1814

/// Undo position
const undoX = float64(columnBlack)
const undoY = float64(row * 15)

func opImage(x, y, scale float64) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(scale, scale)
	return op
}

func drawImage(screen, image *ebiten.Image, x, y, scale float64) {
	op := opImage(x, y, scale)
	screen.DrawImage(image, op)
}

func drawImagePulse(screen, image *ebiten.Image, x, y, scale, alpha float64) {
	op := opImage(x, y, scale)
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(image, op)
}

func drawBackground(screen *ebiten.Image) {
	if dogeMode == true {
		drawImage(screen, imgBackground, 0, 0, 1)
	} else {
		if background == nil {
			screen.Fill(color.RGBA{0xaf, 0xaf, 0xff, 0xff})
		} else {
			screen.Fill(background)
		}
	}
}

func drawGoban(screen *ebiten.Image, g *game) {
	drawImage(screen, imgGoban, 885, 80, scale)
}

func stoneX(position int8) float64 {
	x := gobanX + (float64(position) * positionWidth)
	return x
}

func stoneY(position int8) float64 {
	y := gobanY + (float64(position) * positionWidth)
	return y
}

func drawStones(screen *ebiten.Image, g *game) {
	var y int8
	var x int8
	for y = 0; y < 19; y++ {
		for x = 0; x < 19; x++ {
			coordinate := coordinate{y, x}
			if positionOccupied(coordinate, &g.goban) == true {
				if positionOccupiedByPlayer(coordinate, &g.goban, false) == true {
					drawImage(screen, imgBlack, stoneX(coordinate.x), stoneY(coordinate.y), scale)
				} else {
					drawImage(screen, imgWhite, stoneX(coordinate.x), stoneY(coordinate.y), scale)
				}
			}
		}
	}
}

func drawHotseatSuggestion(screen *ebiten.Image, g *game, alpha float64) {
	if isPlayerHotseat(g) == true && g.won == false {
		coordinate := g.ai0.suggest
		if g.player == true {
			coordinate = g.ai1.suggest
		}
		if g.player == false {
			drawImagePulse(screen, imgBlack, stoneX(coordinate.x), stoneY(coordinate.y), scale, 1-alpha)
		} else {
			drawImagePulse(screen, imgWhite, stoneX(coordinate.x), stoneY(coordinate.y), scale, 1-alpha)
		}
	}
}

func textColorBlack(alpha float64) color.RGBA {
	var textColor color.RGBA
	textColor.A = uint8(alpha * 255)
	return textColor
}

func textColorWhite(alpha float64) color.RGBA {
	var textColor color.RGBA
	textColor.R = 255
	textColor.G = 255
	textColor.B = 255
	textColor.A = uint8(alpha * 255)
	return textColor
}

func drawBlackMessage(screen *ebiten.Image, msg string, alpha float64) {
	textColor := textColorBlack(alpha)
	text.Draw(screen, msg, mplusNormalFont, columnBlack, row*2, textColor)
}

func drawBlackWin(screen *ebiten.Image, msg string, alpha float64) {
	if dogeMode == true {
		drawImagePulse(screen, imgDogeBig, 60, 1, 0.50, alpha)
		msg = "          Wins!"
	}
	textColor := textColorBlack(alpha)
	text.Draw(screen, msg, mplusBigFont, columnBlack, row*1+50, textColor)
}

func drawWhiteMessage(screen *ebiten.Image, msg string, alpha float64) {
	textColor := textColorWhite(alpha)
	text.Draw(screen, msg, mplusNormalFont, columnWhite, row*2, textColor)
}

func drawWhiteWin(screen *ebiten.Image, msg string, alpha float64) {
	if dogeMode == true {
		drawImagePulse(screen, imgCorgBig, float64(columnWhite)-125, 5, 1, alpha)
		msg = "        Wins!"
	}
	textColor := textColorWhite(alpha)
	text.Draw(screen, msg, mplusBigFont, columnWhite, row*1+50, textColor)
}

func drawMessage(screen *ebiten.Image, g *game, alpha float64) {
	if g.won == false {
		if g.gui.drawIntro == true {
			if g.player == false {
				if isPlayerHuman(g) == true || isPlayerHotseat(g) == true {
					text.Draw(screen, `Black to Move`, mplusNormalFont, columnBlack, row, color.Black)
					if dogeMode == true {
						drawImage(screen, imgDogeBig, 60, 0, 0.42)
					}
				} else {
					text.Draw(screen, `Black Thinking...`, mplusNormalFont, columnBlack, row, color.Black)
					if dogeMode == true {
						drawImage(screen, imgDogeBig, 60, 1, 0.42)
					}
				}
				drawBlackMessage(screen, g.gui.message, alpha)
			} else {
				if isPlayerHuman(g) == true || isPlayerHotseat(g) == true {
					text.Draw(screen, `White to Move`, mplusNormalFont, columnWhite, row, color.White)
					if dogeMode == true {
						drawImage(screen, imgCorgBig, float64(columnWhite)-125, 5, 1)
					}
				} else {
					text.Draw(screen, `White Thinking...`, mplusNormalFont, columnWhite, row, color.White)
					if dogeMode == true {
						drawImage(screen, imgCorgBig, float64(columnWhite)-125, 5, 1)
					}
				}
				drawWhiteMessage(screen, g.gui.message, alpha)
			}
		}
	} else {
		if g.gui.message == "Black Wins!" {
			drawBlackWin(screen, g.gui.message, alpha)
		} else {
			drawWhiteWin(screen, g.gui.message, alpha)
		}
	}
}

func drawPlayerID(screen *ebiten.Image, g *game, p ai, column int, color color.Color) {
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

func drawCaptured(screen *ebiten.Image, g *game, captured uint8, alpha float64, column int, color color.Color) {
	if capturedTen(g) == true {
		if g.player == true && column == columnBlack {
			color = textColorBlack(alpha)
		}
		if g.player == false && column == columnWhite {
			color = textColorWhite(alpha)
		}
	}
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

func drawPlayerText(screen *ebiten.Image, g *game, alpha float64, player bool) {
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
	drawPlayerID(screen, g, p, column, c)
	drawCaptured(screen, g, captured, alpha, column, c)
	drawTimer(screen, g, p, column, c)
}

func drawMove(screen *ebiten.Image, g *game) {
	if g.gui.drawIntro == true {
		text.Draw(screen, `move:`, mplusNormalFont, columnBlack, row*13, color.Black)
		text.Draw(screen, strconv.Itoa(int(g.move)), mplusNormalFont, columnBlack+160, row*13, color.Black)
	}
}

func drawText(screen *ebiten.Image, g *game, alpha float64) {
	drawMessage(screen, g, alpha)
	drawPlayerText(screen, g, alpha, false)
	drawPlayerText(screen, g, alpha, true)
	drawMove(screen, g)
}

func drawBluePulse(screen *ebiten.Image, g *game, alpha float64, blue *ebiten.Image) {
	drawImagePulse(screen, blue, stoneX(g.lastMove.x), stoneY(g.lastMove.y), scale, alpha*2)
}

func drawLastMove(screen *ebiten.Image, g *game, second, pulse, alpha float64) {
	if g.gui.drawLastMove == true && g.move > 0 {
		drawBluePulse(screen, g, alpha4(second, pulse), imgBlue)
		drawBluePulse(screen, g, alpha3(second, alpha), imgBlue2)
		drawBluePulse(screen, g, alpha2(second, pulse), imgBlue3)
		drawBluePulse(screen, g, alpha1(second, alpha), imgBlue4)
	}
}

func drawWinMove(screen *ebiten.Image, g *game, alpha float64) {
	if g.won == true && g.gui.drawWinMove == true {
		drawImagePulse(screen, imgRed, stoneX(g.winMove.x), stoneY(g.winMove.y), scale, alpha)
	}
}

func drawCapturedPosition(screen *ebiten.Image, g *game, position coordinate, alpha float64) {
	drawImagePulse(screen, imgCapture, stoneX(position.x), stoneY(position.y), scale, alpha)
}

func drawCapture(screen *ebiten.Image, g *game, alpha float64) {
	if g.gui.drawCapture == true {
		for _, position := range g.gui.capturedPositions {
			drawCapturedPosition(screen, g, position, alpha)
		}
	}
}

func drawUndo(screen *ebiten.Image, g *game) {
	if g.gui.undo == true && (isPlayerHuman(g) || isOpponentHuman(g)) && g.won == false {
		drawImage(screen, imgUndo, undoX, undoY, scale)
	}
}

func drawNewGame(screen *ebiten.Image, g *game) {
	drawImage(screen, imgNewGame, newGameX, newGameY, newGameScale)
}

func drawExit(screen *ebiten.Image, g *game) {
	drawImage(screen, imgExit, exitX, exitY, scale)
}

// draw draws the current game state
func draw(screen *ebiten.Image, g *game) {
	second, pulse, alpha := alphaTime()
	drawBackground(screen)
	if g.gui.newGame == true {
		drawNewGameOptions(screen, g, second, pulse, alpha)
	} else {
		drawGoban(screen, g)
		drawText(screen, g, alpha)
		if g.gui.drawIntro == false {
			drawIntro(screen)
		} else {
			drawStones(screen, g)
			drawHotseatSuggestion(screen, g, alpha)
			drawLastMove(screen, g, second, pulse, alpha)
			drawWinMove(screen, g, alpha)
			drawCapture(screen, g, alpha)
			drawTips(screen, g, alpha)
			drawUndo(screen, g)
		}
	}
	drawNewGame(screen, g)
	drawExit(screen, g)
}
