package core

func XYZToGraphics(x, y, z float64) (float64, float64, float64) {
	// Y doesn't need normalization since we're not going under the floor
	// TODO: let's investigate Z
	return -(x*2 - 1.), y * 2 - 1, z
}
