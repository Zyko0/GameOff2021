package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core"
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/Zyko0/GameOff2021/shaders"
	"github.com/Zyko0/GameOff2021/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	paused bool

	splashView   *ui.SplashView
	mainMenuView *ui.MainMenuView
	pauseView    *ui.PauseView
	gameOverView *ui.GameoverView
	augmentView  *ui.AugmentView
	hud          *ui.HUD

	core           *core.Core
	augmentManager *augments.Manager

	cache *graphics.Cache

	needsRedraw bool
}

func New() *Game {
	level := core.NewCore(assets.DefaultSFXManager())
	return &Game{
		paused:      false,
		needsRedraw: false,

		splashView:   ui.NewSplashView(),
		mainMenuView: ui.NewMainMenuView(),
		pauseView:    ui.NewPauseView(),
		gameOverView: ui.NewGameoverView(),
		augmentView:  ui.NewAugmentView(),
		hud:          ui.NewHUD(level.PlayerHP),

		core:           level,
		augmentManager: augments.NewManager(),

		cache: graphics.NewCache(),
	}
}

func (g *Game) Update() error {
	g.needsRedraw = false
	// Check for any user input changing quality
	if graphics.UpdateQualitySettings() {
		g.needsRedraw = true
	}
	// Splash screen view
	if g.splashView.Active() {
		g.splashView.Update()
		return nil
	}
	// Main menu view with background demo
	if g.mainMenuView.Active() {
		g.mainMenuView.Update()
		return nil
	}
	// Handle game reset first
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		rand.Seed(time.Now().UnixNano())
		g.core = core.NewCore(assets.DefaultSFXManager())
		g.pauseView.Reset()
		g.augmentView.Reset()
		g.augmentManager.Reset()
		g.cache.Reset()
		// TODO: Not sure we want to rewind this audio player is a spam "R" is going on
		assets.ReplayInGameMusic()
	}
	// Gameover view having checked for a restart
	g.gameOverView.Update(g.core.PlayerHP, g.core.Settings.HpToGameOver)
	if g.gameOverView.Active() {
		return nil
	}
	// Pause view
	g.pauseView.Update()
	if g.pauseView.Active() {
		return nil
	}
	// Augments management => if eligible for an augment, show view
	if g.core.IsWaveOver() {
		// If needs an augment selection but the view is not active yet, roll, activate and abort
		if !g.augmentView.Active() {
			rolls := g.augmentManager.RollAugments(g.core.Wave.Number)
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
		// Otherwise pick augment
		g.augmentManager.AddAugment(a)
		g.core.Settings.ApplyAugments(g.augmentManager.CurrentAugments)
		// Start next wave
		g.core.StartNextWave()
		g.augmentView.Reset()
	}
	// Require a draw
	g.needsRedraw = true
	// Reset cache
	g.cache.Reset()
	// Reset player's moving intents
	g.core.Player.SetIntentX(0)
	g.core.Player.SetIntentJump(false)
	// Jump
	if kpd := inpututil.KeyPressDuration(ebiten.KeySpace); kpd > 0 {
		g.core.Player.SetIntentJump(true)
	}
	// Move player right
	if kpd := inpututil.KeyPressDuration(ebiten.KeyRight); (!g.core.Settings.PerfectStep && kpd > 0) || kpd == 1 {
		g.core.Player.SetIntentX(1)
	}
	// Move player left
	if kpd := inpututil.KeyPressDuration(ebiten.KeyLeft); (!g.core.Settings.PerfectStep && kpd > 0) || kpd == 1 {
		g.core.Player.SetIntentX(-1)
	}
	// Game update
	g.core.Update()
	assets.ResumeInGameMusic()
	// Set graphic data
	g.cache.BlockCount = len(g.core.Blocks)
	g.cache.BlockSeeds = g.core.GetBlockSeeds()
	for i, b := range g.core.Blocks {
		x, y, z := core.XYZToGraphics(b.GetX(), b.GetY(), b.GetZ())
		g.cache.BlockPositions[i*3+0] = float32(x)
		g.cache.BlockPositions[i*3+1] = float32(y)
		g.cache.BlockPositions[i*3+2] = float32(z)
		g.cache.BlockSizes[i*2+0] = float32(b.GetWidth())
		g.cache.BlockSizes[i*2+1] = float32(b.GetHeight())
		g.cache.BlockKinds[i] = float32(b.GetKind())
	}
	// Update HUD
	g.hud.Update(g.core.PlayerHP)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Splash screen view
	if g.splashView.Active() {
		g.splashView.Draw(screen)
		return
	}
	// Main menu view with background demo
	if g.mainMenuView.Active() {
		g.mainMenuView.Draw(screen)
		return
	}
	// Offscreen intermediate draw
	if g.needsRedraw { // Save gpu resources if game is paused or if there has not been any update
		x, y, z := core.XYZToGraphics(g.core.Player.GetX(), g.core.Player.GetY(), g.core.Player.GetZ())
		graphics.GetOffscreenImage().DrawRectShader(logic.GameSquareDim, logic.GameSquareDim, shaders.RaymarchingShader, &ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"ScreenSize":     []float32{float32(logic.GameSquareDim / graphics.GetQuality()), float32(logic.GameSquareDim / graphics.GetQuality())},
				"PlayerPosition": []float32{float32(x), float32(y), float32(z)},
				"PlayerRadius":   float32(g.core.Player.GetRadius()),
				"Camera":         g.core.Settings.CameraPosition,
				"Distance":       float32(g.core.Wave.Distance),
				"DebugLines":     g.core.Settings.DebugLines,

				"BlockCount":     float32(len(g.core.Blocks)),
				"BlockPositions": g.cache.BlockPositions,
				"BlockSizes":     g.cache.BlockSizes,
				"BlockKinds":     g.cache.BlockKinds,
				"BlockSeeds":     g.cache.BlockSeeds,

				"PalettePlayer":       graphics.PalettePlayer,
				"PaletteBlock":        graphics.PaletteBlock,
				"PaletteRoad":         graphics.PaletteRoad,
				"PaletteBlockHarder":  graphics.PaletteBlockHarder,
				"PaletteBlockHarder2": graphics.PaletteBlockHarder2,
				"PaletteHeart":        graphics.PaletteHeart,
				"PaletteGoldenHeart":  graphics.PaletteGoldenHeart,
			},
		})
		// Draw HUD on offscreen
		g.hud.Draw(graphics.GetOffscreenImage())
		// Let Update decide whenever there's a need for drawing the whole game scene again
		g.needsRedraw = false
	}
	// Draw buffer to screen
	screen.DrawImage(graphics.GetOffscreenImage(), graphics.GetOffscreenOpts())
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
	// TODO: this debug took 17sec out of 120sec which is pretty huge
	str := fmt.Sprintf("TPS %.2f - FPS %.2f - Wave %d - BlockCount %d - Score %d - Speed %.2f - HP %d - Distance %.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		g.core.Wave.Number,
		len(g.core.Blocks),
		g.core.GetScore(),
		g.core.GetSpeed(),
		g.core.PlayerHP,
		g.core.Wave.Distance,
	)
	text.Draw(screen, str, assets.CardBodyDescriptionTextFontFace, 0, 20, color.RGBA{255, 255, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	// Note: setTimeout is called when FPSMoveVsyncOffMaximum which might create lag
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)

	assets.PlayMainmenuMusic()

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
