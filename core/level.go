package core

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/solarlune/resolv"
)

const (
	RoadWidth  = 1.
	RoadHeight = 1.
	RoadDepth  = 4.

	DefaultSpeed = 1.
	InvulnTime   = 30
)

type Level struct {
	tick   uint64
	paused bool

	invulnTime int
	playerHP   int
	score      uint64

	hSpace     *resolv.Space
	depthSpace *resolv.Space

	Speed    float64
	Player   *Player
	Blocks   []*Block
	Settings *Settings
}

func NewLevel() *Level {
	level := &Level{
		tick:   0,
		paused: false,

		invulnTime: 0,
		playerHP:   4,
		score:      0,

		hSpace:     internal.NewSpace(RoadWidth, RoadHeight),
		depthSpace: internal.NewSpace(RoadDepth, RoadHeight),

		Speed:    DefaultSpeed,
		Player:   NewPlayer(),
		Blocks:   []*Block{},
		Settings: newSettings(),
	}
	level.hSpace.Add(level.Player.hCollider)
	level.depthSpace.Add(level.Player.depthCollider)

	return level
}

func spawnBlocks(speed float64, maxSpawn int) []*Block {
	blockCount := rand.Intn(maxSpawn) + 1
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

		blocks[i] = newBlock(x, 0, width, height, speed)
	}

	return blocks
}

func (l *Level) Update() {
	// Update player on X axis
	if l.Player.intentX != 0 {
		// Check collisions with a wall
		// TODO: don't forget about circular augments
		dx := l.Player.intentX * l.Player.SpeedX
		nextX := l.Player.x + dx
		if nextX-l.Player.radius < 0 {
			l.Player.x = l.Player.radius
			l.Player.hCollider.X = 0.
		} else if nextX+l.Player.radius > 1 {
			l.Player.x = 1. - l.Player.radius
			l.Player.hCollider.X = (1. - l.Player.radius*2) * internal.SpaceSizeRatio
		} else {
			l.Player.x = nextX
			l.Player.hCollider.X = (l.Player.x - l.Player.radius) * internal.SpaceSizeRatio
		}
		l.Player.hCollider.Update()
	}
	// Every 240 ticks, add a block TODO: this is tmp
	if l.tick%240 == 0 {
		// blocks := spawnBlocks(l.Speed*0.075, l.Settings.actualSettings.maxBlocksSpawn)
		blocks := spawnBlocks(0.075, l.Settings.actualSettings.maxBlocksSpawn)
		for _, b := range blocks {
			l.hSpace.Add(b.hCollider)
			l.depthSpace.Add(b.depthCollider)
			l.Blocks = append(l.Blocks, b)
		}
	}
	// Update all blocks
	for _, b := range l.Blocks {
		b.z -= b.speed
		b.depthCollider.X -= b.speed * internal.SpaceSizeRatio
		// Depth space update
		b.depthCollider.Update()
	}
	// Remove any blocks that have fallen off the screen
	for i := 0; i < len(l.Blocks); i++ {
		b := l.Blocks[i]
		if b.z < 0 {
			l.hSpace.Remove(b.hCollider)
			l.depthSpace.Remove(b.depthCollider)
			l.Blocks[i] = l.Blocks[len(l.Blocks)-1]
			l.Blocks = l.Blocks[:len(l.Blocks)-1]
			i--
		}
	}
	// Take hp if player collided with a block if not in invulnerability frame
	if l.invulnTime > 0 {
		l.invulnTime--
	} else if colDepth := l.Player.depthCollider.Check(0, 0, "block"); colDepth != nil {
		if colHorizon := l.Player.hCollider.Check(0, 0, "block"); colHorizon != nil {
			hit := false
			for _, od := range colDepth.Objects {
				for _, oh := range colHorizon.Objects {
					if od.Data == oh.Data {
						hit = true
					}
				}
			}
			if hit {
				l.playerHP--
				l.invulnTime = InvulnTime
			}
		}
	}

	l.tick++
	l.score++

	// Every 500 score increase speed by 0.5
	if l.score%500 == 0 {
		l.Speed += 0.5
	}
}

func (l *Level) GetPlayerHP() int {
	return l.playerHP
}

func (l *Level) GetSpeed() float64 {
	return l.Speed
}

func (l *Level) GetScore() uint64 {
	return l.score
}

func (l *Level) TogglePause() {
	l.paused = !l.paused
}
