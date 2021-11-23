package ui

import (
	"github.com/Zyko0/GameOff2021/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

type HUD struct {
	offscreen   *ebiten.Image
	needsRedraw bool

	playerHP int
}

func NewHUD(playerHP int) *HUD {
	return &HUD{
		offscreen:   ebiten.NewImage(logic.GameSquareDim, logic.GameSquareDim),
		needsRedraw: true,

		playerHP: playerHP,
	}
}

// Update updates the hud information before a new draw, only call this on game state change
func (h *HUD) Update(playerHP int) {
	if playerHP != h.playerHP {
		h.needsRedraw = true
	}
	h.playerHP = playerHP
}

func (h *HUD) Draw(offscreen *ebiten.Image) {
	if h.needsRedraw {
		// TODO: draw offscreen
	}

	// TODO: Draw hearts as hp, make sur not to go over max hp, also handle negative hp
	// TODO: Draw augments symbols as miniatures somewhere
	h.needsRedraw = false
	// offscreen.DrawImage(h.offscreen, &ebiten.DrawImageOptions{})
}
