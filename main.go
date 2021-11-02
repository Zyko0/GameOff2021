package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 480
	screenHeight = 480
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetMaxTPS(60)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)

	g := &Game{}
	ebiten.RunGame(g)
}
