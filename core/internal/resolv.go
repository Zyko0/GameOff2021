package internal

import "github.com/solarlune/resolv"

const (
	SpaceSizeRatio = 100.

	spaceCellSize = 1.
)

func NewBlockObject(x, y, w, h float64, id uint64, tag string) *resolv.Object {
	w *= SpaceSizeRatio
	h *= SpaceSizeRatio
	x *= SpaceSizeRatio
	y *= SpaceSizeRatio
	obj := resolv.NewObject(
		x, y, w, h, tag,
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
		x+w/2, y+w/2, w, h, "player",
	)
	obj.SetShape(resolv.NewCircle(w/2, h/2, w/2))

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
