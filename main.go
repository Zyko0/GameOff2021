package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Zyko0/GameOff2021/core"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/Zyko0/GameOff2021/shaders"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	geom = ebiten.GeoM{}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	geom.Translate(float64(logic.ScreenWidth-logic.GameSquareDim)/2, 0)
}

type Game struct {
	paused bool

	level *core.Level

	offscreen *ebiten.Image
	cache     *graphics.Cache

	needsRedraw bool
}

func New() *Game {
	return &Game{
		paused:      false,
		needsRedraw: false,

		level: core.NewLevel(),

		offscreen: ebiten.NewImage(logic.GameSquareDim, logic.GameSquareDim),
		cache:     graphics.NewCache(),
	}
}

func (g *Game) Update() error {
	g.needsRedraw = true
	// Pause
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.paused = !g.paused
	}
	if g.paused {
		return nil
	}
	// Reset cache
	g.cache.Reset()
	// Reset player's moving intents
	g.level.Player.SetIntentX(0)
	g.level.Player.SetIntentAction(false)
	// Quit
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// TODO: Don't forget to remove this
		return errors.New("soft kill")
	}
	// Restart
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		// Reset to a new level for now
		g.level = core.NewLevel()
		rand.Seed(time.Now().UnixNano())
	}
	// Jump
	// TODO:
	if kpd := inpututil.KeyPressDuration(ebiten.KeySpace); kpd > 0 {
		g.level.Player.SetIntentAction(true)
	}
	// Move player right
	if kpd := inpututil.KeyPressDuration(ebiten.KeyRight); kpd > 0 {
		g.level.Player.SetIntentX(1)
	}
	// Move player left
	if kpd := inpututil.KeyPressDuration(ebiten.KeyLeft); kpd > 0 {
		g.level.Player.SetIntentX(-1)
	}

	// Check game over before updating
	if g.level.GetPlayerHP() > 0 {
		// Update level
		g.level.Update()
	}
	// Set graphic data
	g.cache.BlockCount = len(g.level.Blocks)
	for i, b := range g.level.Blocks {
		x, y, z := core.XYZToGraphics(b.GetX(), b.GetY(), b.GetZ())
		g.cache.BlockPositions[i*3+0] = float32(x)
		g.cache.BlockPositions[i*3+1] = float32(y)
		g.cache.BlockPositions[i*3+2] = float32(z)
		g.cache.BlockSizes[i*2+0] = float32(b.GetWidth())
		g.cache.BlockSizes[i*2+1] = float32(b.GetHeight())
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Buffer intermediate draw
	if !g.paused && g.needsRedraw { // Save gpu resources if game is paused
		x, y, z := core.XYZToGraphics(g.level.Player.GetX(), g.level.Player.GetY(), g.level.Player.GetZ())
		g.offscreen.DrawRectShader(logic.GameSquareDim, logic.GameSquareDim, shaders.RaymarchShader, &ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"ScreenSize":     []float32{float32(logic.GameSquareDim), float32(logic.GameSquareDim)},
				"PlayerPosition": []float32{float32(x), float32(y), float32(z)},
				"PlayerRadius":   float32(g.level.Player.GetRadius()),
				"Camera":         g.level.Settings.ActualSettings.CameraPosition,
				"Distance":       float32(g.level.Distance),

				"BlockCount":     float32(len(g.level.Blocks)),
				"BlockPositions": g.cache.BlockPositions,
				"BlockSizes":     g.cache.BlockSizes,

				"Palette0": graphics.PlayerPalette,
				"Palette1": graphics.BlockPalette,
				"Palette2": graphics.RoadPalette,
			},
		})
		// Let Update() decide whenever there's a need for drawing the whole scene again
		g.needsRedraw = false
	}
	// Draw buffer to screen
	screen.DrawImage(g.offscreen, &ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterLinear,
	})
	// Debug
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS %.2f - FPS %.2f - BlockCount %d - Score %d - Speed %.2f - HP %d",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			len(g.level.Blocks),
			g.level.GetScore(),
			g.level.GetSpeed(),
			g.level.GetPlayerHP(),
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
