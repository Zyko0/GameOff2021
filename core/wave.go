package core

const (
	DefaultSpeed = 2.0
	MaxSpeed     = 4.0
)

type Wave struct {
	intSpeed        uint64
	prevIntDistance uint64

	Number      int
	IntDistance uint64
	Speed       float64
	Distance    float64
}

func newWave(number int) *Wave {
	speed := DefaultSpeed + 0.25*float64(number)
	if speed > MaxSpeed {
		speed = MaxSpeed
	}

	return &Wave{
		intSpeed: 1,

		Number:      number,
		IntDistance: 0,
		Speed:       speed,
		Distance:    0,
	}
}

func (w *Wave) Update() {
	w.IntDistance++
	w.Distance += (w.Speed * BlockDefaultSpeed)
}

func (w *Wave) Endless() bool {
	return w.Number >= 6
}
