package core

const (
	DefaultSpeed = 2.
)

// TODO: handle self-accelerating endless wave past 9 (? Number of possible negative augments)
type Wave struct {
	intSpeed        uint64
	prevIntDistance uint64

	Number      int
	IntDistance uint64
	Speed       float64
	Distance    float64
}

func newWave(number int) *Wave {
	return &Wave{
		intSpeed: 1,

		Number:      number,
		IntDistance: 0,
		Speed:       DefaultSpeed + 0.25*float64(number-1),
		Distance:    0,
	}
}

func (w *Wave) Update() {
	w.IntDistance++ // += w.intSpeed
	w.Distance += (w.Speed * BlockDefaultSpeed)
}
