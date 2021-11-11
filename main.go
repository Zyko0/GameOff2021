package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Zyko0/GameOff2021/core"
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/Zyko0/GameOff2021/shaders"
	"github.com/Zyko0/GameOff2021/ui"
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

	pauseView    *ui.PauseView
	gameOverView *ui.GameoverView
	augmentView  *ui.AugmentView
	hud          *ui.HUD

	level          *core.Level
	augmentManager *core.AugmentManager

	offscreen *ebiten.Image
	cache     *graphics.Cache

	needsRedraw bool
}

func New() *Game {
	level := core.NewLevel()
	return &Game{
		paused:      false,
		needsRedraw: false,

		pauseView:    ui.NewPauseView(),
		gameOverView: ui.NewGameoverView(),
		augmentView:  ui.NewAugmentView(),
		hud:          ui.NewHUD(level.PlayerHP, nil),

		level:          level,
		augmentManager: core.NewAugmentManager(),

		offscreen: ebiten.NewImage(logic.GameSquareDim, logic.GameSquareDim),
		cache:     graphics.NewCache(),
	}
}

func (g *Game) Update() error {
	g.needsRedraw = false
	// Handle game reset first
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		rand.Seed(time.Now().UnixNano())
		g.level = core.NewLevel()
		g.pauseView.Reset()
		g.augmentView.Reset()
	}
	// Gameover view having checked for a restart
	g.gameOverView.Update(g.level.PlayerHP, g.level.Settings.HpToGameOver)
	if g.gameOverView.Active() {
		return nil
	}
	// Pause view
	g.pauseView.Update()
	if g.pauseView.Active() {
		return nil
	}
	// Augments management => if eligible for an augment, show view
	// TODO: first at 1000 ?
	if g.level.GetTicks()%g.level.Settings.AugmentsTicksInterval == 0 {
		// If needs an augment selection but the view is not active yet, roll, activate and abort
		if !g.augmentView.Active() {
			rolls := g.augmentManager.RollAugments()
			g.augmentView.SetAugments(rolls)
			return nil
		}
		// Update and check for user input, if an augment is picked, the view isn't active anymore
		g.augmentView.Update()
		if g.augmentView.Active() {
			// Abort because there's still an augment to pick
			return nil
		}
		// If the view is not active anymore, check for selection
		a := g.augmentView.Augments[g.augmentView.SelectedIndex]
		g.augmentManager.AddAugment(a)
		g.level.Settings.ApplyAugments(g.augmentManager.CurrentAugments)
		if a.Cost.Kind == augments.CostKindHP {
			g.level.PlayerHP -= a.Cost.Value
		}
	}
	// Require a draw
	g.needsRedraw = true
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
	// Game update
	g.level.Update()
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
	// Update HUD
	g.hud.Update(g.level.PlayerHP, g.augmentManager.CurrentAugments)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Offscreen intermediate draw
	if g.needsRedraw { // Save gpu resources if game is paused
		x, y, z := core.XYZToGraphics(g.level.Player.GetX(), g.level.Player.GetY(), g.level.Player.GetZ())
		g.offscreen.DrawRectShader(logic.GameSquareDim, logic.GameSquareDim, shaders.RaymarchShader, &ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"ScreenSize":     []float32{float32(logic.GameSquareDim), float32(logic.GameSquareDim)},
				"PlayerPosition": []float32{float32(x), float32(y), float32(z)},
				"PlayerRadius":   float32(g.level.Player.GetRadius()),
				"Camera":         g.level.Settings.CameraPosition,
				"Distance":       float32(g.level.Distance),

				"BlockCount":     float32(len(g.level.Blocks)),
				"BlockPositions": g.cache.BlockPositions,
				"BlockSizes":     g.cache.BlockSizes,

				"Palette0": graphics.PlayerPalette,
				"Palette1": graphics.BlockPalette,
				"Palette2": graphics.RoadPalette,
			},
		})
		// Draw HUD on offscreen
		g.hud.Draw(g.offscreen)
		// Let Update() decide whenever there's a need for drawing the whole scene again
		g.needsRedraw = false
	}
	// Draw buffer to screen
	screen.DrawImage(g.offscreen, &ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	})
	// Gameover view
	if g.gameOverView.Active() {
		g.gameOverView.Draw(screen)
		return
	}
	// Pause view
	if g.pauseView.Active() {
		g.pauseView.Draw(screen)
		return
	}
	// Augment view
	if g.augmentView.Active() {
		g.augmentView.Draw(screen)
		return
	}
	// Debug
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS %.2f - FPS %.2f - Tick %d - BlockCount %d - Score %d - Speed %.2f - HP %d - Distance %.2f",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			g.level.GetTicks(),
			len(g.level.Blocks),
			g.level.GetScore(),
			g.level.GetSpeed(),
			g.level.PlayerHP,
			g.level.Distance,
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	// TODO: set vsync on
	// Note: setTimeout is called when FPSMoveVsyncOffMaximum which might create lag
	// ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
