package ui

import (
	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	endlessWaveCardOffsetX = 96
	endlessWaveCardOffsetY = 208
	endlessWaveCardWidth   = logic.GameSquareDim - pauseCardOffsetX*2
	endlessWaveCardHeight  = logic.GameSquareDim - pauseCardOffsetY*2
)

type EndlessWaveView struct {
	active bool

	card *ebiten.Image
}

func NewEndlessWaveView() *EndlessWaveView {
	card := ebiten.NewImage(endlessWaveCardWidth, endlessWaveCardHeight)
	// Background and border
	graphics.DrawRect(card, 0, 0, endlessWaveCardWidth, endlessWaveCardHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(card, 0, 0, endlessWaveCardWidth, endlessWaveCardHeight, 1, []float32{1, 1, 1, 0.9})
	// Title
	str := "Endless Wave"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(endlessWaveCardWidth/2-rect.Max.X/2),
		float64(48),
	)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Resume text
	str = "[x] Enable this to step laterally by a column (need to spam left/right)"
	rect = text.BoundString(assets.CardBodyTextFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(endlessWaveCardWidth/2-rect.Max.X/2),
		float64(96),
	)
	text.DrawWithOptions(card, str, assets.CardBodyTextFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &EndlessWaveView{
		card: card,
	}
}

func (ev *EndlessWaveView) Active() bool {
	return ev.active
}

func (ev *EndlessWaveView) Reset() {
	ev.active = false
}

func (ev *EndlessWaveView) Update() {
	// TODO: if something is clicked, enable it + make view inactive
}

func (ev *EndlessWaveView) Draw(screen *ebiten.Image) {
	const (
		offsetX = float32(logic.ScreenWidth-logic.GameSquareDim) / 2
	)

	// Inform about pause, inform about resume by pressing 'P' again
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(offsetX+endlessWaveCardOffsetX),
		float64(endlessWaveCardOffsetY),
	)
	screen.DrawImage(ev.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}
