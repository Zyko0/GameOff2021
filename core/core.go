package core

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/assets"
	"github.com/Zyko0/GameOff2021/core/internal"
)

const (
	RoadWidth  = 1.
	RoadHeight = 1.
	RoadDepth  = 4.

	InvulnTime = 20

	BlockDefaultSpeed = 0.075

	HeartChance        = 0.05 // Let's keep it possible to farm them
	GoldenHeartChance  = 0.25
	BlockHarderChance  = 1. / 2.
	BlockHarder2Chance = 1. / 4.
)

type Core struct {
	tick uint64

	invulnTime      int
	blockSeeds      []float32
	score           uint64
	scoreMultiplier uint64

	sfxManager assets.SFXManager

	Wave     *Wave
	PlayerHP int
	Player   *Player
	Blocks   []*Block
	Settings *Settings
}

func NewCore(sfxManager assets.SFXManager) *Core {
	seeds := make([]float32, 7)
	for i := range seeds {
		seeds[i] = 0.5 + rand.Float32()*0.5
	}

	c := &Core{
		tick: 0,

		blockSeeds:      seeds,
		invulnTime:      0,
		score:           0,
		scoreMultiplier: 1,

		sfxManager: sfxManager,

		Wave:     newWave(0),
		PlayerHP: 3, // 3 is okay, since heart are farmable up to 10 early
		Player:   NewPlayer(),
		Blocks:   []*Block{},
		Settings: newSettings(),
	}

	return c
}

func spawnBlocks(settings *BlockSettings) []*Block {
	blockCount := settings.MinBlocksSpawn + rand.Intn(settings.MaxBlocksSpawn-settings.MinBlocksSpawn+1)
	blocks := make([]*Block, blockCount)
	indices := []int{0, 1, 2, 3, 4}
	for i := 0; i < blockCount; i++ {
		kind := BlockKindRegular
		r := rand.Float32()
		switch {
		case settings.Heart && r < HeartChance:
			kind = BlockKindHeart
			if settings.GoldenHeart && rand.Float64() < GoldenHeartChance {
				kind = BlockKindGoldenHeart
			}
		default:
			r = rand.Float32()
			switch {
			case settings.Harder2 && r < BlockHarder2Chance:
				kind = BlockKindHarder2
			case settings.Harder && r < BlockHarderChance:
				kind = BlockKindHarder
			default:
				if !settings.Regular {
					kind = BlockKindHarder
				}
			}
		}

		width := BlockWidth0
		height := BlockHeight0
		taller := settings.TallerBlocks && rand.Intn(2) == 0
		// TODO: ignore taller hearts here
		if taller && kind != BlockKindHeart && kind != BlockKindGoldenHeart {
			height = BlockHeight1
		}
		y := 0.
		higher := !taller && settings.HigherSpawn && rand.Intn(2) == 0
		if higher {
			y = BlockHeight0
		}

		idx := rand.Intn(len(indices))
		x := float64(indices[idx]) * width
		indices[idx] = indices[len(indices)-1]
		indices = indices[:len(indices)-1]

		blocks[i] = newBlock(x, y, settings.SpawnDepth, width, height, kind)
	}

	return blocks
}

