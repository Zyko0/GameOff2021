package core

const (
	DefaultPlayerRadius = 0.0425
)

type Player struct {
	x, y, z float64
	radius  float64

	intentX      float64
	intentAction bool
}

func NewPlayer() *Player {
	return &Player{
		x:            RoadWidth / 2,
		y:            DefaultPlayerRadius,
		z:            2.5,
		intentX:      0,
		intentAction: false,
		radius:       DefaultPlayerRadius,
	}
}

func (p *Player) SetIntentX(x float64) {
	p.intentX = x
}

func (p *Player) SetIntentAction(action bool) {
	p.intentAction = action
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
