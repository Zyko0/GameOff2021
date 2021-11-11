package ui

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	augmentBgOffsetX = 12
	augmentBgOffsetY = 96
	augmentBgWidth   = logic.GameSquareDim - augmentBgOffsetX*2
	augmentBgHeight  = logic.GameSquareDim - augmentBgOffsetY*2
)

type AugmentView struct {
	background *ebiten.Image

	SelectedIndex int
	Augments      []*augments.Augment
}

func NewAugmentView() *AugmentView {
	bg := ebiten.NewImage(augmentBgWidth, augmentBgHeight)
	graphics.DrawRect(bg, 0, 0, augmentBgWidth, augmentBgHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(bg, 0, 0, augmentBgWidth, augmentBgHeight, 1, []float32{1, 1, 1, 0.9})

	return &AugmentView{
		background: bg,

		SelectedIndex: 0,
		Augments:      nil,
	}
}

func (av *AugmentView) SetAugments(augments []*augments.Augment) {
	av.Augments = augments
}

func (av *AugmentView) Update(screen *ebiten.Image) {
	// TODO: pick selected on hover + arrow key pressed
	// TODO: Activate selected on click/enter key
}

func (av *AugmentView) Draw(screen *ebiten.Image) {
	const (
		offsetX = float32(logic.ScreenWidth-logic.GameSquareDim) / 2

		AugmentCardIntervalOffset = 12
		AugmentCardWidth          = (augmentBgWidth - AugmentCardIntervalOffset*4) / 3
		AugmentCardHeight         = augmentBgHeight/5*4 - AugmentCardIntervalOffset
		AugmentCardOffsetY        = augmentBgOffsetY + (augmentBgHeight - AugmentCardHeight) - AugmentCardIntervalOffset
	)

	// Background card
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(augmentBgOffsetX+offsetX),
		float64(augmentBgOffsetY),
	)
	screen.DrawImage(av.background, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Augment cards
	for i := range av.Augments {
		graphics.DrawRect(
			screen,
			offsetX+augmentBgOffsetX+float32(i+1)*AugmentCardIntervalOffset+float32(i)*AugmentCardWidth,
			AugmentCardOffsetY,
			AugmentCardWidth,
			AugmentCardHeight,
			[]float32{0.7, 0., 0., 0.5},
		)
	}
}
