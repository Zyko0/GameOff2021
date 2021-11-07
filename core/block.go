package core

import (
	"github.com/Zyko0/GameOff2021/core/internal"
	"github.com/solarlune/resolv"
)

const (
	DefaultSpawnDepth = 30.

	BlockBaseSpeed = 0.1

	BlockWidth0 = 0.2
	BlockWidth1 = 0.4

	BlockHeight0 = 0.2
	BlockHeight1 = 0.8

	BlockPosY0 = 0.
	BlockPosY1 = 0.4
)

type Block struct {
	x, y, z       float64
	hCollider     *resolv.Object
	depthCollider *resolv.Object

	speed float64
}

func newBlock(x, y, width, height, speed float64) *Block {
	id := internal.GetNextID()
	return &Block{
		x: x + width/2,
		y: height / 2,
		z: DefaultSpawnDepth,
		hCollider: internal.NewBlockObject(
			x,
			y,
			width,
			height,
			id,
		),
		depthCollider: internal.NewBlockObject(
			DefaultSpawnDepth,
			y,
			width,
			height,
			id,
		),

		speed: speed,
	}
}

func (b *Block) GetHCollider() *resolv.Object {
	return b.hCollider
}

func (b *Block) GetX() float64 {
	return b.x
}

func (b *Block) GetY() float64 {
	return b.y
}

func (b *Block) GetZ() float64 {
	return b.z
}

func (b *Block) GetWidth() float64 {
	return BlockWidth0
}

func (b *Block) GetHeight() float64 {
	return BlockHeight0
}
