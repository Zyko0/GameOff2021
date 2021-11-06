package core

import (
	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/solarlune/resolv"
)

const (
	Width  = 1.
	Height = 1.
	Depth  = 4.

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

	Speed  float64
	Player *Player
	Blocks []*Block
}

func NewLevel() *Level {
	level := &Level{
		tick:   0,
		paused: false,

		invulnTime: 0,
		playerHP:   4,
		score:      0,

		hSpace:     internal.NewSpace(Width, Height),
		depthSpace: internal.NewSpace(Depth, Height),

		leftWall:  internal.NewLeftWall(),
		rightWall: internal.NewRightWall(),
		depthWall: internal.NewDepthWall(),

		Speed:  DefaultSpeed,
		Player: NewPlayer(),
		Blocks: []*Block{},
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

func (l *Level) Update() {
	if l.Player.intentX != 0 {
		// Check collisions with a wall
		dx := l.Player.intentX * l.Player.SpeedX * internal.SpaceSizeRatio
		if collision := l.Player.hCollider.Check(dx/2, 0, "wall"); collision != nil {
			dx = collision.ContactWithObject(collision.Objects[0]).X()
		}
		l.Player.hCollider.X += dx / 2
		l.Player.hCollider.Update()
		// Re-divide to make sense for graphics
		l.Player.x += (dx / internal.SpaceSizeRatio)
	}
	// Every 120 ticks, add a block TODO: this is tmp
	if l.tick%120 == 0 {
		b := newBlock(BlockBaseSpeed * l.Speed)
		l.hSpace.Add(b.hCollider)
		l.depthSpace.Add(b.depthCollider)
		l.Blocks = append(l.Blocks, b)
	}
	// Update all blocks
	for _, b := range l.Blocks {
		b.z -= b.speed
		b.depthCollider.X -= b.speed * internal.SpaceSizeRatio
		// Depth space update
		b.depthCollider.Update()
	}
	// Remove any blocks that have fallen off the screen
	for i, b := range l.Blocks {
		if b.z < 0 {
			l.hSpace.Remove(b.hCollider)
			l.depthSpace.Remove(b.depthCollider)
			l.Blocks = append(l.Blocks[:i], l.Blocks[i+1:]...)
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
