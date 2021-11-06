package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Zyko0/GameOff2021/drafts/01/automata"
	"github.com/Zyko0/GameOff2021/drafts/01/graphics"
	"github.com/Zyko0/GameOff2021/drafts/01/levels"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 1024
	screenHeight = 768

	boardWidth  = 480
	boardHeight = 480
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	automata *automata.Automata
	graphics *graphics.Graphics
}

func NewRandomAutomata() *automata.Automata {
	const gridcount = 13

	walls := make([]*automata.Wall, 4)
	walls[automata.WallPositionLeft] = automata.NewWall(automata.WallPositionLeft, automata.WallEffectReverseDirection)
	walls[automata.WallPositionRight] = automata.NewWall(automata.WallPositionRight, automata.WallEffectReverseDirection)
	walls[automata.WallPositionTop] = automata.NewWall(automata.WallPositionTop, automata.WallEffectReverseDirection)
	walls[automata.WallPositionBottom] = automata.NewWall(automata.WallPositionBottom, automata.WallEffectReverseDirection)

	particles := levels.GetPlusShapeParticles(gridcount)

	return automata.New(gridcount, walls, particles)
}

func (g *Game) Update() error {
	// Quit
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("soft kill")
	}
	// Instanciate a new random automata
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.automata = NewRandomAutomata()
		g.graphics.SetPalette(graphics.NewRandomPalette())
	}
	// Advance step
	if kpd := inpututil.KeyPressDuration(ebiten.KeyRight); kpd == 1 || kpd > 10 && kpd%2 == 0 {
		g.automata.Advance()
	}
	// Rewind step
	if kpd := inpututil.KeyPressDuration(ebiten.KeyLeft); kpd == 1 || kpd > 10 && kpd%2 == 0 {
		g.automata.Rewind()
	}
	// Debug
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cellSize := float64(boardWidth) / float64(g.automata.GetGridCount())
	g.graphics.DrawParticles(screen, float32(cellSize), g.automata.GetParticlesAtStep())
	// Debug
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS %.2f - FPS - %.2f - Step %d - Particles %d",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			g.automata.GetCurrentStep(),
			len(g.automata.GetParticlesAtStep()),
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	f, err := os.Create("beat.prof")
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

	ebiten.SetMaxTPS(60)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetFullscreen(true)

	g := &Game{
		automata: NewRandomAutomata(),
		graphics: graphics.New(),
	}
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("rungame:", err)
	}
}
