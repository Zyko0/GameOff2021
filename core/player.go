package core

import (
	"github.com/Zyko0/GameOff2021/logic"
)

const (
	DefaultPlayerRadius = 0.0425

	DefaultJumpDuration = 0.5 * logic.TPS
)

type jump struct {
	icd      uint
	duration uint

	cooldown        uint
	currentDuration uint

	y float64
}

func newJump() *jump {
	return &jump{
		icd:      0,
		duration: DefaultJumpDuration,

		currentDuration: DefaultJumpDuration,

		y: DefaultPlayerRadius,
	}
}

func (m *jump) Update(intent bool) {
	if m.currentDuration < m.duration {
		d := float64(m.currentDuration) / float64(m.duration)
		c := (-(d * d) + d)
		m.y = DefaultPlayerRadius + c
		m.currentDuration++
	} else if m.cooldown > 0 {
		m.cooldown--
	} else if intent {
		m.currentDuration = 1
		m.cooldown = m.icd
	}
}

type Player struct {
	x, y, z float64
	radius  float64

	intentX    float64
	intentJump bool

	jump *jump
}

func NewPlayer() *Player {
	return &Player{
		x:      RoadWidth / 2,
		y:      DefaultPlayerRadius,
		z:      2.5,
		radius: DefaultPlayerRadius,

		intentX:    0,
		intentJump: false,

		jump: newJump(),
	}
}

func (p *Player) SetIntentX(x float64) {
	p.intentX = x
}

func (p *Player) SetIntentJump(intent bool) {
	p.intentJump = intent
}

func (p *Player) GetRadius() float64 {
	return p.radius
}

func (p *Player) GetX() float64 {
	return p.x
}

func (p *Player) GetY() float64 {
	if p.jump.currentDuration < p.jump.duration {
		return p.jump.y
	}

	return p.y
}

func (p *Player) GetZ() float64 {
	return p.z
}
