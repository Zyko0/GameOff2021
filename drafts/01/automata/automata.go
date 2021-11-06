package automata

const (
	MinSize = 12
	Steps   = 20
)

type Automata struct {
	gridCount int
	walls     []*Wall

	currentStep     int
	particlesByStep [Steps][]*Particle
}

func New(gridRowCount int, walls []*Wall, particles []*Particle) *Automata {
	particlesByStep := [Steps][]*Particle{}
	for i := 0; i < Steps-1; i++ {
		particlesByStep[i] = make([]*Particle, 0)
	}
	particlesByStep[Steps-1] = particles

	a := &Automata{
		gridCount: gridRowCount,
		walls:     walls,

		currentStep:     Steps - 1,
		particlesByStep: particlesByStep,
	}
	a.propagateBackward()

	return a
}

func (a *Automata) GetGridCount() int {
	return a.gridCount
}

func fightParticles(particles []*Particle) {
	for i, p := range particles {
		pv := p.GetValue()
		for j, po := range particles {
			if i == j {
				continue
			}
			if po.GetValue() > pv {
				p.value *= po.value
			}
		}
	}
}

func (a *Automata) propagateBackward() {
	for step := Steps - 1; step > 0; step-- {
		newParticles := make([]*Particle, len(a.particlesByStep[step]))
		for i, p := range a.particlesByStep[step] {
			newParticles[i] = p.Clone()
		}
		// Move particles in their respective directions first
		for _, particle := range newParticles {
			d := particle.GetDirection()
			particle.position[0] += d[0]
			particle.position[1] += d[1]
			// Check if it hits a wall
			wallReaction := WallEffectNone
			wallPosition := WallPositionBottom
			if particle.position[0] < 0 {
				particle.position[0] = 0
				wallReaction = a.walls[WallPositionLeft].reaction
				wallPosition = WallPositionLeft
			}
			if particle.position[0] >= a.gridCount {
				particle.position[0] = a.gridCount - 1
				wallReaction = a.walls[WallPositionRight].reaction
				wallPosition = WallPositionRight
			}
			if particle.position[1] < 0 {
				particle.position[1] = 0
				wallReaction = a.walls[WallPositionTop].reaction
				wallPosition = WallPositionTop
			}
			if particle.position[1] >= a.gridCount {
				particle.position[1] = a.gridCount - 1
				wallReaction = a.walls[WallPositionBottom].reaction
				wallPosition = WallPositionBottom
			}
			// Apply reaction if exists
			if wallReaction != WallEffectNone {
				switch wallReaction {
				case WallEffectTeleportOpposite:
					particle.SetDirectionModifier(1)
					switch wallPosition {
					case WallPositionLeft:
						particle.position[0] = a.gridCount - 1
					case WallPositionRight:
						particle.position[0] = 0
					case WallPositionTop:
						particle.position[1] = a.gridCount - 1
					case WallPositionBottom:
						particle.position[1] = 0
					}
				case WallEffectReverseDirection:
					particle.SetDirectionModifier(-1)
				}
			}
		}
		// Fight particles sharing the same coordinates
		for y := 0; y < a.gridCount; y++ {
			for x := 0; x < a.gridCount; x++ {
				var particles []*Particle

				for _, p := range newParticles {
					if p.position[0] == x && p.position[1] == y {
						particles = append(particles, p)
					}
				}
				if len(particles) > 1 {
					fightParticles(particles)
				}
			}
		}
		// Calculate particles
		for _, p := range newParticles {
			p.Calculate()
		}
		// Set the new particles to the previous step
		a.particlesByStep[step-1] = newParticles
	}
}

func (a *Automata) GetParticlesAtStep() []*Particle {
	return a.particlesByStep[a.currentStep]
}

func (a *Automata) GetCurrentStep() int {
	return a.currentStep
}

func (a *Automata) Advance() {
	if a.currentStep >= Steps-1 {
		return
	}

	a.currentStep++
}

func (a *Automata) Rewind() {
	if a.currentStep == 0 {
		return
	}

	a.currentStep--
}
