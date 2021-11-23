package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomonobold"
)

const (
	TPS          = 60
	QuitTime     = TPS * 60
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
		Size: 10,
	})
}

type Game struct {
	tps uint64
}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	g.tps++

	if g.tps >= QuitTime {
		return errors.New("end of profiling")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("soft kill")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Debug str
	str := fmt.Sprintf("TPS %.2f - FPS %.2f - Rand0 %.6f - Rand1 %.6f - Rand2 %.6f - Rand3 %.6f Rand4 %.6f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	)
	// Text drawing
	ebitenutil.DebugPrint(screen, str)
	//ebitenutil.DebugPrintWithTrianglesAt(screen, str, 0, 20)
	text.Draw(screen, str, face, 0, 50, color.RGBA{255, 255, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	f, err := os.Create("profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println("couldn't profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	ebiten.SetMaxTPS(TPS)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
