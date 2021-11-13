package core

const (
	BlockWidth0 = 0.2
	BlockWidth1 = 0.4

	BlockHeight0 = 0.2
	BlockHeight1 = 0.8

	BlockPosY0 = 0.
	BlockPosY1 = 0.4
)

type BlockKind float32

const (
	BlockKindRegular BlockKind = 4
	BlockKindHarder  BlockKind = 5
	BlockKindHarder2 BlockKind = 6
	BlockKindHeart   BlockKind = 7
)

type Block struct {
	x, y, z float64
	kind    BlockKind
}

func newBlock(x, y, z, width, height float64, kind BlockKind) *Block {
	return &Block{
		x:    x + width/2,
		y:    height / 2,
		z:    z,
		kind: kind,
	}
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

func (b *Block) GetKind() BlockKind {
	return b.kind
}