func (c *Core) Update() {
	// Update player's jump
	c.Player.jump.Update(c.Player.intentJump)
	// Update player on X axis
	// TODO: below code is trash, but just making it work for now
	if c.Player.intentX != 0 {
		if c.Settings.PerfectStep {
			c.Player.x += c.Player.intentX * 0.2
			if c.Player.x <= 0 {
				c.Player.x = 0.1
				if c.Settings.Circular {
					c.Player.x = 0.9
				}
			}
			if c.Player.x >= 1 {
				c.Player.x = 0.9
				if c.Settings.Circular {
					c.Player.x = 0.1
				}
			}
		} else {
			pSpeed := c.Settings.PlayerSpeed * c.Settings.PlayerSpeedModifier * (1 + (c.Wave.Speed - 1.))
			dx := c.Player.intentX * pSpeed
			// Check collisions with a wall
			if diff := c.Player.x + dx - c.Player.radius; diff < 0 {
				dx -= diff
				if c.Settings.Circular {
					dx = (1 - c.Player.radius) - c.Player.x
				}
			} else if diff := c.Player.x + dx + c.Player.radius; diff > 1 {
				dx -= (diff - 1)
				if c.Settings.Circular {
					dx = -c.Player.x + c.Player.radius
				}
			}
			c.Player.x += dx
		}
	}
	// Every distance interval, spawn some blocks
	// TODO: trying on distance but broken yet
	distMod := c.Wave.IntDistance % uint64(float64(c.Settings.BlockSettings.SpawnDistanceInterval)/c.Wave.Speed)
	if (c.Wave.Endless() || c.Wave.Distance < c.Settings.EndWaveDistance) && distMod == 0 {
		blocks := spawnBlocks(&c.Settings.BlockSettings)
		c.Blocks = append(c.Blocks, blocks...)
	}
	// If in an invulnerability frame, decrement it
	if c.invulnTime > 0 {
		c.invulnTime--
	}
	// Check collisions for blocks
	dz := -(c.Wave.Speed * BlockDefaultSpeed)
	// If there's a depth hit and not in an invulnerability frame, check for damage loss or hp up
	if c.invulnTime <= 0 {
		for _, b := range c.Blocks {
			if collides, tdz := internal.DepthCollisionPlayerTest(c.Player, b, dz); collides {
				switch b.kind {
				case BlockKindHeart:
					c.sfxManager.PlayHeartSound()
					c.PlayerHP++
				case BlockKindGoldenHeart:
					c.sfxManager.PlayHeartSound()
					c.PlayerHP += 2
				case BlockKindRegular:
					c.sfxManager.PlayHitSound()
					c.PlayerHP--
					c.scoreMultiplier = 1
				case BlockKindHarder:
					// TODO: Make a different sound
					c.sfxManager.PlayHitSound()
					c.PlayerHP -= 2
					c.scoreMultiplier = 1
				case BlockKindHarder2:
					// TODO: Make a different sound
					c.sfxManager.PlayHitSound()
					c.PlayerHP -= 3
					c.scoreMultiplier = 1
				}

				c.invulnTime = InvulnTime
				// If we know the player is dead, let's adjust the distance of all blocks
				if c.PlayerHP <= c.Settings.defaultSettings.HpToGameOver {
					dz = tdz
				}
				if c.PlayerHP > int(c.Settings.HeartContainers) {
					c.PlayerHP = int(c.Settings.HeartContainers)
				}
				break
			}
		}
	}
	// Update blocks
	for _, b := range c.Blocks {
		b.z += dz
	}
	// Remove any dead blocks
	for i := 0; i < len(c.Blocks); i++ {
		b := c.Blocks[i]
		// Fallen off the screen
		if b.z < 2. {
			c.Blocks[i] = c.Blocks[len(c.Blocks)-1]
			c.Blocks = c.Blocks[:len(c.Blocks)-1]
			i--
		}
	}

	c.tick++
	c.Wave.Update()

	c.score += uint64(c.Wave.Number+1) * c.scoreMultiplier
}

func (c *Core) IsWaveOver() bool {
	return !c.Wave.Endless() && (c.Wave.Distance > c.Settings.EndWaveDistance+c.Settings.BlockSettings.SpawnDepth+1)
}

func (c *Core) StartNextWave() {
	c.Wave = newWave(c.Wave.Number + 1)
}

func (c *Core) GetBlockSeeds() []float32 {
	return c.blockSeeds
}

func (c *Core) GetSpeed() float64 {
	return c.Wave.Speed
}

func (c *Core) GetScore() uint64 {
	return c.score
}

func (c *Core) GetTicks() uint64 {
	return c.tick
}
