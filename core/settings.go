package core

import "github.com/Zyko0/GameOff2021/core/augments"

type baseSettings struct {
	action augments.Action

}

type Settings struct {
	augments []*Augment
}

func newSettings() *Settings {
	return &Settings{
		augments: make([]*Augment, 0),
	}
}

func (s *Settings) applyAugments() {
	
}

type Augment struct {
}

func (s *Settings) AddAugment(a *Augment) {
	s.augments = append(s.augments, a)
	s.applyAugments()
}
