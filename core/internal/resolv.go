package internal

import "github.com/solarlune/resolv"

const (
	SpaceSizeRatio = 100.

	spaceCellSize = 1.
)

func NewBlockObject(x, y, w, h float64, id uint64) *resolv.Object {
	w *= SpaceSizeRatio
	h *= SpaceSizeRatio
	x *= SpaceSizeRatio
	y *= SpaceSizeRatio
	obj := resolv.NewObject(
		x, y, w, h, "block",
	)
	obj.SetShape(resolv.NewRectangle(0, 0, w, h))
	obj.Data = id

	return obj
}

func NewPlayerObject(x, y, r float64) *resolv.Object {
	w := r * 2 * SpaceSizeRatio
	h := r * 2 * SpaceSizeRatio
	x *= SpaceSizeRatio
	y *= SpaceSizeRatio
	obj := resolv.NewObject(
		x, y, w, h, "player",
	)
	obj.SetShape(resolv.NewCircle(x+r, y+r, r))

	return obj
}

func NewSpace(width, height float64) *resolv.Space {
	return resolv.NewSpace(
		int(width*SpaceSizeRatio),
		int(height*SpaceSizeRatio),
		spaceCellSize,
		spaceCellSize,
	)
}
