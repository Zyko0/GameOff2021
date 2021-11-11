package ui

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	augmentOffsetX = float32(logic.ScreenWidth-logic.GameSquareDim) / 2

	augmentBgOffsetX = 12
	augmentBgOffsetY = 96
	augmentBgWidth   = logic.GameSquareDim - augmentBgOffsetX*2
	augmentBgHeight  = logic.GameSquareDim - augmentBgOffsetY*2

	augmentCardIntervalOffset = 12
	augmentCardWidth          = (augmentBgWidth - augmentCardIntervalOffset*4) / 3
	augmentCardHeight         = augmentBgHeight/5*4 - augmentCardIntervalOffset
	augmentCardOffsetY        = augmentBgOffsetY + (augmentBgHeight - augmentCardHeight) - augmentCardIntervalOffset
)

type AugmentView struct {
	active bool

	lastCursorX, lastCursorY int
	background               *ebiten.Image

	SelectedIndex int
	Augments      []*augments.Augment
}

func NewAugmentView() *AugmentView {
	bg := ebiten.NewImage(augmentBgWidth, augmentBgHeight)
	graphics.DrawRect(bg, 0, 0, augmentBgWidth, augmentBgHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(bg, 0, 0, augmentBgWidth, augmentBgHeight, 1, []float32{1, 1, 1, 0.9})

	return &AugmentView{
		active: false,

		background: bg,

		SelectedIndex: 0,
		Augments:      nil,
	}
}

func (av *AugmentView) Reset() {
	av.active = false
	av.SelectedIndex = 0
}

func (av *AugmentView) Active() bool {
	return av.active
}

func (av *AugmentView) SetAugments(augments []*augments.Augment) {
	av.Augments = augments
	av.active = true
}

func (av *AugmentView) Update() {
	// Update selection based on keyboard input at last
	var kbInput bool

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		kbInput = true
		av.SelectedIndex++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		kbInput = true
		av.SelectedIndex--
	}
	if av.SelectedIndex < 0 {
		av.SelectedIndex = len(av.Augments) - 1
	}
	if av.SelectedIndex >= len(av.Augments) {
		av.SelectedIndex = 0
	}

	// Check what is on mouse hover only if there has not been any keyboard input and an actual mouse input
	var hovered bool

	x, y := ebiten.CursorPosition()
	hoveredIndex := 0
	y0 := augmentCardOffsetY
	y1 := y0 + augmentCardHeight
	for i := range av.Augments {
		x0 := augmentOffsetX + augmentBgOffsetX + float32(i+1)*augmentCardIntervalOffset + float32(i)*augmentCardWidth
		x1 := x0 + augmentCardWidth
		if float32(x) >= x0 && float32(x) <= x1 && y >= y0 && y <= y1 {
			hoveredIndex = i
			hovered = true
			break
		}
	}
	if !kbInput && (x != av.lastCursorX || y != av.lastCursorY) {
		av.SelectedIndex = hoveredIndex
	}

	// Check if a selection is made
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		av.active = false
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && hovered {
		av.active = false
	}

	av.lastCursorX, av.lastCursorY = x, y
}

func (av *AugmentView) Draw(screen *ebiten.Image) {
	// Background card
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(augmentBgOffsetX+augmentOffsetX),
		float64(augmentBgOffsetY),
	)
	screen.DrawImage(av.background, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Augment cards
	for i := range av.Augments {
		graphics.DrawRect(
			screen,
			augmentOffsetX+augmentBgOffsetX+float32(i+1)*augmentCardIntervalOffset+float32(i)*augmentCardWidth,
			augmentCardOffsetY,
			augmentCardWidth,
			augmentCardHeight,
			[]float32{0.7, 0., 0., 0.5},
		)
		// Highlight selection
		if i == av.SelectedIndex {
			graphics.DrawRectBorder(
				screen,
				augmentOffsetX+augmentBgOffsetX+float32(i+1)*augmentCardIntervalOffset+float32(i)*augmentCardWidth,
				augmentCardOffsetY,
				augmentCardWidth,
				augmentCardHeight,
				1, []float32{1, 1, 1, 0.9},
			)
		}
	}
}
