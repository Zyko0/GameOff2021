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

type BlockSettings struct {
	MaxBlocksSpawn                  int
	SpawnInterval                   uint64
	SpawnDepth                      float64
	Regular, Harder, Harder2, Heart bool
}

type baseSettings struct {
	Action                Action
	HpToGameOver          int
	SlowMotion            bool
	HeartContainers       uint
	PerfectStep           bool
	DebugLines            bool
	CameraPosition        []float32
	Circular              bool
	PlayerSpeed           float64
	PlayerSpeedModifier   float64
	AugmentsTicksInterval uint64
	BlockSettings         BlockSettings
}

func newBaseSettings() *baseSettings {
	return &baseSettings{
		Action:                ActionNone,
		HpToGameOver:          0,
		SlowMotion:            false,
		HeartContainers:       3,
		PerfectStep:           false,
		DebugLines:            false,
		CameraPosition:        []float32{0, 0, -1.25},
		Circular:              false,
		PlayerSpeed:           0.01,
		PlayerSpeedModifier:   1.,
		AugmentsTicksInterval: logic.TPS * 20,
		BlockSettings: BlockSettings{
			MaxBlocksSpawn: 3,
			SpawnInterval:  logic.TPS * 2,
			SpawnDepth:     27,
			Regular:        true,
			Harder:         true,
			Harder2:        true,
			Heart:          true,
		},
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
			s.BlockSettings.Heart = true
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
			s.BlockSettings.MaxBlocksSpawn = 4
		case augments.IDTallerBlocks:
			// TODO: taller blocks
		case augments.IDTopView:
			// TODO: let's see if we want to do that camera stuff, might break shader optimizations
		case augments.IDMoreSpawns:
			s.BlockSettings.SpawnInterval = logic.TPS * 1.5
		case augments.IDEvenMoreSpawns:
			s.BlockSettings.SpawnInterval = logic.TPS * 1
		case augments.IDCloserSpawns:
			s.BlockSettings.SpawnDepth = 27 * 3 / 4
		case augments.IDCloserSpawns2:
			s.BlockSettings.SpawnDepth = 27 * 1 / 2
		case augments.IDNothing, augments.IDNothing2:
			// TODO: nothing
		case augments.IDLessAugments:
			s.AugmentsTicksInterval = logic.TPS * 40
		case augments.IDHarderBlocks:
			s.BlockSettings.Harder = true
		case augments.IDHarderBlocks2:
			s.BlockSettings.Harder2 = true
		case augments.IDNoRegularBlocks:
			s.BlockSettings.Regular = false
		case augments.IDFourTimesFaster:
			s.PlayerSpeedModifier = 2
		}
	}
}
