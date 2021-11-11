package ui

import (
	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	pauseCardOffsetX = 96
	pauseCardOffsetY = 208
	pauseCardWidth   = logic.GameSquareDim - pauseCardOffsetX*2
	pauseCardHeight  = logic.GameSquareDim - pauseCardOffsetY*2
)

type PauseView struct {
	active bool

	card *ebiten.Image
}

func NewPauseView() *PauseView {
	card := ebiten.NewImage(pauseCardWidth, pauseCardHeight)
	// Background and border
	graphics.DrawRect(card, 0, 0, pauseCardWidth, pauseCardHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(card, 0, 0, pauseCardWidth, pauseCardHeight, 1, []float32{1, 1, 1, 0.9})
	// Title
	str := "Pause"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(pauseCardWidth/2-rect.Max.X/2),
		float64(48),
	)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Resume text
	str = "Press 'P' to resume the game"
	rect = text.BoundString(assets.CardBodyText, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(pauseCardWidth/2-rect.Max.X/2),
		float64(96),
	)
	text.DrawWithOptions(card, str, assets.CardBodyText, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &PauseView{
		card: card,
	}
}

func (pv *PauseView) Active() bool {
	return pv.active
}

func (pv *PauseView) Reset() {
	pv.active = false
}

func (pv *PauseView) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		pv.active = !pv.active
	}
}

func (pv *PauseView) Draw(screen *ebiten.Image) {
	const (
		offsetX = float32(logic.ScreenWidth-logic.GameSquareDim) / 2
	)

	// Inform about pause, inform about resume by pressing 'P' again
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(offsetX+pauseCardOffsetX),
		float64(pauseCardOffsetY),
	)
	screen.DrawImage(pv.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}
