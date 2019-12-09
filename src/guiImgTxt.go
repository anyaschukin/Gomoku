package gomoku

import (
	"log"
	"time"
	
	"image/color"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

/// Images
var imgGoban *ebiten.Image
var imgBlackStone *ebiten.Image
var imgWhiteStone *ebiten.Image
var imgBlack *ebiten.Image
var imgWhite *ebiten.Image
var imgRed *ebiten.Image
var imgBlue *ebiten.Image
var imgBlue2 *ebiten.Image
var imgBlue3 *ebiten.Image
var imgBlue4 *ebiten.Image
var imgCapture *ebiten.Image
var imgNewGame *ebiten.Image
var imgExit *ebiten.Image
var imgSelect *ebiten.Image
var imgUndo *ebiten.Image
var imgCorg *ebiten.Image
var imgCorgBig *ebiten.Image
var imgDoge *ebiten.Image
var imgDogeBig *ebiten.Image
var imgColor *ebiten.Image
var imgBackground *ebiten.Image

var background color.Color

/// Text
var (
	mplusNormalFont font.Face
	mplusMediumFont font.Face
	mplusBigFont    font.Face
)

func init() {
	/// Initialize images
	var err error
	imgGoban, _, err = ebitenutil.NewImageFromFile("src/img/goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlackStone, _, err = ebitenutil.NewImageFromFile("src/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgWhiteStone, _, err = ebitenutil.NewImageFromFile("src/img/white.png", ebiten.FilterDefault)
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
	imgBlue2, _, err = ebitenutil.NewImageFromFile("src/img/blue2.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlue3, _, err = ebitenutil.NewImageFromFile("src/img/blue3.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBlue4, _, err = ebitenutil.NewImageFromFile("src/img/blue4.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgCapture, _, err = ebitenutil.NewImageFromFile("src/img/capture.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgNewGame, _, err = ebitenutil.NewImageFromFile("src/img/newGame.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgExit, _, err = ebitenutil.NewImageFromFile("src/img/exit.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgSelect, _, err = ebitenutil.NewImageFromFile("src/img/select.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgUndo, _, err = ebitenutil.NewImageFromFile("src/img/undo.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgCorg, _, err = ebitenutil.NewImageFromFile("src/img/back/corg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgDoge, _, err = ebitenutil.NewImageFromFile("src/img/back/doge.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgCorgBig, _, err = ebitenutil.NewImageFromFile("src/img/back/corgBig.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgDogeBig, _, err = ebitenutil.NewImageFromFile("src/img/back/dogeBig.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgBackground, _, err = ebitenutil.NewImageFromFile("src/img/back/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	imgColor, _, err = ebitenutil.NewImageFromFile("src/img/back/Color.png", ebiten.FilterDefault)
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
	mplusMediumFont = truetype.NewFace(tt, &truetype.Options{
		Size:    62,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	imgBlack = imgBlackStone
	imgWhite = imgWhiteStone
}

// every second - HotseatSuggestion, WinMove, & CapturedPosition
func alphaPulse(pulse float64) float64 {
	alpha := pulse
	if alpha > 1 {
		alpha = 2 - alpha
	}
	return alpha
}

// centralized time call
func alphaTime() (second, pulse, alpha float64) {
	second = float64(time.Now().Second() % 2)
	pulse = float64(time.Now().Nanosecond()) / 500000000
	alpha = alphaPulse(pulse)
	return
}

// alpha1 to 4 - Highlight LastMove rotates pulses
func alpha1(second, alpha float64) float64 {
	alpha = alpha * second
	return alpha
}

func alpha2(second, pulse float64) float64 {
	var alpha float64
	if second == 0 {
		if pulse > 1 {
			alpha = pulse - 1
		}
	} else {
		if pulse < 1 {
			alpha = 1 - pulse
		}
	}
	return alpha
}

func alpha3(second, alpha float64) float64 {
	alpha = alpha * (1 - second)
	return alpha
}

func alpha4(second, pulse float64) float64 {
	var alpha float64
	if second == 0 {
		if pulse < 1 {
			alpha = 1 - pulse
		}
	} else {
		if pulse > 1 {
			alpha = pulse - 1
		}
	}
	return alpha
}

// for drawIntro red pulse
func alphaDelay(second, pulse float64) float64 {
	var alpha float64
	if second == 0 {
		if pulse > 1 {
			alpha = pulse - 1
		}
	} else {
		alpha = 1 - (pulse - 2)
	}
	return alpha
}
