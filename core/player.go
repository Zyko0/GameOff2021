package core

import (
	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/solarlune/resolv"
)

const (
	DefaultPlayerRadius = 0.0425
)

type Player struct {
	x, y, z float64
	radius  float64

	intentX      float64
	intentAction bool

	hCollider     *resolv.Object
	depthCollider *resolv.Object
}

func NewPlayer() *Player {
	return &Player{
		x:            RoadWidth / 2.,
		y:            DefaultPlayerRadius,
		z:            2.5,
		intentX:      0,
		intentAction: false,
		radius:       DefaultPlayerRadius,
		hCollider: internal.NewPlayerObject(
			RoadWidth/2.-DefaultPlayerRadius,
			0,
			DefaultPlayerRadius,
		),
		depthCollider: internal.NewPlayerObject(
			2.5,
			0,
			DefaultPlayerRadius,
		),
	}
}

func (p *Player) SetIntentX(x float64) {
	p.intentX = x
}

func (p *Player) SetIntentAction(action bool) {
	p.intentAction = action
}

func (p *Player) GetHCollider() *resolv.Object {
	return p.hCollider
}

func (p *Player) GetDCollider() *resolv.Object {
	return p.depthCollider
}

func (p *Player) GetRadius() float64 {
	return p.radius
}

func (p *Player) GetX() float64 {
	return p.x
}

func (p *Player) GetY() float64 {
	return p.y
}

func (p *Player) GetZ() float64 {
	return p.z
}
