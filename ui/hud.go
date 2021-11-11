package ui

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/hajimehoshi/ebiten/v2"
)

type HUD struct {
	playerHP int
	augments []*augments.Augment
}

func NewHUD(playerHP int, augments []*augments.Augment) *HUD {
	return &HUD{
		playerHP: playerHP,
		augments: augments,
	}
}

// Update updates the hud information before a new draw, only call this on game state change
func (h *HUD) Update(playerHP int, augments []*augments.Augment) {
	h.playerHP = playerHP
	h.augments = augments
}

func (h *HUD) Draw(screen *ebiten.Image) {
	// TODO: Draw hearts as hp, make sur not to go over max hp, also handle negative hp
	// TODO: Draw augments symbols as miniatures somewhere
}
