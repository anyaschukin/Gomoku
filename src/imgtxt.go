package play //gui

import (
	"log"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

/// Images
var imgGoban *ebiten.Image
var imgBlack *ebiten.Image
var imgWhite *ebiten.Image
var imgRed *ebiten.Image
var imgBlue *ebiten.Image
var imgBlue2 *ebiten.Image
var imgBlue3 *ebiten.Image
var imgBlue4 *ebiten.Image
var imgExit *ebiten.Image
var imgNewGame *ebiten.Image
var imgSelect *ebiten.Image

/// Text
var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

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

func alphaPulse() float64 {
	alpha := float64(time.Now().Nanosecond()) / 500000000 //% 1 /////
	if alpha > 1 {
		alpha = 2 - alpha
	}
	return alpha
}

func alpha1() float64 {
	second := float64(time.Now().Second() % 2)
	alpha := alphaPulse() * second
	return alpha
}

func alpha2() float64 {
	second := float64(time.Now().Second() % 2)
	nano := time.Now().Nanosecond()
	var alpha = 0.0
	if second == 0 {
		if nano > 500000000 {
			alpha = (float64(time.Now().Nanosecond()) / 500000000) - 1
		}
	} else {
		if nano < 500000000 {
			alpha = 1 - (float64(time.Now().Nanosecond()) / 500000000)
		}
	}
	return alpha
}

func alpha3() float64 {
	second := float64(time.Now().Second() % 2)
	alpha := alphaPulse() * (1 - second)
	return alpha
}

func alpha4() float64 {
	second := float64(time.Now().Second() % 2)
	nano := time.Now().Nanosecond()
	var alpha = 0.0
	if second == 0 {
		if nano < 500000000 {
			alpha = 1 - (float64(time.Now().Nanosecond()) / 500000000)
		}
	} else {
		if nano > 500000000 {
			alpha = (float64(time.Now().Nanosecond()) / 500000000) - 1
		}
	}
	return alpha
}
