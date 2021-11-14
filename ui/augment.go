package ui

import (
	"fmt"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/graphics"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
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

	augmentDescriptionCardOffset  = 8
	augmentDescriptionCardWidth   = augmentCardWidth - augmentDescriptionCardOffset*2
	augmentDescriptionCardHeight  = augmentCardHeight/5*3 - augmentDescriptionCardOffset
	augmentDescriptionCardOffsetY = augmentCardHeight - augmentDescriptionCardHeight - augmentDescriptionCardOffset
)

var (
	augmentBgColorCommon    = []float32{0.36, 0.48, 0.92, 0.7}
	augmentBgColorEpic      = []float32{2. / 3., 0.078, 0.94, 0.7}
	augmentBgColorLegendary = []float32{1.0, 0.75, 0, 0.7}
	augmentBgColorNegative  = []float32{0.7, 0, 0, 0.7}

	augmentDescriptionBgColor = []float32{1, 1, 1, 0.3}
)

type AugmentView struct {
	active bool

	lastCursorX int
	lastCursorY int
	card        *ebiten.Image

	SelectedIndex int
	Augments      []*augments.Augment
}

func NewAugmentView() *AugmentView {
	card := ebiten.NewImage(augmentBgWidth, augmentBgHeight)
	graphics.DrawRect(card, 0, 0, augmentBgWidth, augmentBgHeight, []float32{0.1, 0.1, 0.1, 0.9})
	graphics.DrawRectBorder(card, 0, 0, augmentBgWidth, augmentBgHeight, 1, []float32{1, 1, 1, 0.9})
	// Title
	str := "Exploit a bug"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(augmentBgWidth/2-rect.Max.X/2),
		float64(48),
	)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &AugmentView{
		active: false,

		card: card,

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
	screen.DrawImage(av.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Augment cards
	for i := range av.Augments {
		var clr []float32

		switch av.Augments[i].Rarity {
		case augments.RarityCommon:
			clr = augmentBgColorCommon
		case augments.RarityEpic:
			clr = augmentBgColorEpic
		case augments.RarityLegendary:
			clr = augmentBgColorLegendary
		case augments.RarityNegative:
			clr = augmentBgColorNegative
		}
		// Card rectangle
		x := augmentOffsetX + augmentBgOffsetX + float32(i+1)*augmentCardIntervalOffset + float32(i)*augmentCardWidth
		y := float32(augmentCardOffsetY)
		graphics.DrawRect(
			screen,
			x, y,
			augmentCardWidth,
			augmentCardHeight,
			clr,
		)
		// Card title text
		rect := text.BoundString(assets.CardBodyTitleFontFace, av.Augments[i].Name)
		geom := ebiten.GeoM{}
		geom.Translate(
			float64(x)+float64(augmentCardWidth)/2-float64(rect.Max.X)/2,
			float64(y)+32,
		)
		text.DrawWithOptions(screen, av.Augments[i].Name, assets.CardBodyTitleFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Description rectangle
		graphics.DrawRect(
			screen,
			x+augmentDescriptionCardOffset,
			y+augmentDescriptionCardOffsetY,
			augmentDescriptionCardWidth,
			augmentDescriptionCardHeight,
			augmentDescriptionBgColor,
		)
		// Description text body
		geom = ebiten.GeoM{}
		geom.Translate(
			float64(x)+float64(augmentDescriptionCardOffset)+4,
			float64(y)+augmentDescriptionCardOffsetY+16,
		)
		rect = text.BoundString(assets.CardBodyDescriptionTextFontFace, av.Augments[i].Description)
		text.DrawWithOptions(screen, av.Augments[i].Description, assets.CardBodyDescriptionTextFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Cost
		if cost := av.Augments[i].GetCost(); cost.Kind != augments.CostNone {
			str := "Cost: "
			switch cost.Kind {
			case augments.CostHP:
				str += fmt.Sprintf("%d HP", cost.Value)
			}
			geom = ebiten.GeoM{}
			geom.Translate(
				float64(x)+float64(augmentDescriptionCardOffset)+4,
				float64(y)+216,
			)
			text.DrawWithOptions(screen, str, assets.CardBodyTextFontFace, &ebiten.DrawImageOptions{
				GeoM: geom,
			})
		}
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
