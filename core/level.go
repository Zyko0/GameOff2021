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

	leftWall  *resolv.Object
	rightWall *resolv.Object
	depthWall *resolv.Object

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

		leftWall:  internal.NewLeftWall(),
		rightWall: internal.NewRightWall(),
		depthWall: internal.NewDepthWall(),

		Speed:    DefaultSpeed,
		Player:   NewPlayer(),
		Blocks:   []*Block{},
		Settings: newSettings(),
	}
	level.hSpace.Add(
		level.Player.hCollider,
		level.leftWall,
		level.rightWall,
	)
	level.depthSpace.Add(
		level.Player.depthCollider,
		level.depthWall,
	)

	return level
}

func spawnBlocks(speed float64, maxSpawn int) []*Block {
	blockCount := rand.Intn(maxSpawn) + 1
	blocks := make([]*Block, blockCount)
	indices := []int{1, 2, 3, 4, 5}
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
	if l.Player.intentX != 0 {
		// Check collisions with a wall
		dx := l.Player.intentX * l.Player.SpeedX * internal.SpaceSizeRatio
		if collision := l.Player.hCollider.Check(dx, 0, "wall"); collision != nil {
			dx = collision.ContactWithObject(collision.Objects[0]).X()
		}
		l.Player.hCollider.X += dx
		l.Player.hCollider.Update()
		// Re-divide to make sense for graphics
		l.Player.x += (dx / internal.SpaceSizeRatio)
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
	} else if l.Player.depthCollider.Check(0, 0, "block") != nil {
		if l.Player.hCollider.Check(0, 0, "block") != nil {
			l.playerHP--
			l.invulnTime = InvulnTime
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
