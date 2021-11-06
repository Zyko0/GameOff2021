package levels

import (
	"math/rand"

	"github.com/Zyko0/GameOff2021/drafts/01/automata"
)

func GetPlusShapeParticles(gridCount int) []*automata.Particle {
	particles := []*automata.Particle{}
	// Middle row
	for x := 0; x < gridCount; x++ {
		nreal := float64(rand.Intn(2))
		nimag := float64(rand.Intn(2))
		if nreal == 0 {
			nreal = -1
		}
		if nimag == 0 {
			nimag = -1
		}
		particles = append(particles, automata.NewParticle(
			complex(float64(rand.Intn(10)+1)*nreal, float64(rand.Intn(10)+1)*nimag),
			[2]int{
				x, gridCount / 2,
			},
		))
	}
	// Middle column
	for y := 0; y < gridCount; y++ {
		nreal := float64(rand.Intn(2))
		nimag := float64(rand.Intn(2))
		if nreal == 0 {
			nreal = -1
		}
		if nimag == 0 {
			nimag = -1
		}
		particles = append(particles, automata.NewParticle(
			complex(float64(rand.Intn(10)+1)*nreal, float64(rand.Intn(10)+1)*nimag),
			[2]int{
				gridCount / 2, y,
			},
		))
	}

	return particles
}
