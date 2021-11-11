package core

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/logic"
)

const (
	MaxPlayerSpeed         = 0.02
	MaxHeartContainers     = 10
	SlowMotionTickInterval = logic.TPS * 20 // Trigger a slowmo every 20 secs
	SlowMotionTickDuration = logic.TPS * 2  // Make it last 2 secs to allow player re-calibration
)

type Action byte

const (
	ActionNone Action = iota
	ActionJump
	ActionDash
)

type baseSettings struct {
	Action                Action
	HpToGameOver          int
	SlowMotion            bool
	HeartSpawn            bool
	HeartContainers       uint
	PerfectStep           bool
	MaxBlocksSpawn        int
	DebugLines            bool
	CameraPosition        []float32
	SpawnInterval         uint64
	SpawnDepth            float64
	StrongerBlocks        bool
	EvenStrongerBlocks    bool
	Circular              bool
	PlayerSpeed           float64
	AugmentsTicksInterval uint64
}

func newBaseSettings() *baseSettings {
	return &baseSettings{
		Action:                ActionNone,
		HpToGameOver:          0,
		SlowMotion:            false,
		HeartSpawn:            false,
		HeartContainers:       3,
		MaxBlocksSpawn:        3,
		DebugLines:            false,
		CameraPosition:        []float32{0, 0, -1.25},
		SpawnInterval:         logic.TPS * 3,
		SpawnDepth:            27,
		StrongerBlocks:        false,
		EvenStrongerBlocks:    false,
		Circular:              false,
		PlayerSpeed:           0.01,
		AugmentsTicksInterval: logic.TPS * 20,
	}
}

type Settings struct {
	*baseSettings

	defaultSettings *baseSettings
}

func newSettings() *Settings {
	return &Settings{
		baseSettings:    newBaseSettings(),
		defaultSettings: newBaseSettings(),
	}
}

func (s *Settings) ApplyAugments(currentAugments []*augments.Augment) {
	*s.baseSettings = *s.defaultSettings

	for _, a := range currentAugments {
		switch a.ID {
		case augments.IDIncreaseSpeed:
			s.PlayerSpeed = s.defaultSettings.PlayerSpeed + 0.002
		case augments.IDDecreaseSpeed:
			s.PlayerSpeed = s.defaultSettings.PlayerSpeed - 0.002
		case augments.IDDebugLines:
			s.DebugLines = true
		case augments.IDActionJump:
			s.Action = ActionJump
		case augments.IDActionDash:
			s.Action = ActionDash
		case augments.IDHighSpawn:
			// TODO: higher spawn possible
		case augments.IDSlowMotion:
			s.SlowMotion = true
		case augments.IDHeartSpawn:
			s.HeartSpawn = true
		case augments.IDHeartContainer:
			s.HeartContainers++
			if s.HeartContainers > MaxHeartContainers {
				s.HeartContainers = MaxHeartContainers
			}
		case augments.IDNegativeHearts:
			s.HpToGameOver = -3
		case augments.IDCircular:
			s.Circular = true
		case augments.IDPerfectStep:
			s.PerfectStep = true
		case augments.IDOneMoreBlock:
			s.MaxBlocksSpawn = 4
		case augments.IDTallerBlocks:
			// TODO: taller blocks
		case augments.IDTopView:
			// TODO: let's see if we want to do that camera stuff, might break shader optimizations
		case augments.IDMoreSpawns:
			s.SpawnInterval = logic.TPS * 2
		case augments.IDEvenMoreSpawns:
			s.SpawnInterval = logic.TPS * 1
		case augments.IDCloserSpawns:
			s.SpawnDepth = 27 * 2 / 3
		case augments.IDCloserSpawns2:
			s.SpawnDepth = 27 * 1 / 3
		case augments.IDNothing, augments.IDNothing2:
			// TODO: nothing
		case augments.IDLessAugments:
			s.AugmentsTicksInterval = logic.TPS * 40
		case augments.IDHarderBlocks:
			s.StrongerBlocks = true
		case augments.IDHarderBlocks2:
			s.EvenStrongerBlocks = true
		}
	}
}
