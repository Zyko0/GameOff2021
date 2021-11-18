package internal

type Player interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	GetRadius() float64
}

type Block interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	GetWidth() float64
	GetHeight() float64
}

func DepthCollisionPlayerTest(p Player, block Block, blockDeltaZ float64) (bool, float64) {
	pX := p.GetX()
	pY := p.GetY()
	pZ := p.GetZ()
	pR := p.GetRadius()
	bX := block.GetX()
	bY := block.GetY()
	bZ := block.GetZ()
	bW := block.GetWidth()
	bH := block.GetHeight()
	// There's a collision on 2D plane
	if pX+pR > bX-bW/2 && pX-pR < bX+bW/2 && pY+pR > bY-bH/2 && pY-pR < bY+bH/2 {
		// Dirty brute force to check for a collision on the Z axis at high speed
		// TODO: make this not bruteforce
		// Note: Udpate => It doesn't seem to loop too much anyways, pretty much a wontfix
		// Step by halfblock width multiple times
		for dz := blockDeltaZ; dz < 0; dz += bW {
			if pZ+pR > bZ+dz-bW && pZ-pR < bZ+dz+bW {
				// There's a collision on 3D plane
				return true, dz
			}
		}
	}
	return false, 0.
}
