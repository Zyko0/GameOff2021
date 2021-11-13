package core

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/Zyko0/GameOff2021/logic"
)

const (
	RoadWidth  = 1.
	RoadHeight = 1.
	RoadDepth  = 4.

	DefaultSpeed = 1.
	InvulnTime   = 30

	BlockDefaultSpeed = 0.075

	HeartChance = 1.0 // TODO: 0.05
)

type Core struct {
	tick uint64

	invulnTime int
	score      uint64

	PlayerHP int
	Speed    float64
	Distance float64
	Player   *Player
	Blocks   []*Block
	Settings *Settings
}

func NewCore() *Core {
	c := &Core{
		tick: 0,

		invulnTime: 0,
		score:      0,

		PlayerHP: 3,
		Speed:    DefaultSpeed,
		Distance: 0,
		Player:   NewPlayer(),
		Blocks:   []*Block{},
		Settings: newSettings(),
	}

	return c
}

func spawnBlocks(settings *BlockSettings) []*Block {
	blockCount := rand.Intn(settings.MaxBlocksSpawn) + 1
	blocks := make([]*Block, blockCount)
	indices := []int{0, 1, 2, 3, 4}
	for i := 0; i < blockCount; i++ {
		width := BlockWidth0
		// TODO: handle 2nd height ?
		height := BlockHeight0

		idx := rand.Intn(len(indices))
		x := float64(indices[idx]) * width
		indices[idx] = indices[len(indices)-1]
		indices = indices[:len(indices)-1]

		kind := BlockKindRegular
		rng := rand.Float32()
		switch {
		case settings.Heart && rng < HeartChance:
			kind = BlockKindHeart
		case settings.Harder2 && rng < 0.35:
			kind = BlockKindHarder2
		case settings.Harder && rng < 0.65:
			kind = BlockKindHarder
		case !settings.Regular:
			kind = BlockKindHarder
		}

		blocks[i] = newBlock(x, 0, settings.SpawnDepth, width, height, kind)
	}

	return blocks
}

func (l *Core) Update() {
	// Update player on X axis
	if l.Player.intentX != 0 {
		// TODO: don't forget about circular augments
		dx := l.Player.intentX * l.Settings.PlayerSpeed
		// Check collisions with a wall
		if diff := l.Player.x + dx - l.Player.radius; diff < 0 {
			dx -= diff
			if l.Settings.Circular {
				dx = (1 - l.Player.radius) - l.Player.x
			}
		} else if diff := l.Player.x + dx + l.Player.radius; diff > 1 {
			dx -= (diff - 1)
			if l.Settings.Circular {
				dx = -l.Player.x + l.Player.radius
			}
		}
		l.Player.x += dx
	}
	// Every spawninterval ticks, spawn some blocks
	if l.tick%l.Settings.BlockSettings.SpawnInterval == 0 {
		blocks := spawnBlocks(&l.Settings.BlockSettings)
		l.Blocks = append(l.Blocks, blocks...)
	}
	// If in an invulnerability frame, decrement it
	if l.invulnTime > 0 {
		l.invulnTime--
	}
	// Check collisions for blocks
	dz := -(l.Speed * BlockDefaultSpeed)
	for _, b := range l.Blocks {
		// If there's a depth hit and not in an invulnerability frame, check for damage loss
		if l.invulnTime <= 0 {
			// Check z intersection
			if collides, tdz := internal.DepthCollisionPlayerBlock(l.Player, b, dz); collides {
				assets.PlayHitSound()
				l.PlayerHP--
				l.invulnTime = InvulnTime
				// If we know the player is dead, let's adjust the distance of all blocks
				if l.PlayerHP <= l.Settings.defaultSettings.HpToGameOver {
					dz = tdz
				}
				break
			}
		}
	}
	// Update blocks
	for _, b := range l.Blocks {
		b.z += dz
	}
	// Remove any blocks that have fallen off the screen
	for i := 0; i < len(l.Blocks); i++ {
		b := l.Blocks[i]
		if b.z < 2. {
			l.Blocks[i] = l.Blocks[len(l.Blocks)-1]
			l.Blocks = l.Blocks[:len(l.Blocks)-1]
			i--
		}
	}

	l.tick++
	l.Distance += (l.Speed * BlockDefaultSpeed)
	// Every 15 second, increase global speed
	if l.tick%(logic.TPS*15) == 0 {
		l.Speed += 0.25 // TODO: need a higher base speed, and additional speed here as well
	}
}

func (l *Core) GetSpeed() float64 {
	return l.Speed
}

func (l *Core) GetScore() uint64 {
	return uint64(l.Distance)
}

func (l *Core) GetTicks() uint64 {
	return l.tick
}
