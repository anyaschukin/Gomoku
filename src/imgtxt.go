package play //gui

import (
	"log"

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
