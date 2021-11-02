package automata

type Cell struct {
	X, Y  int
	Value float64
}

func newCell(x, y int, value float64) *Cell {
	return &Cell{
		X:     x,
		Y:     y,
		Value: value,
	}
}
