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
	tick uint64

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
		tick: 0,

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
		// TODO: don't forget about circular augments
		dx := l.Player.intentX * l.Settings.ActualSettings.PlayerSpeed
		// Check collisions with a wall
		if diff := l.Player.x + dx - l.Player.radius; diff < 0 {
			dx -= diff
			if l.Settings.ActualSettings.Circular {
				dx = (1 - l.Player.radius) - l.Player.x
			}
		} else if diff := l.Player.x + dx + l.Player.radius; diff > 1 {
			dx -= (diff - 1)
			if l.Settings.ActualSettings.Circular {
				dx = -l.Player.x + l.Player.radius
			}
		}
		l.Player.x += dx
		l.Player.hCollider.X += (dx * internal.SpaceSizeRatio)
		l.Player.hCollider.Update()
	}
	// Every spawninterval ticks, spawn some
	if l.tick%l.Settings.ActualSettings.SpawnInterval == 0 {
		blocks := spawnBlocks(l.Speed*0.075, l.Settings.ActualSettings.MaxBlocksSpawn)
		for _, b := range blocks {
			l.hSpace.Add(b.hCollider)
			l.depthSpace.Add(b.depthCollider)
			l.Blocks = append(l.Blocks, b)
		}
	}
	// If in an invulnerability frame, decrement it
	if l.invulnTime > 0 {
		l.invulnTime--
	}
	// Update all blocks
	for _, b := range l.Blocks {
		dw := -b.speed * internal.SpaceSizeRatio
		// If there's a depth hit and not in an invulnerability frame, check for damage loss
		if l.invulnTime <= 0 {
			// Check z intersection
			if depthHit := b.depthCollider.Shape.Intersection(dw, 0, l.Player.depthCollider.Shape); depthHit != nil {
				// Check x intersection with the same object
				for _, oh := range l.hSpace.Objects() {
					if oh.Data != b.depthCollider.Data {
						continue
					}
					if oh.Shape.Intersection(0, 0, l.Player.hCollider.Shape) != nil {
						l.playerHP--
						l.invulnTime = InvulnTime
						// If player is dead, let's show where the block has hit
						// TODO: not sure it's nice
						/*
							if l.playerHP <= l.Settings.actualSettings.hpToGameOver {
								dw = depthHit.MTV.X()
							}
						*/
					}
				}
			}
		}
		b.z += (dw / internal.SpaceSizeRatio)
		// Depth space update
		b.depthCollider.X += dw
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
