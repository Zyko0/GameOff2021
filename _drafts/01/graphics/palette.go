package graphics

import "math/rand"

func NewRandomPalette() []float32 {
	return []float32{
		rand.Float32(), rand.Float32(), rand.Float32(), 0,
		rand.Float32(), rand.Float32(), rand.Float32(), 1. / 3.,
		rand.Float32(), rand.Float32(), rand.Float32(), 2. / 3.,
		rand.Float32(), rand.Float32(), rand.Float32(), 1,
	}
}
