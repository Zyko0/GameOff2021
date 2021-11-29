package ui

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/Zyko0/GameOff2021/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenuView struct {
	active      bool
	needsRedraw bool

	lastIntentX int

	demoLevel *core.Core
	cache     *graphics.Cache
}

func NewMainMenuView() *MainMenuView {
	demoLevel := core.NewCore(assets.NoopSFXManager())
	demoLevel.Settings.BlockSettings.Heart = true
	demoLevel.Settings.BlockSettings.GoldenHeart = true
	demoLevel.Settings.BlockSettings.HigherSpawn = true
	demoLevel.Settings.BlockSettings.TallerBlocks = true
	demoLevel.Settings.BlockSettings.Harder = true
	demoLevel.Settings.BlockSettings.Harder2 = true
	demoLevel.Settings.EndWaveDistance = 9999999 // sorry for this magic number

	return &MainMenuView{
		active: true,

		lastIntentX: -1,

		demoLevel: demoLevel,
		cache:     graphics.NewCache(),
	}
}

func (mv *MainMenuView) Active() bool {
	return mv.active
}

func (mv *MainMenuView) Update() {
	// Leave on 'Enter' key
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		mv.active = false
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		assets.StopMainmenuMusic()
		return
	}
	// Play a dummy player in the background
	if rand.Intn(logic.TPS/4) == 0 {
		mv.lastIntentX *= -1
	}

	mv.demoLevel.Player.SetIntentX(float64(mv.lastIntentX))
	mv.demoLevel.Player.SetIntentJump(rand.Intn(logic.TPS) == 0)
	// Update level's logic
	mv.demoLevel.Update()
	mv.needsRedraw = true
	// Set graphic data
	mv.cache.BlockCount = len(mv.demoLevel.Blocks)
	mv.cache.BlockSeeds = mv.demoLevel.GetBlockSeeds()
	for i, b := range mv.demoLevel.Blocks {
		x, y, z := core.XYZToGraphics(b.GetX(), b.GetY(), b.GetZ())
		mv.cache.BlockPositions[i*3+0] = float32(x)
		mv.cache.BlockPositions[i*3+1] = float32(y)
		mv.cache.BlockPositions[i*3+2] = float32(z)
		mv.cache.BlockSizes[i*2+0] = float32(b.GetWidth())
		mv.cache.BlockSizes[i*2+1] = float32(b.GetHeight())
		mv.cache.BlockKinds[i] = float32(b.GetKind())
	}
}

func (mv *MainMenuView) Draw(screen *ebiten.Image) {
	// Draw a dummy level
	if mv.needsRedraw { // Save gpu resources if game is paused or if there has not been any update
		x, y, z := core.XYZToGraphics(mv.demoLevel.Player.GetX(), mv.demoLevel.Player.GetY(), mv.demoLevel.Player.GetZ())
		graphics.GetOffscreenImage().DrawRectShader(logic.GameSquareDim, logic.GameSquareDim, shaders.RaymarchingShader, &ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"ScreenSize":     []float32{float32(logic.GameSquareDim / graphics.GetQuality()), float32(logic.GameSquareDim / graphics.GetQuality())},
				"PlayerPosition": []float32{float32(x), float32(y), float32(z)},
				"PlayerRadius":   float32(mv.demoLevel.Player.GetRadius()),
				"Camera":         mv.demoLevel.Settings.CameraPosition,
				"Distance":       float32(mv.demoLevel.Wave.Distance),
				"DebugLines":     mv.demoLevel.Settings.DebugLines,

				"BlockCount":     float32(len(mv.demoLevel.Blocks)),
				"BlockPositions": mv.cache.BlockPositions,
				"BlockSizes":     mv.cache.BlockSizes,
				"BlockKinds":     mv.cache.BlockKinds,
				"BlockSeeds":     mv.cache.BlockSeeds,

				"PalettePlayer":       graphics.PalettePlayer,
				"PaletteBlock":        graphics.PaletteBlock,
				"PaletteRoad":         graphics.PaletteRoad,
				"PaletteBlockHarder":  graphics.PaletteBlockHarder,
				"PaletteBlockHarder2": graphics.PaletteBlockHarder2,
				"PaletteHeart":        graphics.PaletteHeart,
				"PaletteGoldenHeart":  graphics.PaletteGoldenHeart,
			},
		})
		// Let Update decide whenever there's a need for drawing the whole game scene again
		mv.needsRedraw = false
	}
	screen.DrawImage(graphics.GetOffscreenImage(), graphics.GetOffscreenOpts())
}
