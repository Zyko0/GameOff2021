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
	MinBlocksSpawn                               int
	MaxBlocksSpawn                               int
	SpawnDistanceInterval                        uint64
	SpawnDepth                                   float64
	Regular, Harder, Harder2, Heart, GoldenHeart bool
	HigherSpawn                                  bool
	TallerBlocks                                 bool
}

type baseSettings struct {
	Action                Action
	HpToGameOver          int
	SlowMotion            bool
	HeartContainers       uint
	PerfectStep           bool
	DebugLines            float32
	CameraPosition        []float32
	Circular              bool
	PlayerSpeed           float64
	PlayerSpeedModifier   float64
	AugmentsTicksInterval uint64
	EndWaveDistance       float64
	BlockSettings         BlockSettings
}

func newBaseSettings() *baseSettings {
	return &baseSettings{
		Action:                ActionNone,
		HpToGameOver:          0,
		SlowMotion:            false,
		HeartContainers:       3,
		PerfectStep:           false,
		DebugLines:            0.,
		CameraPosition:        []float32{0, -0.2, -1.15},
		Circular:              false,
		PlayerSpeed:           0.01,
		PlayerSpeedModifier:   1.,
		AugmentsTicksInterval: logic.TPS * 20,
		EndWaveDistance:       150.,
		BlockSettings: BlockSettings{
			MinBlocksSpawn:        3,
			MaxBlocksSpawn:        4,
			SpawnDistanceInterval: 80,
			SpawnDepth:            27,
			Regular:               true,
			Harder:                false,
			Harder2:               false,
			Heart:                 false,
			GoldenHeart:           false,
			HigherSpawn:           false,
			TallerBlocks:          false,
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
		case augments.IDDebugLines:
			s.DebugLines = 1.
		case augments.IDHighSpawn:
			s.BlockSettings.HigherSpawn = true
		case augments.IDSlowMotion:
			s.SlowMotion = true
		case augments.IDHeartSpawn:
			s.BlockSettings.Heart = true
		case augments.IDGoldHeartSpawn:
			s.BlockSettings.GoldenHeart = true
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
		case augments.IDMoreBlocks:
			s.BlockSettings.MaxBlocksSpawn = 4
			s.BlockSettings.MinBlocksSpawn = 3
		case augments.IDTallerBlocks:
			s.BlockSettings.TallerBlocks = true
		case augments.IDMoreSpawns:
			s.BlockSettings.SpawnDistanceInterval = 70
		case augments.IDEvenMoreSpawns:
			s.BlockSettings.SpawnDistanceInterval = 60
		case augments.IDCloserSpawns:
			s.BlockSettings.SpawnDepth = 27 * 3 / 4
		case augments.IDCloserSpawns2:
			s.BlockSettings.SpawnDepth = 27 * 1 / 2
		case augments.IDNothing, augments.IDNothing2, augments.IDNothing3, augments.IDNothing4:
			// Nothing
		case augments.IDHarderBlocks:
			s.BlockSettings.Harder = true
		case augments.IDHarderBlocks2:
			s.BlockSettings.Harder2 = true
		case augments.IDNoRegularBlocks:
			s.BlockSettings.Regular = false
		}
	}
}
