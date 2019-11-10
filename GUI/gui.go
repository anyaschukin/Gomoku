package gui

import (
	"fmt"
	"log"
	"os"

	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// type coordinate struct {
// 	y int8
// 	x int8
// }

// type position struct {
// 	occupied bool
// 	Player   bool
// }

// type align5 struct { //winning move for checking if opponent breaks it in the next move
// 	break5   bool
// 	capture8 bool // is it possible for the opponent to win by capturing 10? (have they already captured 8, and is there an available capture move)
// 	winner   bool /// rm?
// 	winmove  coordinate
// }

// type Game struct {
// 	Goban    [19][19]position
// 	Player   bool   // whose move is it? (Player 0 - black first)
// 	capture0 uint8  // capture 10 and win
// 	capture1 uint8  // capture 10 and win
// 	align5   align5 // one Player has aligned 5, however it can be broken. The other Player must break it, capture 10, or lose.
// 	// move		uint32				// how many moves have been played in total (is this desirable/necessary?)
// }

// func DumpGoban(Goban *[19][19]position) {
// 	fmt.Printf("     ")
// 	for x := 0; x < 19; x++ {
// 		fmt.Printf("{%2d         } ", x)
// 	}
// 	fmt.Printf("\n")
// 	for y := 0; y < 19; y++ {
// 		fmt.Printf("{%2d} ", y)
// 		for x := 0; x < 19; x++ {
// 			if Goban[y][x].occupied == true {
// 				fmt.Printf("\x1B[31m")
// 			}
// 			fmt.Printf("{%v\x1B[0m ", Goban[y][x].occupied)
// 			if Goban[y][x].occupied == true {
// 				fmt.Printf(" ")
// 			}
// 			color := ""
// 			if Goban[y][x].occupied == true {
// 				if Goban[y][x].Player == true {
// 					color = "\x1B[32m"
// 				} else {
// 					color = "\x1B[33m"
// 				}
// 			}
// 			fmt.Printf("%s%v\x1B[0m", color, Goban[y][x].Player)
// 			if Goban[y][x].Player == true {
// 				fmt.Printf(" ")
// 			}
// 			fmt.Printf("} ")
// 		}
// 		fmt.Printf("\n")
// 	}
// 	fmt.Printf("\n")
// }

// const (
// 	screenWidth  = 1315
// 	screenHeight = 1500
// )

var img_Goban *ebiten.Image
var img_black *ebiten.Image
var img_white *ebiten.Image

func init() {
	var err error
	img_Goban, _, err = ebitenutil.NewImageFromFile("GUI/img/Goban.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img_black, _, err = ebitenutil.NewImageFromFile("GUI/img/black.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img_white, _, err = ebitenutil.NewImageFromFile("GUI/img/white.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// func update(screen *ebiten.Image) error {
// 	if ebiten.IsDrawingSkipped() {
// 		return nil
// 	}

// 	return nil
// }

// func main() {
// 	if err := ebiten.Run(update, 640, 480, 1, "Render an image"); err != nil {
// 		log.Fatal(err)
// 	}
// }

func update(screen *ebiten.Image) error {
	// if err := GameLoop(); err != nil {//////////update Game state
	// 	return err
	// }
	if ebiten.IsDrawingSkipped() { /// do we want this (see cheat sheet)?
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) == true { // quit if press escape
		os.Exit(0) ////// rm, just for test. Return win message to GUI
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) == true {
		x, y := ebiten.CursorPosition()
		fmt.Printf("Mouse pressed x:%d y:%d\n", x, y)
	}
	op_Goban := &ebiten.DrawImageOptions{}
	op_Goban.GeoM.Translate(825, 20)
	op_Goban.GeoM.Scale(0.7, 0.7)

	op_stone := &ebiten.DrawImageOptions{}
	op_stone.GeoM.Translate(838, 34)
	op_stone.GeoM.Scale(0.7, 0.7)

	op_stone2 := &ebiten.DrawImageOptions{}
	op_stone2.GeoM.Translate((838 + 1045), 34)
	op_stone2.GeoM.Scale(0.7, 0.7)

	op_stone3 := &ebiten.DrawImageOptions{}
	op_stone3.GeoM.Translate(838, (34 + 1045))
	op_stone3.GeoM.Scale(0.7, 0.7)

	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(img_Goban, op_Goban)
	screen.DrawImage(img_black, op_stone)
	screen.DrawImage(img_black, op_stone2)
	screen.DrawImage(img_black, op_stone3)
	// ebitenutil.DebugPrint(screen, "Our first Game in Ebiten!") //////
	// Draw(screen) ////////// draw new image based on new Game state
	return nil
}

func RunEbiten(G *Game) {
	// if err := ebiten.Run(update, 1500, 1315, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
	DumpGoban(&G.Goban) //////
	w, h := ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	// ebiten.SetCursorVisible(true)//// helpful?
	if err := ebiten.Run(update, w, h, 1, "Gomoku"); err != nil {
		log.Fatal(err)
	}
	// if err := ebiten.Run(update, screenHeight, screenWidth, 1, "Gomoku"); err != nil {
	// 	log.Fatal(err)
	// }
}
