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

	HeartChance = 0.075 // Let's make it possible to farm them
)

type Core struct {
	tick uint64

	blockSeeds []float32

	invulnTime int
	score      uint64

	PlayerHP      int
	Speed         float64
	Distance      float64
	Player        *Player
	ActionManager *ActionManager
	Blocks        []*Block
	Settings      *Settings
}

func NewCore() *Core {
	seeds := make([]float32, 7)
	for i := range seeds {
		seeds[i] = rand.Float32()
	}
	c := &Core{
		tick: 0,

		blockSeeds: seeds,

		invulnTime: 0,
		score:      0,

		PlayerHP:      3,
		Speed:         DefaultSpeed,
		Distance:      0,
		Player:        NewPlayer(),
		ActionManager: NewActionManager(),
		Blocks:        []*Block{},
		Settings:      newSettings(),
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
		case settings.Harder2 && rng < HeartChance+0.3:
			kind = BlockKindHarder2
		case settings.Harder && rng < HeartChance+0.6:
			kind = BlockKindHarder
		case !settings.Regular:
			kind = BlockKindHarder
		}

		blocks[i] = newBlock(x, 0, settings.SpawnDepth, width, height, kind)
	}

	return blocks
}

func (c *Core) Update() {
	// Update player on X axis
	if c.Player.intentX != 0 {
		// TODO: don't forget about circular augments
		dx := c.Player.intentX * c.Settings.PlayerSpeed * c.Settings.PlayerSpeedModifier
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
	// Every spawninterval ticks, spawn some blocks
	if c.tick%c.Settings.BlockSettings.SpawnInterval == 0 {
		blocks := spawnBlocks(&c.Settings.BlockSettings)
		c.Blocks = append(c.Blocks, blocks...)
	}
	// If in an invulnerability frame, decrement it
	if c.invulnTime > 0 {
		c.invulnTime--
	}
	// Check collisions for blocks
	dz := -(c.Speed * BlockDefaultSpeed)
	for _, b := range c.Blocks {
		// If there's a depth hit and not in an invulnerability frame, check for damage loss
		if c.invulnTime <= 0 {
			// Check z intersection
			if collides, tdz := internal.DepthCollisionPlayerBlock(c.Player, b, dz); collides {
				switch b.kind {
				case BlockKindHeart:
					assets.PlayHeartSound()
					if c.PlayerHP < int(c.Settings.HeartContainers) {
						c.PlayerHP++
					}
				case BlockKindRegular:
					assets.PlayHitSound()
					c.PlayerHP--
				case BlockKindHarder:
					// TODO: Make a different sound
					assets.PlayHitSound()
					c.PlayerHP -= 2
				default:
					// TODO: Make a different sound
					assets.PlayHitSound()
					c.PlayerHP -= 3
				}
				c.invulnTime = InvulnTime
				// If we know the player is dead, let's adjust the distance of all blocks
				if c.PlayerHP <= c.Settings.defaultSettings.HpToGameOver {
					dz = tdz
				}
				break
			}
		}
	}
	// Update blocks
	for _, b := range c.Blocks {
		b.z += dz
	}
	// Remove any blocks that have fallen off the screen
	for i := 0; i < len(c.Blocks); i++ {
		b := c.Blocks[i]
		if b.z < 2. {
			c.Blocks[i] = c.Blocks[len(c.Blocks)-1]
			c.Blocks = c.Blocks[:len(c.Blocks)-1]
			i--
		}
	}

	c.tick++
	c.Distance += (c.Speed * BlockDefaultSpeed)
	// Every 10 seconds, increase global speed
	if c.tick%(logic.TPS*10) == 0 {
		c.Speed += 0.5 // TODO: need a higher base speed, and additional speed here as well
	}
}

func (c *Core) GetBlockSeeds() []float32 {
	return c.blockSeeds
}

func (c *Core) GetSpeed() float64 {
	return c.Speed
}

func (c *Core) GetScore() uint64 {
	return uint64(c.Distance)
}

func (c *Core) GetTicks() uint64 {
	return c.tick
}
