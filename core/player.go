package core

import (
	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/solarlune/resolv"
)

const (
	DefaultPlayerRadius = 0.1
	DefaultPlayerSpeed  = 0.015
)

type Player struct {
	x, y, z float64
	radius  float64

	intentX      float64
	intentAction bool

	hCollider     *resolv.Object
	depthCollider *resolv.Object

	SpeedX float64
}

func NewPlayer() *Player {
	id := internal.GetNextID()
	return &Player{
		x:            RoadWidth / 2.,
		y:            DefaultPlayerRadius,
		z:            2.5,
		intentX:      0,
		intentAction: false,
		radius:       DefaultPlayerRadius,
		hCollider: internal.NewObject(
			RoadWidth/2.-DefaultPlayerRadius,
			0,
			DefaultPlayerRadius*2.,
			DefaultPlayerRadius*2.,
			"player",
			id,
		),
		depthCollider: internal.NewObject(
			2.5,
			0,
			DefaultPlayerRadius*2.,
			DefaultPlayerRadius*2.,
			"player",
			id,
		),

		SpeedX: DefaultPlayerSpeed,
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
