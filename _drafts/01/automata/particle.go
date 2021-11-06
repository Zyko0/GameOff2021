package automata

import "math"

type Particle struct {
	position [2]int
	value    complex128

	directionModifier int

	direction  [2]int
	floatValue float64
}

func NewParticle(value complex128, position [2]int) *Particle {
	p := &Particle{
		position: position,
		value:    value,

		directionModifier: 1,
	}
	p.Calculate()

	return p
}

func (p *Particle) Calculate() {
	// Float value
	var v float64

	realSign := math.Signbit(real(p.value))
	imagSign := math.Signbit(imag(p.value))
	if realSign == imagSign {
		if realSign {
			v = 0
		} else {
			v = 1
		}
	} else {
		if realSign {
			v = 1. / 3.
		} else {
			v = 2. / 3.
		}
	}
	p.floatValue = v
	// Direction
	var delta = 1

	v = real(p.value) * imag(p.value)
	if math.Signbit(v) {
		delta = -1
	}
	index := int(math.Abs(math.Round(v))) % 2
	p.direction[0], p.direction[1] = 0, 0
	p.direction[index] = delta * p.directionModifier
}

func (p *Particle) Clone() *Particle {
	return &Particle{
		position: p.position,
		value:    p.value,

		directionModifier: p.directionModifier,

		direction:  p.direction,
		floatValue: p.floatValue,
	}
}

func (p *Particle) SetDirectionModifier(modifier int) {
	p.directionModifier = modifier
}

func (p *Particle) GetValue() float64 {
	return p.floatValue
}

func (p *Particle) GetPosition() [2]int {
	return p.position
}

func (p *Particle) GetDirection() [2]int {
	return p.direction
}
