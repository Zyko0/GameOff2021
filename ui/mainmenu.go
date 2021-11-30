package ui

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/Zyko0/GameOff2021/shaders"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// TODO: lots of magic numbers and ugly file, because rushed last, need to fix

const (
	mainmenuCardLogicalOffsetX = (logic.ScreenWidth - logic.GameSquareDim) / 2

	mainmenuCardOffsetX = 96
	mainmenuCardOffsetY = 96
	mainmenuCardWidth   = logic.GameSquareDim - mainmenuCardOffsetX*2
	mainmenuCardHeight  = logic.GameSquareDim - mainmenuCardOffsetY*2

	mainmenuButtonOffsetX = 16
	mainmenuButtonOffsetY = mainmenuCardHeight - 64
	mainmenuButtonWidth   = mainmenuCardWidth - mainmenuButtonOffsetX*2
	mainmenuButtonHeight  = mainmenuCardHeight - mainmenuButtonOffsetY - 16
)

type MainMenuView struct {
	active      bool
	needsRedraw bool

	lastIntentX int

	demoLevel *core.Core
	cache     *graphics.Cache

	card *ebiten.Image
}

func NewMainMenuView() *MainMenuView {
	demoLevel := core.NewCore(assets.NoopSFXManager())
	demoLevel.Settings.BlockSettings.Heart = true
	demoLevel.Settings.BlockSettings.GoldenHeart = true
	demoLevel.Settings.BlockSettings.HigherSpawn = true
	demoLevel.Settings.BlockSettings.TallerBlocks = true
	demoLevel.Settings.BlockSettings.Harder = true
	demoLevel.Settings.BlockSettings.Harder2 = true
	demoLevel.Settings.EndWaveDistance = 99999999 // sorry for this magic number

	card := ebiten.NewImage(mainmenuCardWidth, mainmenuCardHeight)
	// Background and border
	graphics.DrawRect(card, 0, 0, mainmenuCardWidth, mainmenuCardHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(card, 0, 0, mainmenuCardWidth, mainmenuCardHeight, 1, []float32{1, 1, 1, 0.9})
	// Game title
	titleStr := "GameOff 2021"
	titleStrBounds := text.BoundString(assets.CardTitleFontFace, titleStr)
	text.Draw(
		card,
		titleStr,
		assets.CardTitleFontFace,
		card.Bounds().Dx()/2-titleStrBounds.Dx()/2,
		titleStrBounds.Dy()+32,
		color.White,
	)
	// Labels
	labelQualityStr := "Quality (Key 1/2/3/4) :"
	labelQualityStrBounds := text.BoundString(assets.CardBodyTextFontFace, labelQualityStr)
	text.Draw(
		card,
		labelQualityStr,
		assets.CardBodyTextFontFace,
		32,
		mainmenuCardOffsetY+16,
		color.White,
	)
	labelPerfectStepStr := "Perfect Step (Key S)  :"
	labelPerfectStepStrBounds := text.BoundString(assets.CardBodyTextFontFace, labelPerfectStepStr)
	text.Draw(
		card,
		labelPerfectStepStr,
		assets.CardBodyTextFontFace,
		32,
		mainmenuCardOffsetY+labelQualityStrBounds.Dy()+32,
		color.White,
	)
	labelDebugLinesStr := "Debug Lines (Key D)   :"
	labelDebugLinesStrBounds := text.BoundString(assets.CardBodyTextFontFace, labelDebugLinesStr)
	text.Draw(
		card,
		labelDebugLinesStr,
		assets.CardBodyTextFontFace,
		32,
		mainmenuCardOffsetY+labelQualityStrBounds.Dy()+labelPerfectStepStrBounds.Dy()+48,
		color.White,
	)
	labelTPSStr := "Ticks per second (60~):"
	labelTPSStrBounds := text.BoundString(assets.CardBodyTextFontFace, labelTPSStr)
	text.Draw(
		card,
		labelTPSStr,
		assets.CardBodyTextFontFace,
		32,
		mainmenuCardOffsetY+labelQualityStrBounds.Dy()+labelPerfectStepStrBounds.Dy()+labelDebugLinesStrBounds.Dy()+64,
		color.White,
	)
	labelFPSStr := "Frames per second     :"
	text.Draw(
		card,
		labelFPSStr,
		assets.CardBodyTextFontFace,
		32,
		mainmenuCardOffsetY+labelQualityStrBounds.Dy()+labelPerfectStepStrBounds.Dy()+labelDebugLinesStrBounds.Dy()+labelTPSStrBounds.Dy()+80,
		color.White,
	)
	// Draw start button
	graphics.DrawRect(card, mainmenuButtonOffsetX, mainmenuButtonOffsetY, mainmenuButtonWidth, mainmenuButtonHeight, []float32{0.4, 0.0, 0.7, 0.9})
	graphics.DrawRectBorder(card, mainmenuButtonOffsetX, mainmenuButtonOffsetY, mainmenuButtonWidth, mainmenuButtonHeight, 1, []float32{1, 1, 1, 0.9})
	// Start button text
	startStr := "Start"
	startStrBounds := text.BoundString(assets.CardBodyTitleFontFace, startStr)
	text.Draw(
		card,
		startStr,
		assets.CardBodyTitleFontFace,
		card.Bounds().Dx()/2-startStrBounds.Dx()/2,
		mainmenuButtonOffsetY+mainmenuButtonHeight/2+startStrBounds.Dy()/2,
		color.White,
	)

	return &MainMenuView{
		active: true,

		lastIntentX: -1,

		demoLevel: demoLevel,
		cache:     graphics.NewCache(),

		card: card,
	}
}

func (mv *MainMenuView) Active() bool {
	return mv.active
}

func (mv *MainMenuView) Update() {
	core.UpdateGlobalSettings()
	// Leave on 'R' key or button click
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		mv.active = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x0 := mainmenuCardLogicalOffsetX + mainmenuCardOffsetX + mainmenuButtonOffsetX
		x1 := x0 + mainmenuButtonWidth
		y0 := mainmenuCardOffsetY + mainmenuButtonOffsetY
		y1 := y0 + mainmenuButtonHeight
		if x >= x0 && x <= x1 && y >= y0 && y <= y1 {
			mv.active = false
		}
	}
	if !mv.active {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		assets.StopMainmenuMusic()
		return
	}
	// Play a dummy player in the background
	if core.PerfectStep() {
		mv.lastIntentX = 0
		if rand.Intn(logic.TPS/4) == 0 {
			mv.lastIntentX = 1
			if rand.Intn(2) == 0 {
				mv.lastIntentX = -1
			}
		}
	} else if rand.Intn(logic.TPS/4) == 0 {
		if mv.lastIntentX == 0 {
			mv.lastIntentX = 1
		}
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
				"DebugLines":     core.DebugLines(),

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
	// Main menu card
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(mainmenuCardLogicalOffsetX+mainmenuCardOffsetX),
		float64(mainmenuCardOffsetY),
	)
	screen.DrawImage(mv.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// States
	mv.drawLabels(screen)
}

func (mv *MainMenuView) drawLabels(screen *ebiten.Image) {
	const labelOffsetY = mainmenuCardOffsetY*2 + 16

	strBoundsY := text.BoundString(assets.CardBodyTextFontFace, "TexT").Dy()
	// Quality
	var qualityStr string
	var qualityColor color.Color

	switch graphics.GetQuality() {
	case graphics.QualityVeryLow:
		qualityStr = "Very low"
		qualityColor = color.RGBA{210, 0, 0, 255}
	case graphics.QualityLow:
		qualityStr = "Low"
		qualityColor = color.RGBA{210, 75, 0, 255}
	case graphics.QualityMedium:
		qualityStr = "Medium"
		qualityColor = color.RGBA{210, 150, 0, 255}
	case graphics.QualityHigh:
		qualityStr = "High"
		qualityColor = color.RGBA{210, 210, 0, 255}
	}
	text.Draw(
		screen,
		qualityStr,
		assets.CardBodyTextFontFace,
		mainmenuCardLogicalOffsetX+mainmenuCardOffsetX+216,
		labelOffsetY,
		qualityColor,
	)
	// Perfect Step
	var perfectStepStr string
	var perfectStepColor color.Color

	if core.PerfectStep() {
		perfectStepStr = "ON"
		perfectStepColor = color.RGBA{0, 210, 0, 255}
	} else {
		perfectStepStr = "OFF"
		perfectStepColor = color.RGBA{210, 0, 0, 255}
	}
	text.Draw(
		screen,
		perfectStepStr,
		assets.CardBodyTextFontFace,
		mainmenuCardLogicalOffsetX+mainmenuCardOffsetX+216,
		labelOffsetY+strBoundsY+20,
		perfectStepColor,
	)
	// Debug Lines
	var debugLinesStr string
	var debugLinesColor color.Color

	if core.DebugLines() > 0 {
		debugLinesStr = "ON"
		debugLinesColor = color.RGBA{0, 210, 0, 255}
	} else {
		debugLinesStr = "OFF"
		debugLinesColor = color.RGBA{210, 0, 0, 255}
	}
	text.Draw(
		screen,
		debugLinesStr,
		assets.CardBodyTextFontFace,
		mainmenuCardLogicalOffsetX+mainmenuCardOffsetX+216,
		labelOffsetY+strBoundsY*2+38,
		debugLinesColor,
	)
	// TPS
	text.Draw(
		screen,
		fmt.Sprintf("%.2f", ebiten.CurrentTPS()),
		assets.CardBodyTextFontFace,
		mainmenuCardLogicalOffsetX+mainmenuCardOffsetX+216,
		labelOffsetY+strBoundsY*3+56,
		color.RGBA{0, 210, 0, 255},
	)
	// FPS
	text.Draw(
		screen,
		fmt.Sprintf("%.2f", ebiten.CurrentFPS()),
		assets.CardBodyTextFontFace,
		mainmenuCardLogicalOffsetX+mainmenuCardOffsetX+216,
		labelOffsetY+strBoundsY*4+74,
		color.RGBA{0, 210, 0, 255},
	)
}
