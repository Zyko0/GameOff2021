package automata

const (
	MinSize = 12
)

type Automata struct {
	cells []*Cell
}

func New(cellsCount int) *Automata {
	return &Automata{
		cells: make([]*Cell, cellsCount),
	}
}
