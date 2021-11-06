package core

import "github.com/Zyko0/GameOff2021/core/augments"

type Action byte

const (
	ActionNone Action = iota
	ActionJump
)

type baseSettings struct {
	action         Action
	hpToGameOver   int
	maxBlocksSpawn int
}

func newBaseSettings() *baseSettings {
	return &baseSettings{
		action:         ActionNone,
		hpToGameOver:   0,
		maxBlocksSpawn: 3,
	}
}

type Settings struct {
	baseSettings   *baseSettings
	actualSettings *baseSettings

	augments []*augments.Augment
}

func newSettings() *Settings {
	return &Settings{
		baseSettings:   newBaseSettings(),
		actualSettings: newBaseSettings(),

		augments: make([]*augments.Augment, 0),
	}
}

func (s *Settings) AddAugment(a *augments.Augment) {
	s.augments = append(s.augments, a)
	s.applyAugments()
}

func (s *Settings) applyAugments() {
	for range s.augments {

	}
}
