package ui

import (
	"image/color"
	"strconv"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	hudOffsetX = (logic.ScreenWidth-logic.GameSquareDim)/2 + 16
	hudOffsetY = 16
)

var (
	whiteImg *ebiten.Image
	heartImg *ebiten.Image

	waveStrBounds         = text.BoundString(assets.CardBodyTitleFontFace, "Wave ∞")
	comboStrBounds        = text.BoundString(assets.CardBodyTitleFontFace, "Combo x999")
	biggestScoreStrBounds = text.BoundString(assets.CardBodyTitleFontFace, "Score 9999999")
)

func init() {
	whiteImg = ebiten.NewImage(1, 1)
	whiteImg.Fill(color.White)

	width := float32(128)
	height := float32(128)
	heartImg = ebiten.NewImage(int(width), int(height))

	p := vector.Path{}
	p.MoveTo(width/2, height/5)
	p.CubicTo(5*width/14, 0, 0, height/15, width/28, 2*height/5)
	p.CubicTo(width/14, 2*height/3, 3*width/7, 5*height/6, width/2, height)
	p.CubicTo(4*width/7, 5*height/6, 13*width/14, 2*height/3, 27*width/28, 2*height/5)
	p.CubicTo(width, height/15, 9*width/14, 0, width/2, height/5)

	vertices, indices := p.AppendVerticesAndIndicesForFilling(nil, nil)
	heartImg.DrawTriangles(vertices, indices, whiteImg, nil)
}

type HUD struct {
	playerHP   int
	score      uint64
	wave       int
	multiplier uint64
}

func NewHUD(playerHP int, score uint64, wave int, multiplier uint64) *HUD {
	return &HUD{
		playerHP:   playerHP,
		score:      score,
		wave:       wave,
		multiplier: multiplier,
	}
}

// Update updates the hud information before a new draw, only call this on game state change
func (h *HUD) Update(playerHP int, score uint64, wave int, multiplier uint64) {
	h.playerHP = playerHP
	h.score = score
	h.wave = wave
	h.multiplier = multiplier
}

func (h *HUD) Draw(screen *ebiten.Image) {
	// Health points
	for i := 0; i < h.playerHP; i++ {
		op := &ebiten.DrawImageOptions{
			Filter: ebiten.FilterLinear,
		}
		op.GeoM.Scale(0.125, 0.125)
		op.GeoM.Translate(hudOffsetX+float64(i*16), hudOffsetY)
		op.ColorM.Scale(1, 0, 0, 1)
		screen.DrawImage(heartImg, op)
	}
	// Print wave number
	waveStr := "Wave " + string('0'+rune(h.wave+1)) // This is ugly but I believe this is cheaper than fmt.Sprintf / strconv stuff
	if h.wave > 5 {
		waveStr = "Wave ∞"
	}
	text.Draw(
		screen,
		waveStr,
		assets.CardBodyTitleFontFace,
		screen.Bounds().Dx()/2-waveStrBounds.Dx()/2,
		hudOffsetY+waveStrBounds.Dy(),
		color.White,
	)
	// Print score
	text.Draw(
		screen,
		"Score "+strconv.FormatUint(h.score, 10),
		assets.CardBodyTitleFontFace,
		screen.Bounds().Dx()-biggestScoreStrBounds.Dx()-hudOffsetX,
		hudOffsetY+comboStrBounds.Dy(),
		color.White,
	)
	// Print score multiplier (combo)
	text.Draw(
		screen,
		"x"+strconv.FormatUint(h.multiplier, 10),
		assets.CardBodyTitleFontFace,
		screen.Bounds().Dx()-biggestScoreStrBounds.Dx()-hudOffsetX,
		hudOffsetY+waveStrBounds.Dy()*2+16,
		color.White,
	)
}
