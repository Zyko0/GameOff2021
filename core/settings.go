package core

import (
	"github.com/Zyko0/GameOff2021/core/augments"
	"github.com/Zyko0/GameOff2021/logic"
)

const (
	MaxPlayerSpeed     = 0.02
	MaxHeartContainers = 10
)

type Action byte

const (
	ActionNone Action = iota
	ActionJump
	ActionDash
)

type BlockSettings struct {
	MinBlocksSpawn        int
	MaxBlocksSpawn        int
	SpawnDistanceInterval uint64
	SpawnDepth            float64
	Regular               bool
	Harder                bool
	Harder2               bool
	Heart                 bool
	GoldenHeart           bool
	LateralHole           bool
	LongHole              bool
	ChargingBeam          bool
	HigherSpawn           bool
	TallerBlocks          bool
}

type baseSettings struct {
	Action                Action
	HpToGameOver          int
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
		HeartContainers:       10,
		PerfectStep:           false, // TODO: Remove
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
			Heart:                 true,
			GoldenHeart:           true,
			LateralHole:           false,
			LongHole:              false,
			ChargingBeam:          false, // TODO: disable
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
		case augments.IDHighSpawn:
			s.BlockSettings.HigherSpawn = true
		case augments.IDCircular:
			s.Circular = true
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
		case augments.IDHarderBlocks:
			s.BlockSettings.Harder = true
		case augments.IDHarderBlocks2:
			s.BlockSettings.Harder2 = true
		case augments.IDNoRegularBlocks:
			s.BlockSettings.Regular = false
		}
	}
}
