package ui

import (
	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	gameoverCardOffsetX = 96
	gameoverCardOffsetY = 208
	gameoverCardWidth   = logic.GameSquareDim - pauseCardOffsetX*2
	gameoverCardHeight  = logic.GameSquareDim - pauseCardOffsetY*2
)

type GameoverView struct {
	active bool

	card *ebiten.Image
}

func NewGameoverView() *GameoverView {
	card := ebiten.NewImage(gameoverCardWidth, gameoverCardHeight)
	graphics.DrawRect(card, 0, 0, gameoverCardWidth, gameoverCardHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(card, 0, 0, gameoverCardWidth, gameoverCardHeight, 1, []float32{1, 1, 1, 0.9})
	// Title
	str := "Game over"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardWidth/2-rect.Max.X/2),
		float64(48),
	)
	colorM := ebiten.ColorM{}
	colorM.Scale(0.8, 0, 0, 1)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Resume text
	str = "Press 'R' to start a new game"
	rect = text.BoundString(assets.CardBodyTextFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardWidth/2-rect.Max.X/2),
		float64(96),
	)
	text.DrawWithOptions(card, str, assets.CardBodyTextFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &GameoverView{
		card: card,
	}
}

func (gv *GameoverView) Active() bool {
	return gv.active
}

func (gv *GameoverView) Update(playerHP, hpToGameOver int) {
	gv.active = playerHP <= hpToGameOver
	if gv.active {
		assets.StopInGameMusic()
		assets.PlayGameoverMusic()
	} else {
		assets.StopGameoverMusic()
	}
}

func (gv *GameoverView) Draw(screen *ebiten.Image) {
	const (
		offsetX = float32(logic.ScreenWidth-logic.GameSquareDim) / 2
	)

	// Inform about pause, inform about resume by pressing 'P' again
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(offsetX+pauseCardOffsetX),
		float64(pauseCardOffsetY),
	)
	screen.DrawImage(gv.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}
