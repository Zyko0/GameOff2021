package core

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/logic"
)

const (
	MaxPlayerSpeed = 0.02
)

type Action byte

const (
	ActionNone Action = iota
	ActionJump
)

type baseSettings struct {
	Action         Action
	HpToGameOver   int
	MaxBlocksSpawn int
	LinesDebug     bool
	CameraPosition []float32
	SpawnInterval  uint64
	Circular       bool
	PlayerSpeed    float64
}

func newBaseSettings() *baseSettings {
	return &baseSettings{
		Action:         ActionNone,
		HpToGameOver:   0,
		MaxBlocksSpawn: 3,
		LinesDebug:     false,
		CameraPosition: []float32{0, 0, -1.25},
		SpawnInterval:  logic.TPS * 3,
		Circular:       false,
		PlayerSpeed:    0.01,
	}
}

type Settings struct {
	BaseSettings   *baseSettings
	ActualSettings *baseSettings

	Augments []*augments.Augment
}

func newSettings() *Settings {
	return &Settings{
		BaseSettings:   newBaseSettings(),
		ActualSettings: newBaseSettings(),

		Augments: make([]*augments.Augment, 0),
	}
}

func (s *Settings) AddAugment(a *augments.Augment) {
	s.Augments = append(s.Augments, a)
	s.applyAugments()
}

func (s *Settings) applyAugments() {
	for range s.Augments {

	}
}
