package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

const (
	TPS          = 60
	ScreenWidth  = 768
	ScreenHeight = 576
)

var (
	face font.Face
)

func init() {
	rand.Seed(time.Now().UnixNano())

	font, err := truetype.Parse(gomonobold.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face = truetype.NewFace(font, &truetype.Options{
		Size: 50,
	})
}

type Game struct {
	tps uint64
}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("soft kill")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		symbol = "*^:°"
	)
	//ebitenutil.DebugPrintWithTrianglesAt(screen, str, 0, 20)
	for _, r := range symbol {
		text.Draw(screen, string(r), face, 0, 50, color.RGBA{255, 255, 255, 255})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetMaxTPS(TPS)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
