package gomoku

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

func drawG(screen *ebiten.Image, alpha float64, img1 *ebiten.Image, img2 *ebiten.Image) {
	drawImagePulse(screen, img1, stoneX(1), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(1), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(1), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(1), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(1), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(1), stoneY(11), scale, alpha)
	drawImagePulse(screen, img1, stoneX(1), stoneY(12), scale, alpha)

	drawImagePulse(screen, img1, stoneX(2), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(2), stoneY(12), scale, alpha)
	drawImagePulse(screen, img1, stoneX(2), stoneY(13), scale, alpha)

	drawImagePulse(screen, img1, stoneX(3), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(3), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(3), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(7), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(8), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(3), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(3), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(3), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(4), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(4), stoneY(5), scale, alpha)
	drawImagePulse(screen, img1, stoneX(4), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(4), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(4), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(4), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(5), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(5), stoneY(5), scale, alpha)
	drawImagePulse(screen, img1, stoneX(5), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(5), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(5), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(5), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(5), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(5), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(5), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(6), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(6), stoneY(5), scale, alpha)
	drawImagePulse(screen, img1, stoneX(6), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(6), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(6), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(6), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(6), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(6), stoneY(13), scale, alpha)

	drawImagePulse(screen, img1, stoneX(7), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(7), stoneY(5), scale, alpha)
	drawImagePulse(screen, img1, stoneX(7), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(7), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(7), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(7), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(7), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(7), stoneY(12), scale, alpha)
	drawImagePulse(screen, img1, stoneX(7), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(7), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(8), stoneY(5), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(8), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(11), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(12), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(8), stoneY(14), scale, alpha)
}

func drawO(screen *ebiten.Image, alpha float64, img1 *ebiten.Image, img2 *ebiten.Image) {
	drawImagePulse(screen, img1, stoneX(10), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(7), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(8), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(11), scale, alpha)
	drawImagePulse(screen, img1, stoneX(10), stoneY(12), scale, alpha)

	drawImagePulse(screen, img1, stoneX(11), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(11), stoneY(12), scale, alpha)
	drawImagePulse(screen, img1, stoneX(11), stoneY(13), scale, alpha)

	drawImagePulse(screen, img1, stoneX(12), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(12), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(12), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(13), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(13), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(13), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(13), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(14), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(14), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(14), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(14), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(15), stoneY(4), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(12), scale, alpha)
	drawImagePulse(screen, img2, stoneX(15), stoneY(13), scale, alpha)
	drawImagePulse(screen, img1, stoneX(15), stoneY(14), scale, alpha)

	drawImagePulse(screen, img1, stoneX(16), stoneY(5), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(6), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(7), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(8), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(9), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(10), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(11), scale, alpha)
	drawImagePulse(screen, img2, stoneX(16), stoneY(12), scale, alpha)
	drawImagePulse(screen, img1, stoneX(16), stoneY(13), scale, alpha)

	drawImagePulse(screen, img1, stoneX(17), stoneY(6), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(7), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(8), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(9), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(10), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(11), scale, alpha)
	drawImagePulse(screen, img1, stoneX(17), stoneY(12), scale, alpha)
}

// drawIntro draws "GO" in stones on the goban at the start of a new game
func drawIntro(screen *ebiten.Image) {
	if g.gui.introTime.IsZero() {
		g.gui.introTime = time.Now()
	}
	elapsed := time.Since(g.gui.introTime)
	second := float64(elapsed.Truncate(time.Second) / 1000000000)
	pulse := float64(elapsed) / 500000000
	alpha := alphaPulse(pulse)
	alpha2 := alphaDelay(second, pulse)

	drawG(screen, alpha, imgWhite, imgBlack)
	drawO(screen, alpha, imgWhite, imgBlack)
	drawG(screen, alpha, imgBlue, imgBlue3)
	drawO(screen, alpha, imgBlue, imgBlue3)

	drawG(screen, alpha2, imgRed, imgRed)
	drawO(screen, alpha2, imgRed, imgRed)

	if elapsed > 1500000000 {
		g.gui.drawIntro = true
	}
}
