package internal

import "github.com/solarlune/resolv"

const (
	SpaceSizeRatio = 100.

	spaceCellSize = 1.
)

func NewObject(x, y, w, h float64, tag string, id uint64) *resolv.Object {
	w *= SpaceSizeRatio
	h *= SpaceSizeRatio
	x *= SpaceSizeRatio
	y *= SpaceSizeRatio
	obj := resolv.NewObject(
		x, y, w, h, tag,
	)
	obj.Data = id
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
