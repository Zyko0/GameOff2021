package internal

import "github.com/solarlune/resolv"

const (
	SpaceSizeRatio = 200.

	spaceCellSize = 1.
)

func NewObject(x, y, w, h float64, tag string) *resolv.Object {
	w *= SpaceSizeRatio
	h *= SpaceSizeRatio
	x *= SpaceSizeRatio
	y *= SpaceSizeRatio
	return resolv.NewObject(
		x-w/2, y-h/2, w, h, tag,
	)
}

func NewLeftWall() *resolv.Object {
	return resolv.NewObject(
		0, 0,
		1.0, SpaceSizeRatio, "wall",
	)
}

func NewRightWall() *resolv.Object {
	return resolv.NewObject(
		SpaceSizeRatio-1.0, 0,
		1.0, SpaceSizeRatio, "wall",
	)
}

func NewDepthWall() *resolv.Object {
	return resolv.NewObject(
		0, 0,
		1.0, SpaceSizeRatio, "wall",
	)
}

func NewSpace(width, height float64) *resolv.Space {
	return resolv.NewSpace(
		// +2 to let space for walls
		int(width*(SpaceSizeRatio+2)),
		int(height*(SpaceSizeRatio+2)),
		spaceCellSize,
		spaceCellSize,
	)
}
